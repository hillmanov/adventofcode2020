package main

import (
	"adventofcode/utils"
	"fmt"
)

type Instruction struct {
	Operation string
	Argument  int
	Executed  bool
}

type BootCodeProgram []*Instruction

type BootCodeComputer struct {
	Accumulator int
	Program     BootCodeProgram
	Index       int
}

func NewBootCodeComputer(p BootCodeProgram) *BootCodeComputer {
	return &BootCodeComputer{
		Accumulator: 0,
		Program:     p,
		Index:       0,
	}
}

func (bcc *BootCodeComputer) Run() int {
	for {
		if bcc.Index >= len(bcc.Program) {
			return 0
		}

		instruction := bcc.Program[bcc.Index]

		if instruction.Executed {
			return 1
		}

		switch instruction.Operation {
		case "acc":
			bcc.Accumulator += instruction.Argument
			bcc.Index++
		case "jmp":
			bcc.Index += instruction.Argument
		case "nop":
			bcc.Index++
		}
		instruction.Executed = true
	}
}

func ReadBootCodeProgram(filename string) BootCodeProgram {
	input, err := utils.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	p := BootCodeProgram{}
	for _, line := range input {
		instruction := Instruction{}
		fmt.Sscanf(line, "%s %d", &instruction.Operation, &instruction.Argument)
		p = append(p, &instruction)
	}

	return p
}

func CopyBootCodeProgram(source BootCodeProgram) BootCodeProgram {
	target := make([]*Instruction, len(source))
	for i, instruction := range source {
		target[i] = &Instruction{
			Operation: instruction.Operation,
			Argument:  instruction.Argument,
			Executed:  instruction.Executed,
		}
	}
	return target
}
