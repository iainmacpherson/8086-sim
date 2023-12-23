package main

import (
	"errors"
	"flag"

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

func Disassemble(istream *DataStream) {
	// output the required directive to indicate the x86 flavour.
	logger.LogRaw("bits 16")
	cpu_running := true
	for cpu_running {
		// Decode
		decoded_instr := DecodeNextInstruction(istream)

		if decoded_instr != nil {
			// disassemble
			decoded_instr.Runnable.Disassemble(decoded_instr)
		}
		if istream.IsEmpty() {
			// Check if we have finished executing all instructions
			cpu_running = false
		}
	}
}

func Execute(istream *DataStream) {
	// TODO(iain)
	panic(errors.New("Unimplemented"))
}

func main() {
	// get command line arguments
	args := parseArguments()
	logger.Initialise(args.verbosity)
	logger.LogInf(NAME, "8086-sim started, reading input program.")

	// read input program
	istream := DataStreamCreate(args.filepath)

	if args.disassemble {
		Disassemble(istream)
	}
	if args.execute {
		Execute(istream)
	}

	// clean up and terminate
	logger.LogInf(NAME, "8086-sim completed, terminating.")
}
