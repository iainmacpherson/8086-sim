package main

import (
	"8086-sim/logger"
	"errors"
)

/* Extracted instruction fields */
type DecodedInstruction struct {
	Opcode     uint8
	SrcRegCode uint8
	DesRegCode uint8
}

func DecodeNextInstruction(istream *InstructionStream) DecodedInstruction {
	logger.LogInf(NAME, "Decoding Next Instruction")
	var decoded_instr = DecodedInstruction{}
	// read byte
	byte_read := InstructionStreamPopByte(istream)

	// decode based on opcode
	decoded_instr.Opcode = (byte_read >> 2) & 0b00111111
	switch decoded_instr.Opcode {
	case Opcodes.MOV_RM_TF_R:
		//var dbit uint8 = (byte_read >> 1) & 0b1 // NOTE: seemingly ignored in reg to reg
		var wbit uint8 = byte_read & 0b1

		// read next byte
		byte_read = InstructionStreamPopByte(istream)

		var modbits uint8 = (byte_read >> 6) & 0b11
		// TODO(iain): currenly only reg to reg operations are supported
		switch modbits {
		case 0b11: // reg to reg MOV
			var regbits uint8 = (byte_read >> 3) & 0b111
			var rmbits uint8 = byte_read & 0b111
			decoded_instr.SrcRegCode = (wbit << 3) | regbits
			decoded_instr.DesRegCode = (wbit << 3) | rmbits
		default:
			panic(errors.New("UnimplementedError"))
		}

	default:
		panic(errors.New("UnimplementedError"))
	}

	return decoded_instr
}
