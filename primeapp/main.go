package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	intro()
	doneChan := make(chan bool)
	go readUserInput(os.Stdin, doneChan)
	<-doneChan
	close(doneChan)
	fmt.Println("End")
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

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumbers(scanner)
		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a number", false
	}

	_, msg := isPrime(numToCheck)

	return msg, false
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
