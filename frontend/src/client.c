/* client.c */

#include "../include/client.h"
#include <string.h>
#include <netinet/in.h>
#include <sys/socket.h>

int init_client(const char *addr, int port) {
	struct sockaddr_in addr_in;

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

	char msg[] = "client connected!";
	if (sendto(fd, msg, strlen(msg), 0, (struct sockaddr*)&addr_in, sizeof(addr_in)) < 0) {
		fprintf(stderr, "Unable to confirm client connection\n");
		return -1;
	}
	printf("Client connected to the backend!\n");

	return fd;
}

int receive_from_serv(int sockfd, struct sockaddr_in servaddr, Packet *p) {

	return 0;
}
