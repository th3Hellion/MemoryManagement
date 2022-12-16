package main

import "sort"

const (
	FIRST_FIT = iota
	BEST_FIT
	WORST_FIT
	END
)

// Fit Start with first fit
var Fit = FIRST_FIT

func (m *Memory) FirstFit(memoryBlock int, size int) (success bool) {
	counter := 0
	hole := -1
	searchRange := len(m.mainMemory)
	for i := 0; i < searchRange; i++ {
		if m.mainMemory[i] != -1 {
			counter = 0
		} else {
			counter++
		}
		if counter == size {
			hole = i - size + 1
			break
		}
	}
	if hole == -1 {
		return false
	} else {
		for i := hole; i < hole+size; i++ {
			m.mainMemory[i] = memoryBlock
		}
		return true
	}

}

func (m *Memory) BestFit(memoryBlock int, size int) (success bool) {
	counter := 0
	holes := make([]int, 0)
	searchRange := len(m.mainMemory)
	for i := 0; i < searchRange; i++ {
		if m.mainMemory[i] != -1 {
			counter = 0
		} else {
			counter++
		}
		if counter == size {
			holes = append(holes, i-size+1)
		}
	}
	smallestHole := sort.IntSlice(holes)
	smallestHole.Sort()
	if len(smallestHole) == 0 {
		return false
	} else {

		for i := smallestHole[0]; i < smallestHole[0]+size; i++ {
			m.mainMemory[i] = memoryBlock
		}
		return true
	}
}

func (m *Memory) WorstFit(memoryBlock int, size int) (success bool) {
	counter := 0
	holes := make([]int, 0)
	searchRange := len(m.mainMemory)
	for i := 0; i < searchRange; i++ {
		if m.mainMemory[i] != -1 {
			counter = 0
		} else {
			counter++
		}
		if counter == size {
			holes = append(holes, i-size+1)
		}
	}
	largestHole := sort.IntSlice(holes)
	largestHole.Sort()
	if len(largestHole) == 0 {
		return false
	} else {
		for i := largestHole[len(largestHole)-1]; i < largestHole[len(largestHole)-1]+size; i++ {
			m.mainMemory[i] = memoryBlock
		}
		return true
	}
}
func FitToString() string {
	switch Fit {
	case FIRST_FIT:
		return "First Fit"
	case BEST_FIT:
		return "Best Fit"
	case WORST_FIT:
		return "Worst Fit"
	default:
		return "Unknown"
	}
}
