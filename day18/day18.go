package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

type any interface{}

type token int

const (
	Number     token = 0
	Add        token = 1
	Multiply   token = 2
	OpenParen  token = 3
	CloseParen token = 4
)

func tokenType(r rune) token {
	switch r {
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
	part2Solution := part2(expressions)

	fmt.Printf("Day 18: Part 1: = %+v\n", part1Solution)
	fmt.Printf("Day 18: Part 2: = %+v\n", part2Solution)
}

func part1(expressions []string) int {
	sum := 0
	for _, expression := range expressions {
		runes := []rune(strings.ReplaceAll(expression, " ", ""))
		sum += evaluateExpression(&runes)
	}
	return sum
}

func part2(expressions []string) int {
	sum := 0
	for _, expression := range expressions {
		runes := []rune(strings.ReplaceAll(expression, " ", ""))
		sum += evaluateExpression2(&runes)
	}
	return sum
}

func evaluateExpression(e *[]rune) int {
	stack := []any{}

	for len(*e) != 0 {
		tok := shift(e)

		switch tokenType(tok) {
		case OpenParen:
			stack = append(stack, evaluateExpression(e))
		case Number:
			value, _ := strconv.Atoi(string(tok))
			stack = append(stack, value)
		case Add, Multiply:
			stack = append(stack, tok)
		case CloseParen:
			return stack[0].(int)
		}

		if len(stack) >= 3 {
			paramA, operator, paramB := stack[0].(int), stack[1].(rune), stack[2].(int)
			stack = nil
			switch tokenType(operator) {
			case Add:
				stack = append(stack, paramA+paramB)
			case Multiply:
				stack = append(stack, paramA*paramB)
			}
		}
	}
	return stack[0].(int)
}

func evaluateExpression2(e *[]rune) int {
	stack := []any{}

	done := false
	for len(*e) != 0 && !done {
		tok := shift(e)

		switch tokenType(tok) {
		case OpenParen:
			stack = append(stack, evaluateExpression2(e))
		case Number:
			value, _ := strconv.Atoi(string(tok))
			stack = append(stack, value)
		case Add, Multiply:
			stack = append(stack, tok)
		case CloseParen:
			done = true
			break
		}

		// Handle additions only
		if len(stack) >= 3 {
			var paramA, paramB int
			var operator rune

			if v, ok := stack[len(stack)-3].(int); ok {
				paramA = v
			}
			if v, ok := stack[len(stack)-2].(rune); ok {
				operator = v
			}
			if v, ok := stack[len(stack)-1].(int); ok {
				paramB = v
			}

			if tokenType(operator) == Add {
				stack = stack[:len(stack)-3]
				stack = append(stack, paramA+paramB)
			}
		}
	}

	// Take care of multiplications
	for len(stack) > 1 {
		paramA, operator, paramB := stack[0].(int), stack[1].(rune), stack[2].(int)
		if tokenType(operator) == Multiply {
			stack = append([]any{paramA * paramB}, stack[3:]...)
		}
	}

	return stack[0].(int)
}

func shift(l *[]rune) rune {
	r := (*l)[0]
	*l = (*l)[1:]
	return r
}
