package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	title  string
	output string
}

var InstructionFilePath = "instructions.txt"

var OutputFile = File{title: filepath.Base(InstructionFilePath), output: ""}

// ReadInstructions Read the instructions from the file
func ReadInstructions() (instructions [][]string) {

	fmt.Println("Select instruction file from current directory")
	// fmt.Scanln(&InstructionFilePath)
	InstructionFilePath = strings.Trim(InstructionFilePath, "'")

	file, err := os.Open(InstructionFilePath)
	if err != nil {
		fmt.Println("Error opening file")
		fmt.Println(InstructionFilePath)
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ";")
		instructions = append(instructions, line)
	}
	return
}

// OutputResultsFile Output the final results to a file
func OutputResultsFile(results string, number string) {
	if number == "" {
		return
	}
	OutputFile.output = results
	OutputFile.title = strings.TrimSuffix(OutputFile.title, filepath.Ext(OutputFile.title))
	file, err := os.Create(OutputFile.title + "_" + number + ".txt")
	if err != nil {
		fmt.Println("Error creating file")
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	_, err = file.WriteString(OutputFile.output)
	if err != nil {
		return
	}

}

// ProduceOutputString Output the current status of the memory in an intermediate output
func (f *File) ProduceOutputString(memory Memory, errors string) (output string) {
	//Allocated Blocks
	allocatedBlocks := make(map[int][]int)
	for i, v := range memory.mainMemory {
		if v != -1 {
			allocatedBlocks[v] = append(allocatedBlocks[v], i)
		}
	}
	startAddress := 0
	endAddress := 0
	for k, v := range allocatedBlocks {
		startAddress = v[0]
		endAddress = v[len(v)-1]
		output += fmt.Sprintf("Allocated Blocks\n %v;%v;%v\n", k, startAddress, endAddress)
	}
	freeBlocks := CalculateFreeBlocks(&memory)
	addresses := ""
	for _, v := range freeBlocks {
		count := 0
		addresses += fmt.Sprintf("%v;%v\n", v[count], v[count+1])
		count += 2
	}
	//fmt.Println(freeBlocks)
	output += fmt.Sprintf("Free Blocks:\n%v", addresses)

	//Errors
	if errors != "" {
		output += fmt.Sprintf("Errors:\n%v\n", errors)
	} else {
		output += "Errors: None\n"
	}

	//Fragmentation
	fragmentation := CalculateFragmentation(memory)
	output += fmt.Sprintf("Fragmentation: %v\n", fragmentation)
	output += "--------------------------------------------------\n"

	return
}

func CalculateFreeBlocks(memory *Memory) (freeBlocks [][]int) {
	freeBlocks = make([][]int, 0)
	var freeBlock []int
	for i, v := range memory.mainMemory {
		if v == -1 {
			if len(freeBlock) == 0 {
				freeBlock = append(freeBlock, i)
			}
		} else {
			if len(freeBlock) != 0 {
				freeBlock = append(freeBlock, i-1)
				freeBlocks = append(freeBlocks, freeBlock)
				freeBlock = make([]int, 0)
			}
		}
	}
	if len(freeBlock) != 0 {
		freeBlock = append(freeBlock, len(memory.mainMemory)-1)
		freeBlocks = append(freeBlocks, freeBlock)
	}

	return
}
