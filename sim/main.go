package main

import (
	"errors"
	"flag"

	ds "8086-sim/datastream"
	"8086-sim/logger"
)

const NAME = "SIM"

const MIN_ARGS = 1
const MAX_ARGS = 2

type cmdlineArgs struct {
	filepath    string
	verbosity   int
	disassemble bool
	execute     bool
}

func main() {
	// get command line arguments
	args := parseArguments()
	logger.Initialise(args.verbosity)
	logger.LogInf(NAME, "8086-sim started, reading input program.")

	// read input program
	istream := ds.DataStreamCreateFromFile(args.filepath)

	decoded_instructions := Decode(istream)

	if args.disassemble {
		printDisassembly(Disassemble(decoded_instructions))
	}
	if args.execute {
		Execute(decoded_instructions)
	}

	// clean up and terminate
	logger.LogInf(NAME, "8086-sim completed, terminating.")
}

/*
Parses the commandline arguments using the Flag package.
*/
func parseArguments() cmdlineArgs {
	filepathPtr := flag.String("in", "", "The filepath to the executable program binary")
	vLevelPtr := flag.Int("verbosity", 0, "The numeric indicator for the verbosity level required.\n 0: Errors only\n 1: Warnings\n 2: Debug messages\n 3: Information messages\n")
	vAllPtr := flag.Bool("v", false, "Enables all logging output at the high verbosity level.\n Overrides other verbosity settings.")
	disassemblePtr := flag.Bool("disassemble", true, "Outputs the disassembly of the input program.")
	executePtr := flag.Bool("execute", false, "Executes the input program [currently unsupported].")

	flag.Parse()

	args := cmdlineArgs{}
	args.filepath = *filepathPtr
	if *vAllPtr {
		args.verbosity = logger.INFO
	} else {
		args.verbosity = *vLevelPtr
	}
	args.disassemble = *disassemblePtr
	args.execute = *executePtr

	return args
}

func Decode(istream *ds.DataStream) []*Instruction {
	var decoded_instructions []*Instruction
	for !istream.IsEmpty() {
		instr := DecodeNextInstruction(istream)
		if instr != nil {
			decoded_instructions = append(decoded_instructions, instr)
		}
	}
	return decoded_instructions
}

func Disassemble(instructions []*Instruction) []string {
	var output []string
	for _, instr := range instructions {
		instr.Runnable.Disassemble(instr)
		output = append(output, instr.Disassembly)
	}
	return output
}

func printDisassembly(disassembly []string) {
	logger.LogRaw("bits 16")
	for _, line := range disassembly {
		logger.LogRaw(line)
	}
}

func Execute(instructions []*Instruction) {
	// TODO(iain)
	panic(errors.New("Unimplemented"))
}
