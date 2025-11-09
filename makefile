CC = gcc
CFLAGS = -I include
SRCS = main.c configParser.c utility.c
TARGET = wall

all:
	$(CC) $(SRCS) $(CFLAGS) -o $(TARGET)

clean:
	rm -f $(TARGET)
