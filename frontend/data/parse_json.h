/* parse_json.h */

#pragma once

#include "../include/packets.h"
#include <string.h>
#include <stdio.h>

/* Parses a line of json into the Packet struct and returns the number of items parsed;
 * Returns -1 when servers shuts down;
 */
int simple_parse_packet(const char *json, Packet *p);

/* Parses a line of json and finds values to hardcoded keys and give them to the packet struct;
 * Returns -1 when server shuts down, 1 on error and 0 on success;
 */
int parse_packet(const char *json, Packet *p);

/* Looks for a string value under given key in the json and extracts it into out;
 * Returns 0 on sucess, 1 on error;
 */
static int extract_string(const char *json, const char *key, char *out, size_t out_size);

/* Looks for a int value under given key in the json and extracts it into out;
 * Returns 0 on sucess, 1 on error;
 */
static int extract_int(const char *json, const char *key, int *out);
