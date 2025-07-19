package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Go Calculator!")
	fmt.Println("Enter your calculation in the format: number1 operator number2")
	fmt.Println("Supported operators: +, -, *, /")
	fmt.Println("Type 'exit' to quit.")

	for {
		fmt.Print("Enter calculation: ")
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Buffer(make([]byte, 1024), 1024)
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")

		if input[0] == "exit" {
			fmt.Println("Exiting the calculator. Goodbye!")
			break
		}

		fmt.Println("Processing input:", input)
		if len(input) != 3 {
			fmt.Println("Error: Please provide a valid calculation in the format: number1 operator number2")
			continue
		}

		num1, err := strconv.ParseFloat(string(input[0]), 64)
		if err != nil {
			fmt.Println("Error: Invalid number format for the first operand.")
			continue
		}
		operator := string(input[1])
		num2, err := strconv.ParseFloat(string(input[2]), 64)
		if err != nil {
			fmt.Println("Error: Invalid number format for the first operand.")
			continue
		}

		switch operator {
		case "+":
			fmt.Printf("Result: %1.f + %1.f = %1.f\n", num1, num2, num1+num2)
		case "-":
			fmt.Printf("Result: %1.f - %1.f = %1.f\n", num1, num2, num1-num2)
		case "*":
			fmt.Printf("Result: %1.f * %1.f = %1.f\n", num1, num2, num1*num2)
		case "/":
			if num2 == 0 {
				fmt.Println("Error: Division by zero is not allowed.")
				return
			}
			fmt.Printf("Result: %1.f / %1.f = %1.f\n", num1, num2, num1/num2)
		default:
			fmt.Println("Error: Unsupported operator. Use +, -, *, or /.")
			return
		}
	}
}
