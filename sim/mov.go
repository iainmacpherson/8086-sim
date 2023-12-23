package main

import (
	"8086-sim/logger"
	"errors"
)

func decodeMOV(istream *DataStream, instr *Instruction) {
	logger.LogInf(NAME, "Decoding MOV instruction.")

	if Opcodes.MOV_RM_TF_R == (instr.FirstByte >> 2) {
		// First Byte
		instr.OPCODE = Opcodes.MOV_RM_TF_R
		instr.D = (instr.FirstByte >> 1) & 0b1
		instr.W = (instr.FirstByte) & 0b1
		// Second Byte
		instr.MOD = istream.PopBits(2)
		instr.REG = istream.PopBits(3)
		instr.RM = istream.PopBits(3)
		// Third Byte
		nr_displacement_bytes := instr.checkForDisplacementBytes()
		if nr_displacement_bytes == 1 {
			instr.DISP_LO = istream.PopByte()
		}
		// Fourth Byte
		if nr_displacement_bytes == 2 {
			instr.DISP_HI = istream.PopByte()
		}
	} else if Opcodes.MOV_I_T_RM == (instr.FirstByte >> 1) {
		// First Byte
		instr.OPCODE = Opcodes.MOV_I_T_RM
		instr.W = (instr.FirstByte) & 0b1
		// Second Byte
		instr.MOD = istream.PopBits(2)
		instr.LIT3 = istream.PopBits(3)
		instr.RM = istream.PopBits(3)
		// Third Byte
		nr_displacement_bytes := instr.checkForDisplacementBytes()
		if nr_displacement_bytes == 1 {
			instr.DISP_LO = istream.PopByte()
		}
		// Fourth Byte
		if nr_displacement_bytes == 2 {
			instr.DISP_HI = istream.PopByte()
		}
		// Fifth Byte
		instr.DATA_LO = istream.PopByte()
		// Sixth Byte
		if instr.W == 1 {
			instr.DATA_HI = istream.PopByte()
		}
	} else if Opcodes.MOV_I_T_R == (instr.FirstByte >> 4) {
		// First Byte
		instr.OPCODE = Opcodes.MOV_I_T_R
		instr.W = (instr.FirstByte >> 3) & 0b1
		instr.REG = (instr.FirstByte) & 0b111
		// Second Byte
		instr.DATA_LO = istream.PopByte()
		// Third Byte
		if instr.W == 1 {
			instr.DATA_HI = istream.PopByte()
		}
	} else if Opcodes.MOV_M_T_A == (instr.FirstByte >> 1) {
		// First Byte
		instr.OPCODE = Opcodes.MOV_M_T_A
		instr.W = (instr.FirstByte) & 0b1
		// Second Byte
		instr.ADDR_LO = istream.PopByte()
		// Third Byte
		instr.ADDR_HI = istream.PopByte()
	} else if Opcodes.MOV_A_T_M == (instr.FirstByte >> 1) {
		// First Byte
		instr.OPCODE = Opcodes.MOV_A_T_M
		instr.W = (instr.FirstByte) & 0b1
		// Second Byte
		instr.ADDR_LO = istream.PopByte()
		// Third Byte
		instr.ADDR_HI = istream.PopByte()
	} else if Opcodes.MOV_RM_T_S == instr.FirstByte {
		// First Byte
		instr.OPCODE = Opcodes.MOV_RM_T_S
		// Second Byte
		instr.MOD = istream.PopBits(2)
		_ = istream.PopBits(1)
		instr.SR = istream.PopBits(2)
		instr.RM = istream.PopBits(3)
		// Third Byte
		nr_displacement_bytes := instr.checkForDisplacementBytes()
		if nr_displacement_bytes == 1 {
			instr.DISP_LO = istream.PopByte()
		}
		// Fourth Byte
		if nr_displacement_bytes == 2 {
			instr.DISP_HI = istream.PopByte()
		}
	} else if Opcodes.MOV_S_T_RM == instr.FirstByte {
		// First Byte
		instr.OPCODE = Opcodes.MOV_S_T_RM
		// Second Byte
		instr.MOD = istream.PopBits(2)
		_ = istream.PopBits(1)
		instr.SR = istream.PopBits(2)
		instr.RM = istream.PopBits(3)
		// Third Byte
		nr_displacement_bytes := instr.checkForDisplacementBytes()
		if nr_displacement_bytes == 1 {
			instr.DISP_LO = istream.PopByte()
		}
		// Fourth Byte
		if nr_displacement_bytes == 2 {
			instr.DISP_HI = istream.PopByte()
		}
	} else {
		logger.LogfErr(NAME, "Unknown MOV instruction, unable to decode. First byte = 0x%x", instr.FirstByte)
		panic(errors.New("IllegalInstruction"))
	}
}

func disassembleMOV(instr *Instruction) {
	logger.LogInf(NAME, "Disassembling MOV instruction.")
	switch instr.OPCODE {
	case Opcodes.MOV_RM_TF_R:
	case Opcodes.MOV_I_T_RM:
	case Opcodes.MOV_I_T_R:
	case Opcodes.MOV_M_T_A:
	case Opcodes.MOV_A_T_M:
	case Opcodes.MOV_RM_T_S:
	case Opcodes.MOV_S_T_RM:
	default:
		logger.LogfErr(NAME, "Unknown MOV instruction, unable to disassemble. First byte = 0x%x", instr.FirstByte)
		panic(errors.New("IllegalInstruction"))
	}
}

func executeMOV(instr *Instruction) {
	logger.LogfErr(NAME, "Attempting to execute unimplemented MOV instruction. First byte = 0x%x", instr.FirstByte)
	panic(errors.New("UnimplementedError"))
}
