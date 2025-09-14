/* parse_json.h */

#pragma once

#include "../include/packets.h"
#include <string.h>
#include <stdio.h>

/* Parses a line of json into the Packet struct and returns the number of items parsed;
 * Returns -1 when servers shuts down;
 */
int parse_packet(const char *json, Packet *p);
