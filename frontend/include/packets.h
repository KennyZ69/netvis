/* packets.h */


#pragma once

typedef struct {
	int len;
	char src_ip[40]; // possible max 39 characters for ipv6
	char dst_ip[40];
	char prot[16];
} Packet
