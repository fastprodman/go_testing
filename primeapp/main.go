package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	intro()
	doneChan := make(chan bool)
}

func intro() {
	fmt.Println("Is it prime?")
	fmt.Println("-------------")
	fmt.Println("Enter a number or press q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		res, done := checkNumbers(scanner)
		if done {
			doneChan <- true
			return
		}

	}
}

func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime by definition", n)
	}

	if n < 0 {
		return false, "Negative numbers are not prime by defenition"
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number", n)
}
