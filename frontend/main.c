/* main.c */

#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include "include/app.h"
#include "include/client.h"
#include "include/packets.h"

static volatile int running = 1;

void handle_exit(int sig) {
	(void)sig;
	printf("\nShutting down program ...\n");
	running = 0;
}

void usage() {
	printf("Not enough arguments!\nRun:\n\t./netvis <address> <port>\n");
	exit(-1);
}

int main(int argc, char *argv[]) {
	if (argc < 3) usage();

	const char *addr = argv[1];
	int port = atoi(argv[2]);

	signal(SIGINT, handle_exit);

	int sockfd = init_client(addr, port);
	if (sockfd == -1) {
		exit(EXIT_FAILURE);
	}

	Packet p;
	while (running) {
		int stat = receive_from_serv(sockfd, &p);
		if (stat == 0) {
			printf("Received a packet:\n%s:%d -> %s:%d | %s | len=%d\n", p.src_ip, p.src_port, p.dst_ip, p.dst_port, p.prot, p.len);
		} else if (stat == -1) {
			printf("The server shut down...\n");
			break;
		} else {
			// fprintf(stderr, "Error receiving a packet\n");
			// break;
		}
	}

	close_client(sockfd);
	return 0;
}
