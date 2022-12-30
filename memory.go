package main

func CalculateFragmentation(memory Memory) (fragmentation float32) {
	var totalFreeMemory float32
	var largestFreeBlock float32
	for _, v := range memory.mainMemory {
		if v == -1 {
			totalFreeMemory++
		}
	}
	counter := 0
	for _, v := range memory.mainMemory {
		if v == -1 {
			counter++
		} else {
			if float32(counter) > largestFreeBlock {
				largestFreeBlock = float32(counter)
			}
			counter = 0
		}
	}
	fragmentation = 1 - largestFreeBlock/totalFreeMemory
	fragmentation = float32(int(fragmentation*1000000)) / 1000000
	return
}

// Fill with -1 to indicate free memoryi (null)
func MakeMemory(memory *Memory) {
	for i := range memory.mainMemory {
		memory.mainMemory[i] = -1
	}
}
