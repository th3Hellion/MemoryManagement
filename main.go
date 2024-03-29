package main

import (
	"fmt"
	"strconv"
)

var OutputNumber = 1
var errors string
var outputString = &OutputFile.output

func main() {
	start()
}

func start() {
	instructions := ReadInstructions()
	MemorySize, _ := strconv.Atoi(instructions[0][0])

	MainMemory := Memory{mainMemory: make([]int, MemorySize)}
	MakeMemory(&MainMemory)
	//Cut off the first line
	instructions = instructions[1:]

	for i, instruction := range instructions {
		switch instruction[0][0] {
		case 'A':
			errors += Allocate(instruction, &MainMemory)

		case 'D':
			errors += Deallocate(instruction, &MainMemory)

		case 'C':
			MainMemory.Compact()

		case 'O':
			Output(outputString, MainMemory)

		default:
			fmt.Println("Error: Invalid instruction")
		}

		// Last instruction
		if len(instructions)-1 == i {
			Output(outputString, MainMemory)
			Fit++
			if Fit == END {
				break
			}
			// Do next fit
			start()

		}
	}
}
