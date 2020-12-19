package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
)

type any interface{}

type token int

func (t token) IsOperator() bool {
	return t == Add || t == Multiply
}

const (
	Space      token = 0
	Number     token = 1
	Add        token = 2
	Multiply   token = 3
	OpenParen  token = 4
	CloseParen token = 5
)

func tokenType(r rune) token {
	switch r {
	case ' ':
		return Space
	case '+':
		return Add
	case '*':
		return Multiply
	case '(':
		return OpenParen
	case ')':
		return CloseParen
	default:
		return Number
	}
}

func main() {
	expressions, _ := utils.ReadLines("./input.txt")

	part1Solution := part1(expressions)
	part2Solution := part2()

	fmt.Printf("Day 18: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 18: Part 2: = %+v\n", part2Solution)
}

func part1(expressions []string) int {
	sum := 0
	for _, expression := range expressions {
		sum += evaluateExpression(expression)
	}
	return sum
}

func part2() int {
	return -1
}

func evaluateExpression(e string) int {
	stack := []any{}

	for _, r := range e {
		switch tokenType(r) {
		case Number:
			num, _ := strconv.Atoi(string(r))
			stack = append(stack, num)
			if len(stack) >= 3 {
				if value, ok := peek(stack, 1).(rune); ok && tokenType(value).IsOperator() {
					var operandB, operation, operandA any
					operandB, stack = pop(stack)
					operation, stack = pop(stack)
					operandA, stack = pop(stack)

					switch tokenType(operation.(rune)) {
					case Add:
						stack = append(stack, operandA.(int)+operandB.(int))
					case Multiply:
						stack = append(stack, operandA.(int)*operandB.(int))
					}
				}
			}

		case OpenParen:
			stack = append(stack, r)
		case CloseParen:
			var value any
			stack = append(stack, r)
			value, stack = popParenResult(stack)
			stack = append(stack, value)

			if len(stack) >= 3 {
				if value, ok := peek(stack, 1).(rune); ok && tokenType(value).IsOperator() {
					var operandB, operation, operandA any
					operandB, stack = pop(stack)
					operation, stack = pop(stack)
					operandA, stack = pop(stack)

					switch tokenType(operation.(rune)) {
					case Add:
						stack = append(stack, operandA.(int)+operandB.(int))
					case Multiply:
						stack = append(stack, operandA.(int)*operandB.(int))
					}
				}
			}

		case Add, Multiply:
			stack = append(stack, r)
		case Space:
			continue
		}
	}

	return stack[0].(int)
}

func popParenResult(s []any) (any, []any) {
	var value any
	_, s = pop(s)
	value, s = pop(s)
	_, s = pop(s)
	return value, s
}

func pop(s []any) (any, []any) {
	value, newS := s[len(s)-1], s[:len(s)-1]
	return value, newS
}

func popUntil(s []any, t token) []any {
	for value, ok := s[len(s)-1].(rune); ok && tokenType(value) != t; {
		pop(s)
	}
	return s
}

func peek(s []any, index int) any {
	return s[len(s)-1-index]
}
