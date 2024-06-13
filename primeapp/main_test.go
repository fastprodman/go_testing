package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
	}{
		{"prime", 7, true},
		{"not prime", 8, false},
		{"not prime by definition", 1, false},
		{"negative not prime", -1, false},
	}

	for _, test := range primeTests {
		result, _ := isPrime(test.testNum)
		if test.expected != result {
			t.Errorf("Test %s: expected %t got %t", test.name, test.expected, result)
		}
	}
}

func Test_prompt(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	prompt()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if string(out) != "-> " {
		t.Errorf("incorrect prompt")
	}
}

func Test_intro(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	intro()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "Enter a number or press q to quit.") {
		t.Errorf("does not contain expected str, got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a number"},
		{name: "Quit", input: "q", expected: ""},
		{name: "Not prime", input: "1", expected: "1 is not prime by definition"},
	}

	for _, test := range tests {

		input := strings.NewReader(test.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, test.expected) {
			t.Errorf("In test %s expected %s got %s", test.name, test.expected, res)
		}

	}
}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)
	var testin bytes.Buffer
	testin.Write([]byte("1\nq\n"))

	go readUserInput(&testin, doneChan)
	<-doneChan
	close(doneChan)
}
