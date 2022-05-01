package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PalindromicFactors() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			break
		}
		input = append(input, t)
	}

	nCases, _ := strconv.Atoi(input[0])
	for tc := 0; tc < nCases; tc++ {
		A, _ := strconv.Atoi(input[tc+1])

		var factors []int
		i := 1
		for i*i <= A {
			if A%i == 0 {
				factors = append(factors, i)
				if A/i != i {
					factors = append(factors, A/i)
				}
			}
			i++
		}

		palindromes := 0
		for _, factor := range factors {
			s := strconv.Itoa(factor)
			if isPalindrome(s) {
				palindromes++
			}
		}

		fmt.Printf("Case #%d: %v\n", tc+1, palindromes)
	}
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
