/* main.c */

#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include "include/app.h"
#include "include/client.h"

static volatile int running = 1;

void handle_exit(int sig) {
	(void)sig;
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

	return 0;
}
