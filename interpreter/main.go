package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

const wrapAroundValue = -8
const wrapTargetValue = 128

func main() {
	var input string
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println("error reading file:", err)
			return
		}
		input = string(data)
	} else {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("error reading stdin:", err)
			return
		}
		input = string(bytes)
	}

	interpret(input)
}

func interpret(code string) {
	counterStack := []int{0}
	secondCounter := 0
	reader := bufio.NewReader(os.Stdin)

	runes := []rune(code)

	for i := 0; i < len(runes); i++ {
		if i+3 < len(runes) && string(runes[i:i+4]) == "goog" {
			fmt.Print("goog...")
			i += 3
			continue
		}

		ch := runes[i]

		switch ch {
		case '[':
			counterStack = append(counterStack, 0)
		case ']':
			if len(counterStack) > 1 {
				counterStack = counterStack[:len(counterStack)-1]
			} else {
				fmt.Println("warning: unmatched closing bracket")
			}
		case '.':
			counterStack[len(counterStack)-1]--
			if counterStack[len(counterStack)-1] < wrapAroundValue {
				counterStack[len(counterStack)-1] = wrapTargetValue
			}
		case '+':
			fmt.Print(counterStack[len(counterStack)-1])
		case '/':
			fmt.Print(string(rune(counterStack[len(counterStack)-1])))
		case '#':
			fmt.Print("input a character: ")
			input, err := reader.ReadByte()
			if err != nil {
				fmt.Println("error reading input:", err)
				continue
			}
			secondCounter = int(input)
		case '*':
			counterStack[len(counterStack)-1] = secondCounter
		case '\n':
			counterStack[len(counterStack)-1] = 0
		case '&':
			fmt.Println()
		default:
		}
	}
}
