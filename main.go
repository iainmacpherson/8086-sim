package main

import (
	"errors"
	"os"
	"strconv"

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
	args := parseArguments(os.Args[1:])
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

func parseArguments(provided_args []string) cmdlineArgs {
	// TODO(iain): improve command line argument handling
	args := cmdlineArgs{}

	// Check if we have the correct arguments
	if len(provided_args) < MIN_ARGS || len(provided_args) > MAX_ARGS {
		logger.LogErr(NAME, "Provide a minimum of one argument: the path to the binary")
		panic(errors.New("IncorrectArgs"))
	}

	// parse the filepath
	args.filepath = provided_args[0]

	// parse the verbosity level
	if len(provided_args) > 1 {
		if provided_args[1] == "-v" {
			args.verbosity = logger.INFO
		} else if len(provided_args[1]) > 2 && provided_args[1][0:2] == "-v" {
			args.verbosity, _ = strconv.Atoi(string(provided_args[1][2]))
		} else {
			logger.LogErr(NAME, "Unknown additional argument provided, ignored.")
		}
	}

	return args
}

func readInputProgram(filepath string) []byte {
	file_data, err := os.ReadFile(filepath)
	check(err)

	return file_data
}
