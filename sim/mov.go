package main

import (
	"8086-sim/logger"
	"errors"
	"fmt"
)

func decodeMOV(istream *DataStream, instr *Instruction) {
	logger.LogfInf(NAME, "Decoding MOV instruction. First Byte = 0x%x", instr.FirstByte)

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
		if nr_displacement_bytes >= 1 {
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
		if nr_displacement_bytes >= 1 {
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
		instr.LIT1 = istream.PopBits(1)
		instr.SR = istream.PopBits(2)
		instr.RM = istream.PopBits(3)
		// Third Byte
		nr_displacement_bytes := instr.checkForDisplacementBytes()
		if nr_displacement_bytes >= 1 {
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
		instr.LIT1 = istream.PopBits(1)
		instr.SR = istream.PopBits(2)
		instr.RM = istream.PopBits(3)
		// Third Byte
		nr_displacement_bytes := instr.checkForDisplacementBytes()
		if nr_displacement_bytes >= 1 {
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
		var mnemonic string = OpcodeMnemonics[instr.OPCODE]
		var reg string = RegNames[(instr.W<<3)|instr.REG]
		var memory_or_reg string
		if instr.MOD == 0b11 {
			// Reg Mode
			memory_or_reg = RegNames[(instr.W<<3)|instr.RM]
		} else {
			// Memory Mode
			memory_or_reg = instr.calculateEffectiveAddress()
		}
		if instr.D == 0 {
			// From reg
			instr.Disassembly = fmt.Sprintf("%s %s, %s", mnemonic, memory_or_reg, reg)
		} else {
			// To reg
			instr.Disassembly = fmt.Sprintf("%s %s, %s", mnemonic, reg, memory_or_reg)
		}
	case Opcodes.MOV_I_T_RM:
		if instr.LIT3 != 0b000 {
			logger.LogfErr(NAME, "Literal provided (0x%x) does not equal expected value (0x%x). Unknown MOV instruction, unable to disassemble. First byte = 0x%x", instr.LIT3, 0, instr.FirstByte)
			panic(errors.New("IllegalInstruction"))
		}
		var mnemonic string = OpcodeMnemonics[instr.OPCODE]
		var memory string = instr.calculateEffectiveAddress()
		var immediate string = fmt.Sprint(uint16(instr.DATA_HI)<<8 | uint16(instr.DATA_LO))
		instr.Disassembly = fmt.Sprintf("%s %s, %s", mnemonic, memory, immediate)
	case Opcodes.MOV_I_T_R:
		var mnemonic string = OpcodeMnemonics[instr.OPCODE]
		var reg string = RegNames[(instr.W<<3)|instr.REG]
		var immediate string = fmt.Sprint(uint16(instr.DATA_HI)<<8 | uint16(instr.DATA_LO))
		instr.Disassembly = fmt.Sprintf("%s %s, %s", mnemonic, reg, immediate)
	case Opcodes.MOV_M_T_A:
		var mnemonic string = OpcodeMnemonics[instr.OPCODE]
		var reg string = RegNames[instr.W<<3]
		var memory string = "[" + fmt.Sprint(uint16(instr.ADDR_HI)<<8|uint16(instr.ADDR_LO)) + "]"
		instr.Disassembly = fmt.Sprintf("%s %s, %s", mnemonic, reg, memory)
	case Opcodes.MOV_A_T_M:
		var mnemonic string = OpcodeMnemonics[instr.OPCODE]
		var reg string = RegNames[instr.W<<3]
		var memory string = "[" + fmt.Sprint(uint16(instr.ADDR_HI)<<8|uint16(instr.ADDR_LO)) + "]"
		instr.Disassembly = fmt.Sprintf("%s %s, %s", mnemonic, memory, reg)
	case Opcodes.MOV_RM_T_S:
		if instr.LIT1 != 0b0 {
			logger.LogfErr(NAME, "Literal provided (0x%x) does not equal expected value(0x%x). Unknown MOV instruction, unable to disassemble. First byte = 0x%x", instr.LIT1, 0, instr.FirstByte)
			panic(errors.New("IllegalInstruction"))
		}
		var mnemonic string = OpcodeMnemonics[instr.OPCODE]
		var memory_or_reg string
		var seg_reg string = SegRegNames[instr.SR]
		if instr.MOD == 0b11 {
			// Reg Mode
			memory_or_reg = RegNames[(instr.W<<3)|instr.RM]
		} else {
			// Memory Mode
			memory_or_reg = instr.calculateEffectiveAddress()
		}
		instr.Disassembly = fmt.Sprintf("%s %s, %s", mnemonic, seg_reg, memory_or_reg)
	case Opcodes.MOV_S_T_RM:
		if instr.LIT1 != 0b0 {
			logger.LogfErr(NAME, "Literal provided (0x%x) does not equal expected value(0x%x). Unknown MOV instruction, unable to disassemble. First byte = 0x%x", instr.LIT1, 0, instr.FirstByte)
			panic(errors.New("IllegalInstruction"))
		}
		var mnemonic string = OpcodeMnemonics[instr.OPCODE]
		var memory_or_reg string
		var seg_reg string = SegRegNames[instr.SR]
		if instr.MOD == 0b11 {
			// Reg Mode
			memory_or_reg = RegNames[(instr.W<<3)|instr.RM]
		} else {
			// Memory Mode
			memory_or_reg = instr.calculateEffectiveAddress()
		}
		instr.Disassembly = fmt.Sprintf("%s %s, %s", mnemonic, memory_or_reg, seg_reg)
	default:
		logger.LogfErr(NAME, "Unknown MOV instruction, unable to disassemble. First byte = 0x%x", instr.FirstByte)
		panic(errors.New("IllegalInstruction"))
	}
}

func executeMOV(instr *Instruction) {
	// TODO(iain): simulate instruction execution
	logger.LogfErr(NAME, "Attempting to execute unimplemented MOV instruction. First byte = 0x%x", instr.FirstByte)
	panic(errors.New("UnimplementedError"))
}
