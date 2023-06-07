#ifndef UTILS_H_
#define UTILS_H_

namespace utils
{

// Exit with failure if the condition is not met, printing a message to stderr
void checkCondition(bool condition, const char* message);

// Exit with failure printing a message to stderr.
void panic(const char* message);

// prints each bit of a byte
void printBinaryByte(char byte);

} // namespace utils

#endif /* UTILS_H_ */
