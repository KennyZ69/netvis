/* parse_json.c */

#include "parse_json.h"

int parse_packet(const char *json, Packet *p) {
	if (strstr(json, "\"shutdown\"") != NULL) {
		return -1; // server is shutting down
	}

	int n = sscanf(json, "{\"src_ip\":\"%63[^\"]\",\"dst_ip\":\"%63[^\"]\",\"src_port\":%d,"
        "\"dst_port\":%d,\"length\":%d,\"proto\":\"%15[^\"]\"}", p->src_ip, p->dst_ip, &p->src_port, &p->dst_port, &p->len, p->prot);

	return (n == 6) ? n : 0;
}
