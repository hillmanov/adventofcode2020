package main

import (
	"fmt"
)

func main() {
	bootCodeProgram := ReadBootCodeProgram("./input.txt")

	// Part 1
	bcc := NewBootCodeComputer(CopyBootCodeProgram(bootCodeProgram))
	bcc.Run()
	fmt.Printf("Day 08: Part 1: = %+v\n", bcc.Accumulator)

	// Part 2
	for i, instruction := range bootCodeProgram {
		if instruction.Operation == "nop" || instruction.Operation == "acc" {
			continue
		}

		candidate := CopyBootCodeProgram(bootCodeProgram)
		if instruction.Operation == "jmp" {
			candidate[i].Operation = "nop"
		}
		if instruction.Operation == "nop" {
			candidate[i].Operation = "jmp"
		}

		bcc := NewBootCodeComputer(candidate)
		if bcc.Run() == 0 {
			fmt.Printf("Day 08: Part 2: = %+v\n", bcc.Accumulator)
			return
		}
	}
}
