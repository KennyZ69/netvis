# CFLAGS = -Wall -Wextra -Iinclude -Idata/cJSON
CFLAGS = -Wall -Wextra -Iinclude -Idata
LDFLAGS=

# SRC = $(wildcard src/*.c) main.c data/cjson.c
SRC = $(wildcard src/*.c) main.c data/parse_json.c
OBJ = $(patsubst %.c,%.o,$(SRC))
OUT = netvis

all: $(OUT)

$(OUT): $(OBJ)
	gcc $(OBJ) -o $(OUT) $(LDFLAGS)

%.o: %.c
	gcc $(CFLAGS) -c $< -o $@

clean:
	rm -f $(OUT) $(OBJ)

run:
	./$(OUT) 127.0.0.1 420
#	./$(OUT)
