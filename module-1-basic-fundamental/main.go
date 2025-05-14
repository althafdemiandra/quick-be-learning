package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(num1 float64, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("Cannot divide by zero")
		}
		return num1 / num2, nil
	default:
		fmt.Println()
		return 0, errors.New("Invalid operator")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a number: ")
	// ReadString reads until the first occurrence of delimiter '\n' (newline)
	// This tells the reader to keep reading input until user presses Enter
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input)
	if input == "" {
		fmt.Print("No input provided. Using default value 0.\n")
		input = "0"
	}

	// convert input to number
	num_input, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error converting first input to number:", err)
		return
	}

	fmt.Print("Enter a s econd number: ")

	input2, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error converting second input to number:", err)
		return
	}

	input2 = strings.TrimSpace(input2)
	num_input2, err := strconv.ParseFloat(input2, 64)
	if err != nil {
		fmt.Println("Error converting input to number:", err)
		return
	}

	fmt.Printf(`you input %.2f and %.2f\n`, num_input, num_input2)

	fmt.Print("Enter operator (+, -, *, /): ")
	operator, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading operator:", err)
		return
	}
	operatorMap := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	operator = strings.TrimSpace(operator)

	if !operatorMap[operator] {
		fmt.Println("Invalid operator. Please use +, -, *, or /")
		return
	}

	res, err := calculate(num_input, num_input2, operator)
	if err != nil {
		fmt.Println("Error calculating:", err)
		return
	}

	fmt.Println("result: ", res)

}
