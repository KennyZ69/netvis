/* parse_json.c */

#include "parse_json.h"
#include <stdio.h>
#include <string.h>

int simple_parse_packet(const char *json, Packet *p) {
	printf("Gotten json: %s\n", json);
	if (strstr(json, "\"shutdown\"") != NULL) {
		return -1; // server is shutting down
	}

	int n = sscanf(json, "{\"src_ip\":\"%63[^\"]\",\"dst_ip\":\"%63[^\"]\",\"src_port\":%d,"
        "\"dst_port\":%d,\"length\":%d,\"proto\":\"%15[^\"]\"}", p->src_ip, p->dst_ip, &p->src_port, &p->dst_port, &p->len, p->prot);

	return (n == 6) ? n : 0;
}

int parse_packet(const char *json, Packet *p) {
	printf("Gotten json: %s\n", json);
	if (strstr(json, "\"shutdown\"") != NULL) {
		return -1; // server is shutting down
	}

	int ok = 0;
	ok &= extract_string(json, "src_ip", p->src_ip, sizeof(p->src_ip));
	ok &= extract_string(json, "dst_ip", p->dst_ip, sizeof(p->dst_ip));

	return ok;
}

static int extract_string(const char *json, const char *key, char *out, size_t out_size) {
	char pattern[64];
	snprintf(pattern, sizeof(pattern), "\"%s\":\"", key);

	char *start = strstr(json, pattern);
	if (!start) return 1; // the key was not found in the line of json
	
	start += strlen(pattern); // moving after the '"'
	char *end = strchr(start, '"');
	if (!end) return 1;

	size_t len = end - start;
	if (len >= out_size) len = out_size -1;

	strncpy(out, start, len);
	out[len] = '\0';

	return 0;
}
