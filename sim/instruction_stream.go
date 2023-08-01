package main

import (
	"8086-sim/logger"
	"os"
)

type InstructionStream struct {
	length int
	bytePosition int
	binaryData []byte
}

func InstructionStreamCreate(filepath string) *InstructionStream {
	file_data, err := os.ReadFile(filepath)
	if err != nil {
		logger.LogfErr(NAME, "Error reading input program. Please ensure that the correct path was provided. Use -h for help.\n\n")
		panic(err)
	}

	istream := InstructionStream {
		length: len(file_data),
		bytePosition: 0,
		binaryData: file_data,
	}
	return &istream
}

func InstructionStreamPopByte(istream *InstructionStream) byte {
	var data uint8 = istream.binaryData[istream.bytePosition]
	istream.bytePosition++
	return data
}

func InstructionStreamIsEmpty(istream *InstructionStream) bool {
	return istream.bytePosition >= istream.length
}
