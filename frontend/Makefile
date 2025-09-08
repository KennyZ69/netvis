CFLAGS = -Wall -Wextra -Iinclude -Idata/cJSON
LDFLAGS=

SRC = $(wildcard src/*.c) main.c data/cjson.c
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
	./$(OUT)
