package main

import (
	"errors"
	"flag"
	"os"

	"8086-sim/logger"
)

const NAME = "SIM"

const MIN_ARGS = 1
const MAX_ARGS = 2

func Disassemble(program_binary []byte) {
	// output the required directive to indicate the x86 flavour.
	logger.LogRaw("bits 16")
	cpu_running := true
	byte_index := 0
	program_size := len(program_binary)
	for cpu_running {
		// Decode
		bytes_consumed, decoded_instr :=
			DecodeInstruction(program_binary[byte_index:])
		byte_index += bytes_consumed

		// disassemble
		logger.LogfRaw("%s %s, %s",
			OpcodeMnemonics[decoded_instr.Opcode],
			RegNames[decoded_instr.DesRegCode],
			RegNames[decoded_instr.SrcRegCode])

		// Check if we have finished executing all instructions
		if byte_index >= program_size {
			cpu_running = false
		}
	}
}

func Execute(program_binary []byte) {
	panic(errors.New("Unimplemented"))
}

func main() {
	// get command line arguments
	args := parseArguments()
	logger.Initialise(args.verbosity)
	logger.LogInf(NAME, "8086-sim started, reading input program.")

	// read input program
	program_binary := readInputProgram(args.filepath)

	if args.disassemble {
		Disassemble(program_binary)
	} else {
		Execute(program_binary)
	}

	// clean up and terminate
	logger.LogInf(NAME, "8086-sim completed, terminating.")
}

type cmdlineArgs struct {
	filepath    string
	verbosity   int
	disassemble bool
}

/*
Parses the commandline arguments using the Flag package.
*/
func parseArguments() cmdlineArgs {
	filepathPtr := flag.String("in", "", "The filepath to the executable program binary")
	vLevelPtr := flag.Int("verbosity", 0, "The numeric indicator for the verbosity level required.\n 0: Errors only\n 1: Warnings\n 2: Debug messages\n 3: Information messages\n")
	vAllPtr := flag.Bool("v", false, "Enables all logging output at the high verbosity level.\n Overrides other verbosity settings.")
	disassemblePtr := flag.Bool("disassemble", false, "Outputs the disassembly of the input program.")

	flag.Parse()

	args := cmdlineArgs{}
	args.filepath = *filepathPtr
	if *vAllPtr {
		args.verbosity = logger.INFO
	} else {
		args.verbosity = *vLevelPtr
	}
	args.disassemble = *disassemblePtr

	return args
}

func readInputProgram(filepath string) []byte {
	file_data, err := os.ReadFile(filepath)
	if err != nil {
		logger.LogfErr(NAME, "Error reading input program. Please ensure that the correct path was provided. Use -h for help.\n\n")
		panic(err)
	}

	return file_data
}
