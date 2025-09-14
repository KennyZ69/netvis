/* client.h */

#pragma once

#include "packets.h"
#include <stdio.h>
#include <unistd.h>
#include <sys/socket.h>
#include <arpa/inet.h>


/* Initialize the listener for given address on given port return the socket fd;
 * Returns -1 on errors;
 */
int init_client(const char *addr, int port);

/* Receives a network packet from the server and decodes it into provided Packet struct;
* Returns 0 on success, -1 on errors, 1 when server shuts down;
*/
int receive_from_serv(int sockfd, Packet *p);

void close_client(int sockfd);
