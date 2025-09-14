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

