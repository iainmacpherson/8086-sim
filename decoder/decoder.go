package decoder

import (
    "errors"
    "8086-sim/logger"
)

const NAME = "DECODE"

/* Extracted instruction fields */
type DecodedInstruction struct {
    Opcode      uint8
    SrcRegCode  uint8
    DesRegCode  uint8
}

func readByte(program []byte, byte_count int) (uint8, int) {
    var byte_read byte = program[byte_count]
    byte_count++
    return byte_read, byte_count
}

/*
    Decode the next instruction in the array of bytes passed.
    returns: number of bytes consumed.
*/
func DecodeInstruction(program []byte) (int, DecodedInstruction) {
    var byte_count int = 0
    var byte_read uint8 = 0
    var decoded_instr = DecodedInstruction{}
    // read byte
    byte_read, byte_count = readByte(program, byte_count)

    // decode based on opcode
    decoded_instr.Opcode = (byte_read >> 2) & 0b00111111
    switch (decoded_instr.Opcode) {
    case Opcodes.MOV_RM_TF_R:
        //var dbit uint8 = (byte_read >> 1) & 0b1 // NOTE: seemingly ignored in reg to reg
        var wbit uint8 = byte_read & 0b1

        // read next byte
        byte_read, byte_count = readByte(program, byte_count)

        var modbits uint8 = (byte_read >> 6) & 0b11
        // TODO(iain): currenly only reg to reg operations are supported
        switch modbits {
        case 0b11:  // reg to reg MOV
            var regbits uint8 = (byte_read >> 3) & 0b111;
            var rmbits  uint8 = byte_read & 0b111;
            decoded_instr.SrcRegCode = (wbit << 3) | regbits
            decoded_instr.DesRegCode = (wbit << 3) | rmbits
        default:
            panic(errors.New("UnimplementedError"))
        }

    default:
        panic(errors.New("UnimplementedError"))
    }

    logger.LogfInf(NAME, "Decoded instruction, consumed %d bytes",  byte_count)
    return byte_count, decoded_instr
}

