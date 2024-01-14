# 8086 Sim

A simulation of a basic 8086 processor written in Go, implementing a
subset of the processor's features.

Written for fun and practice learning go.

## Current state of the project
Currently, this project is only capable of decoding and outputting a disassembly of a binary containing a limited subset of the instructions supported by the 8086, and it does not support execution of any.
### Supported Instructions

| Instruction | Decoding | Disassembly | Execution | Description |
| --- | :---: | :---: | :---: | --- |
| MOV | y | y | n | Transfer a byte or word from source to destination. |
| PUSH | n | n | n | Decrements the stack pointer and transfers a word from the source to the stack. |

### Current TODO list
* implement disassembly for push
    * disambiguate properly between operations with the same first byte
        * need to implement a peak function in DataStream

### Future Plans
This is very much in a proof of concept phase at the moment.
I may eventually have the time to implement support for the full 8086 instruction set but for now I am happy to have written a framework that would make it relatively simple (if time consuming) to do so.

## Requirements
* Go
* Make

## Building
1. Clone this repository
2. Run `make` in the root directory to build the project.
3. The binary is in the `bin` directory

## Running the project
There are a few sample 8086 binaries in the `data` folder along with their source code. They were assembled with NASM.

To run the binary you will need to provide a path to the binary to run. For example you could run this from the root directory:
```
./bin/8086-sim -disassemble -in data/many_mov
```
To access the help run the binary with the argument `-h`.

## Credits
* Project inspired by C. Muratori's [Performance Aware Programming Course](https://www.computerenhance.com/). However, all code and techniques are my own.