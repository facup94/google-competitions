package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func InfinityArea() {
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
		RAB := strings.Split(input[tc+1], " ")
		R, _ := strconv.Atoi(RAB[0])
		A, _ := strconv.Atoi(RAB[1])
		B, _ := strconv.Atoi(RAB[2])

		var total, r float64
		r = float64(R)
		total += math.Pi * math.Pow(r, 2)

		for {
			R *= A
			total += math.Pi * math.Pow(float64(R), 2)

			R /= B
			if R == 0 {
				break
			}
			total += math.Pi * math.Pow(float64(R), 2)
		}

		fmt.Printf("Case #%d: %v\n", tc+1, total)
	}
}
