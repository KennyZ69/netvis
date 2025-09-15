/* client.c */

#include "../include/client.h"
#include "../data/parse_json.h"
#include <string.h>
#include <netinet/in.h>
#include <sys/socket.h>
#include <unistd.h>

const char con_msg[] = "client connected!";
const char discon_msg[] = "client disconnected!";
struct sockaddr_in addr_in;

int init_client(const char *addr, int port) {
	int fd = socket(AF_INET, SOCK_DGRAM, IPPROTO_UDP);
	if (fd == -1) {
		fprintf(stderr, "Error setting up UDP socket\n");
		return -1;
	}
	printf("Client socket created successfully\n");

	memset(&addr_in, 0, sizeof(addr_in));

	addr_in.sin_addr.s_addr = inet_addr(addr);
	addr_in.sin_family = AF_INET;
	addr_in.sin_port = htons(port);

	if (sendto(fd, con_msg, strlen(con_msg), 0, (struct sockaddr*)&addr_in, sizeof(addr_in)) < 0) {
		fprintf(stderr, "Unable to confirm client connection\n");
		return -1;
	}
	printf("Client connected to the backend!\n");

	return fd;
}

int receive_from_serv(int sockfd, Packet *p) {
	char buf[2048];
	socklen_t len = sizeof(addr_in);
	int n = recvfrom(sockfd, &buf, 2048 - 1, 0, (struct sockaddr*)&addr_in, &len);

	if (n < 0) {
		return -1;
	}
	buf[n] = '\0';

	n = parse_packet(buf, p);
	if (n == -1) {
		// fprintf(stderr, "Error receiving a packet: server shutdown\n");
		return -1;
	} else if (n == 1) {
		// fprintf(stderr, "Error parsing the packet json\n");
		return 1;
	}

	return 0;
}

void close_client(int sockfd) {
	if (sockfd > 0) {
		sendto(sockfd, discon_msg, strlen(discon_msg), 0, (struct sockaddr*)&addr_in, sizeof(addr_in));
		close(sockfd);
		printf("Client disconnected from the backend server!\n");
	}
}
