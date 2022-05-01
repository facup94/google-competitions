package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func UnlockThePadlock() {
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
		var movements int
		ND := strings.Split(input[tc*2+1], " ")
		N, _ := strconv.Atoi(ND[0])
		D, _ := strconv.Atoi(ND[1])

		dials := make([]int, N)
		distances := make([]int, N)
		for i, v := range strings.Split(input[tc*2+2], " ") {
			dials[i], _ = strconv.Atoi(v)
			dist := dials[i]
			if D-dials[i] < dist {
				dist = D - dials[i]
			}
			distances[i] = dist
		}

		var i, j = 0, len(dials) - 1
		for {
			for i < j && distances[i] == 0 {
				i++
			}
			for i < j && distances[j] == 0 {
				j--
			}
			if i >= j && distances[i] == 0 {
				break
			}
			movements++

			if distances[i] < distances[j] {
				moveUp := true
				if dials[i] == distances[i] {
					moveUp = false
				}
				for k := i; k <= j; k++ {
					if moveUp {
						dials[k]++
						if dials[k] == D {
							dials[k] = 0
						}
					} else {
						dials[k]--
						if dials[k] == -1 {
							dials[k] = D - 1
						}
					}
				}
			} else {
				moveUp := true
				if dials[j] == distances[j] {
					moveUp = false
				}
				for k := i; k <= j; k++ {
					if moveUp {
						dials[k]++
						if dials[k] == D {
							dials[k] = 0
						}
					} else {
						dials[k]--
						if dials[k] == -1 {
							dials[k] = D - 1
						}
					}
				}
			}

			// recalculate distances
			for k, v := range dials {
				dist := v
				if D-v < dist {
					dist = D - v
				}
				distances[k] = dist
			}
			fmt.Println(dials)
		}

		fmt.Printf("Case #%d: %v\n", tc+1, movements)
	}
}
