package main

import (
	"fmt"
	"strconv"
)

type Memory struct {
	mainMemory []int
}

func Allocate(instruction []string, MainMemory *Memory) (errors string) {
	memoryBlock, _ := strconv.Atoi(instruction[1])
	size, _ := strconv.Atoi(instruction[2])

	switch Fit {
	case FIRST_FIT:
		errors += VerifyInstruction(MainMemory.FirstFit(memoryBlock, size), instruction)
	case BEST_FIT:
		errors += VerifyInstruction(MainMemory.BestFit(memoryBlock, size), instruction)
	case WORST_FIT:
		errors += VerifyInstruction(MainMemory.WorstFit(memoryBlock, size), instruction)
	}
	return
}

func Deallocate(instruction []string, MainMemory *Memory) (errors string) {
	memoryBlock, _ := strconv.Atoi(instruction[1])

	if !MainMemory.Remove(memoryBlock) {
		errors += instruction[0] + ";" + instruction[1] + "\n"
	}
	return
}

// Remove from array
func (m *Memory) Remove(memoryBlock int) (success bool) {
	if !hasMemoryBlock(m.mainMemory, memoryBlock) {
		return false
	}
	for _, v := range m.mainMemory {
		if v == memoryBlock {
			break
		}
	}

	for i := range m.mainMemory {
		if m.mainMemory[i] == memoryBlock {
			m.mainMemory[i] = -1
		}
	}
	return true
}

func (m *Memory) Compact() {
	var tempArray []int
	for _, v := range m.mainMemory {
		if v != -1 {
			tempArray = append(tempArray, v)
		}
	}
	for i := len(tempArray); i < len(m.mainMemory); i++ {
		tempArray = append(tempArray, -1)
	}
	m.mainMemory = tempArray
}

func Output(outputString *string, MainMemory Memory) {
	*outputString += FitToString() + "\n"
	*outputString += OutputFile.ProduceOutputString(MainMemory, errors)
	fmt.Println(*outputString)
	OutputResultsFile(*outputString, strconv.Itoa(OutputNumber))
	OutputNumber++
}
