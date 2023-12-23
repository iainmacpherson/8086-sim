package main

import (
	"8086-sim/logger"
)

/* Extracted instruction fields */
type Instruction struct {
	FirstByte uint8                // indicates instruction operation - byte
	OPCODE    uint8                // Instruction opcode - variable bit length
	LIT3      uint8                // literal - 3 bits
	LIT8      uint8                // literal - 8 bits
	MOD       uint8                // mode - 2 bits
	REG       uint8                // register - 3 bits
	RM        uint8                // reg/mem - 3 bits
	SR        uint8                // Segment reg code - 2 bits
	W         uint8                // wide - 1 bit
	S         uint8                // - 1 bit
	D         uint8                // direction - 1 bit
	V         uint8                // - 1 bit
	Z         uint8                // - 1 bit
	DISP_LO   uint8                // low byte of optional 8 or 16 bit displacement - byte
	DISP_HI   uint8                // low byte of optional 16 bit displacement - byte
	DATA_8    uint8                // immediate 8 bit const - byte
	DATA_LO   uint8                // low byte of 16 bit const - byte
	DATA_HI   uint8                // high byte of 16 bit const - byte
	DATA_SX   uint8                // immediate, sign extended to 16 bits before use - byte
	IP_INC_8  uint8                // signed increment to Instruction Pointer (IP) - byte
	IP_INC_LO uint8                // low byte of signed 16 bit IP increment - byte
	IP_INC_HI uint8                // high byte of signed 16 bit IP increment - byte
	IP_LO     uint8                // low byte of new IP value - byte
	IP_HI     uint8                // high byte of new IP value - byte
	CS_LO     uint8                // low byte of new CS value - byte
	CS_HI     uint8                // high byte of new CS value - byte
	SEG_LO    uint8                // low byte of seg value - byte
	SEG_HI    uint8                // high byte of seg value - byte
	ADDR_LO   uint8                // low byte of direct address - byte
	ADDR_HI   uint8                // high byte of direct address - byte
	Runnable  InstructionFunctions // function pointers providing required methods
}

type InstructionFunctions struct {
	DecodeFields func(*DataStream, *Instruction)
	Disassemble  func(*Instruction) // call to disassemble this instruction
	Execute      func(*Instruction) // call to execute this instruction
}

func DecodeNextInstruction(istream *DataStream) *Instruction {
	if istream.IsEmpty() {
		return nil
	}
	logger.LogInf(NAME, "Decoding Next Instruction")
	var instr = Instruction{}

	// decode first byte
	instr.FirstByte = istream.PopByte()
	// Use first byte to index a table of function pointers
	instr.Runnable = InstructionLookup[instr.FirstByte]
	// Decode the data in the rest of the instruction
	instr.Runnable.DecodeFields(istream, &instr)

	return &instr
}

func (instr *Instruction) checkForDisplacementBytes() uint {
	nr_displacement_bytes := 0
	if instr.MOD == 0b01 {
		nr_displacement_bytes = 1
	} else if instr.MOD == 0b10 {
		nr_displacement_bytes = 2
	} else if instr.MOD == 0b00 && instr.RM == 0b110 {
		nr_displacement_bytes = 2
	}

	return uint(nr_displacement_bytes)
}
