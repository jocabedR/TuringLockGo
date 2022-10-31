package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	commandAnalyzer(lines, 0, 0, 0)
}

func commandAnalyzer(text []string, counterLine int, a int, b int) {
	if counterLine < len(text)-1 {
		fmt.Println(counterLine, ". ", text[counterLine])
		words := strings.Split(text[counterLine], " ")
		instruction := words[0]

		fmt.Println(instruction)
		var offset string
		var aux int = 0
		var value int

		if instruction[0:1] == "j" {
			if instruction == "jmp" {
				offset = words[1]
				aux = jumpTo(counterLine, offset, len(text)-1)
			} else {
				offset = words[2]
				value = getRegisterValue(words[1][0:1], a, b)
				if instruction == "jie" && value%2 == 0 {
					aux = jumpTo(counterLine, offset, len(text)-1)
				} else if instruction == "jio" && value == 1 {
					aux = jumpTo(counterLine, offset, len(text)-1)
				} else {
					aux = counterLine + 1
				}
			}
		} else {
			aux = counterLine + 1
			if words[1] == "a" {
				a = setRegisterValue(instruction, a)
			} else {
				b = setRegisterValue(instruction, b)
			}
		}

		fmt.Println("Jump to: ", aux)
		fmt.Println("a = ", a, "b = ", b, "\n")
		commandAnalyzer(text, aux, a, b)
	} else {
		fmt.Println("Program finished. \na = ", a, "b = ", b)
	}
}

func jumpTo(counterLine int, offset string, limit int) int {
	var longitudJump int
	direction := offset[0:1]
	longitud, _ := strconv.Atoi(offset[1:])

	if direction == "+" {
		longitudJump = counterLine + longitud
	} else {
		longitudJump = counterLine - longitud
	}
	if longitudJump > limit || longitudJump < 0 {
		longitudJump = limit
	}
	return longitudJump
}

func getRegisterValue(register string, a, b int) int {
	if register == "a" {
		return a
	} else {
		return b
	}
}

func setRegisterValue(operator string, register int) int {
	var response int
	switch operator {
	case "hlf":
		response = register / 2
		break
	case "tpl":
		response = register * 3
		break
	case "inc":
		response = register + 1
		break
	}
	return response
}
