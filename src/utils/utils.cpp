#include "utils.h"

#include <stdio.h>
#include <stdlib.h>

namespace utils
{

void checkCondition(bool condition, const char* message)
{
    if (condition)
        return;
    fprintf(stderr, "ERROR: %s\n", message);
    exit(EXIT_FAILURE);
}

void panic(const char* message)
{
    fprintf(stderr, "ERROR: %s\n", message);
    exit(EXIT_FAILURE);
}

#define BYTE_TO_BINARY_PATTERN "%c%c%c%c%c%c%c%c"
#define BYTE_TO_BINARY(byte)  \
  ((byte) & 0x80 ? '1' : '0'), \
  ((byte) & 0x40 ? '1' : '0'), \
  ((byte) & 0x20 ? '1' : '0'), \
  ((byte) & 0x10 ? '1' : '0'), \
  ((byte) & 0x08 ? '1' : '0'), \
  ((byte) & 0x04 ? '1' : '0'), \
  ((byte) & 0x02 ? '1' : '0'), \
  ((byte) & 0x01 ? '1' : '0')

void printBinaryByte(char byte)
{
    printf("0b" BYTE_TO_BINARY_PATTERN "\n", BYTE_TO_BINARY(byte));
}


} // namespace utils
