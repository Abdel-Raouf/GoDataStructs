package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func accept() float64 {
	tok := ""
	_, err := fmt.Scan(&tok)

	if err != nil || {
		os.Exit(0)
	}

	switch tok {					//what is tok?
	case "+":
		return accept() + accept()
	case "-":
		return accept() - accept()
	case "*":
		return accept() * accept()
	case "/":
		return accept() / accept()
	case "sqrt":
		return math.Sqrt(accept())
	case "expt":
		return math.Pow(accept(), accept())
	default:
		r, _ := strconv.ParseFloat(tok, 64)		//is a number

		return r
	}
}

func main() {
	fmt.Println("Prefix calculator")
	fmt.Println("Tokens are separated by whitespaces.")
	fmt.Println("Operators:")
	fmt.Println("\tUnary: sqrt")
	fmt.Println("\tBinary: + - * / expt")
	for {
		fmt.Print("> ")
		fmt.Printf("%.16g\n", accept())
	}
	f, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Println(f)
}
