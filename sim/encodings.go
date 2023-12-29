package main

type opcodes_t struct {
	MOV_RM_TF_R uint8
	MOV_I_T_RM  uint8
	MOV_I_T_R   uint8
	MOV_M_T_A   uint8
	MOV_A_T_M   uint8
	MOV_RM_T_S  uint8
	MOV_S_T_RM  uint8
}

/* Holds opcodes. Do not modify. */
var Opcodes = opcodes_t{
	MOV_RM_TF_R: 0b100010,
	MOV_I_T_RM:  0b1100011,
	MOV_I_T_R:   0b1011,
	MOV_M_T_A:   0b1010000,
	MOV_A_T_M:   0b1010001,
	MOV_RM_T_S:  0b10001110,
	MOV_S_T_RM:  0b10001100,
}

type registers_t struct {
	AL uint8
	CL uint8
	DL uint8
	BL uint8
	AH uint8
	CH uint8
	DH uint8
	BH uint8
	AX uint8
	CX uint8
	DX uint8
	BX uint8
	SP uint8
	BP uint8
	SI uint8
	DI uint8
}

/* Holds IDs of named registers. Do not modify. */
var Registers = registers_t{
	AL: 0b0000, CL: 0b0001, DL: 0b0010, BL: 0b0011,
	AH: 0b0100, CH: 0b0101, DH: 0b0110, BH: 0b0111,
	AX: 0b1000, CX: 0b1001, DX: 0b1010, BX: 0b1011,
	SP: 0b1100, BP: 0b1101, SI: 0b1110, DI: 0b1111,
}

// NOTE: these maps are good for looking up instructions and registers for debugging
//      but they are slow. Performant code should use the above structs for logic
//		wherever possible.

/* Mapping of Binary opcodes to instruction names */
var OpcodeMnemonics = map[uint8]string{
	Opcodes.MOV_RM_TF_R: "mov",
	Opcodes.MOV_I_T_RM:  "mov",
	Opcodes.MOV_I_T_R:   "mov",
	Opcodes.MOV_M_T_A:   "mov",
	Opcodes.MOV_A_T_M:   "mov",
	Opcodes.MOV_RM_T_S:  "mov",
	Opcodes.MOV_S_T_RM:  "mov",
}

/*
Mapping of Binary Register ids to register names.

regcode = concat(W, R/M)
*/
var RegNames = map[uint8]string{
	// W = 0
	Registers.AL: "al", Registers.CL: "cl", Registers.DL: "dl", Registers.BL: "bl",
	Registers.AH: "ah", Registers.CH: "ch", Registers.DH: "dh", Registers.BH: "bh",
	// W = 1
	Registers.AX: "ax", Registers.CX: "cx", Registers.DX: "dx", Registers.BX: "bx",
	Registers.SP: "sp", Registers.BP: "bp", Registers.SI: "si", Registers.DI: "di",
}

var AddressingRegs = map[uint8]string{
	0b000: "bx + si",
	0b001: "bx + di",
	0b010: "bp + si",
	0b011: "bp + di",
	0b100: "si",
	0b101: "di",
	0b110: "bp",
	0b111: "bx",
}

var SegRegNames = map[uint8]string{
	0b00: "es",
	0b01: "cs",
	0b10: "ss",
	0b11: "ds",
}
