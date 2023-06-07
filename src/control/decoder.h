#ifndef DECODER_H_
#define DECODER_H_

namespace control
{
    namespace ops
    {
        typedef enum
        {
            MOV = 0b100010
        } opcode_t;
        extern const char* mnemonics[]; // for looking up the opcode mnemonics
    } // namespace ops

    namespace regs
    {
        /* regcode = concat(W, R/M) */
        typedef enum
        {
            // W = 0
            AL = 0b0000, CL = 0b0001, DL = 0b0010, BL = 0b0011,
            AH = 0b0100, CH = 0b0101, DH = 0b0110, BH = 0b0111,
            // W = 1
            AX = 0b1000, CX = 0b1001, DX = 0b1010, BX = 0b1011,
            SP = 0b1100, BP = 0b1101, SI = 0b1110, DI = 0b1111,
        } regcode_t;
        extern const char* regnames[]; // for looking up the register names
    } // namespace regs

    typedef struct
    {
        ops::opcode_t opcode;
        regs::regcode_t src_reg;
        regs::regcode_t dest_reg;
    } instruction_t;

    /*
        returns: number of bytes consumed from the stream.
    */
    int decodeInstruction(char* instruction_stream, instruction_t& decoded_instruction);
} // namespace control

#endif /* DECODER_H_ */
