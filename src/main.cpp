#include "decoder.h"

#include "utils.h"

#include <stdlib.h>
#include <stdio.h>

int main(int argc, char** argv)
{
    /* Program Preparation */

    // check correct number of args
    utils::checkCondition(argc == 2, "Provide one argument: the path to the binary.");
    // grab the binary file path and open the file
    char* filepath = argv[1];
    FILE* fileptr = fopen(filepath, "rb");
    utils::checkCondition(fileptr, "Unable to open provided file.");
    // work out the length of the file from the offset of the end.
    utils::checkCondition(fseek(fileptr, 0, SEEK_END) == 0, "Failed seeking the file end.");
    long filelen = ftell(fileptr);
    utils::checkCondition(filelen != -1, "Failed to get the position of the end of the file.");
    rewind(fileptr);
    // allocate memory for the file contents
    char* binary_buffer = (char*)malloc(filelen * sizeof(char));
    // read the file into the binary_buffer
    fread(binary_buffer, filelen, 1, fileptr);
    // done with the file so close it.
    fclose(fileptr);

    /* Program Execution Loop */
    // runs until the end of the buffer consuming bytes as needed.
    char* positionptr = binary_buffer;
    while (positionptr != binary_buffer + filelen*sizeof(char))
    {
        // Decode an instruction.
        // increment the current position by the number of bytes consumed.
        control::instruction_t decoded_instr = {};
        positionptr += control::decodeInstruction(positionptr, decoded_instr);

        // For testing: print out the instruction.
        printf("%s %s, %s\n", control::ops::mnemonics[decoded_instr.opcode],
                control::regs::regnames[decoded_instr.dest_reg],
                control::regs::regnames[decoded_instr.src_reg]);

        // TODO(iain): Dispatch the instruction for execution.
    }

    free(binary_buffer);
    return EXIT_SUCCESS;
}


