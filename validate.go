package main

func hasMemoryBlock(memory []int, memoryBlock int) bool {
  for _, i := range memory {
    if i == memoryBlock {
      return true
    }
  }
  return false
}

// Check for errors
func VerifyInstruction(operation bool, instruction []string) (failedInstruction string) {
  if !operation {
    return instruction[0] + ";" + instruction[1] + ";" + instruction[2] + "\n"
  }
  return
}
