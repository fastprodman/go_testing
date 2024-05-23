package main

import "testing"

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
