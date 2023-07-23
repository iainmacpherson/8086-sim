package main

import (
	"flag"
	"os"

	"8086-sim/decoder"
	"8086-sim/logger"
)

const NAME = "MAIN"
const MIN_ARGS = 1
const MAX_ARGS = 2

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// get command line arguments
	args := parseArguments()
	logger.Initialise(args.verbosity)
	logger.LogInf(NAME, "8086-sim started, reading input program.")

	// read input program
	program_binary := readInputProgram(args.filepath)

	// program execution loop
	cpu_running := true
	byte_index := 0
	program_size := len(program_binary)
	for cpu_running {
		// Decode
		bytes_consumed, decoded_instr :=
			decoder.DecodeInstruction(program_binary[byte_index:])
		byte_index += bytes_consumed

		// TEMP: output the disassembled instruction
		logger.LogfRaw("%s %s, %s",
			decoder.OpcodeMnemonics[decoded_instr.Opcode],
			decoder.RegNames[decoded_instr.DesRegCode],
			decoder.RegNames[decoded_instr.SrcRegCode])

		// Check if we have finished executing all instructions
		if byte_index >= program_size {
			cpu_running = false
		}
	}

	// clean up and terminate
	logger.LogInf(NAME, "8086-sim completed program execution, terminating.")
}

type cmdlineArgs struct {
	filepath  string
	verbosity int
}

/*
	Parses the commandline arguments using the Flag package.
*/
func parseArguments() cmdlineArgs {
	filepathPtr := flag.String("in", "", "The filepath to the executable program binary")
	vLevelPtr := flag.Int("verbosity", 0, "The numeric indicator for the verbosity level required.\n 0: Errors only\n 1: Warnings\n 2: Debug messages\n 3: Information messages\n")
	vAllPtr := flag.Bool("v", false, "Enables all logging output at the high verbosity level.\n Overrides other verbosity settings.")

	flag.Parse()

	args := cmdlineArgs{}
	args.filepath = *filepathPtr
	if *vAllPtr {
		args.verbosity = logger.INFO
	} else {
		args.verbosity = *vLevelPtr
	}

	return args
}

func readInputProgram(filepath string) []byte {
	file_data, err := os.ReadFile(filepath)
	if err != nil {
		logger.LogfErr(NAME, "Error reading input program. Please ensure that the correct path was provided. Use -h for help.\n\n" + err.Error())
		os.Exit(1)
	}

	return file_data
}
