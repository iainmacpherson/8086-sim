#include "decoder.h"

#include "utils.h"

namespace control
{
    namespace ops
    {
        const char* mnemonics[] = {
            "x","x","x","x",
            "x","x","x","x",
            "x","x","x","x",
            "x","x","x","x",
            "x","x","x","x",
            "x","x","x","x",
            "x","x","x","x",
            "x","x","x","x",
            "x","x","mov","x",
            "x","x","x","x",
            "x","x","x","x",
            "x","x","x","x",
            "x","x","x","x",
            "x","x","x","x",
            "x","x","x","x",
            "x","x","x","x"};
    }
    namespace regs
    {
        const char* regnames[] = {"al", "cl", "dl", "bl", "ah", "ch", "dh", "bh",
            "ax", "cx", "dx", "bx", "sp", "bp", "si", "di"};
    } // namespace regs

    int decodeInstruction(char* instruction_stream, instruction_t& decoded_instr)
    {
        char* initial_stream_position = instruction_stream;
        // read byte
        char byte = *instruction_stream++;

        // grab opcode
        decoded_instr.opcode = (control::ops::opcode_t)((byte >> 2) & 0b00111111);
        switch (decoded_instr.opcode)
        {
            case ops::MOV:
            {
                char dbit = (byte >> 1) & 0b1; // NOTE: seemingly ignored in reg to reg?
                char wbit = byte & 0b1;

                // read next byte
                byte = *instruction_stream++;
                char modbits = (byte >> 6) & 0b11;
                char regbits = (byte >> 3) & 0b111;
                char rmbits = byte & 0b111;

                // currently only support reg to reg operations.
                if (modbits != 0b11) { utils::panic("Not implemented yet."); }

                decoded_instr.src_reg =
                    (control::regs::regcode_t)((wbit << 3) | regbits);
                decoded_instr.dest_reg =
                    (control::regs::regcode_t)((wbit << 3) | rmbits);
            } break;
            // TODO(iain): Implement the remaining opcodes
            default:
            {
                utils::panic("Provided opcode not implemented.");
            } break;
        }
        int bytes_consumed = instruction_stream - initial_stream_position;
        return bytes_consumed;
    }
} // namespace control

