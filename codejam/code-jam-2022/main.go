package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input []string

	//scanner := bufio.NewScanner(os.Stdin)
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}
		input = append(input, scanner.Text())
	}
	fmt.Println("scanner.Err()", scanner.Err())

	cases, _ := strconv.Atoi(input[0])
	for tc := 0; tc < cases; tc++ {
		SS := strings.Fields(input[2+tc*2])
		S := make([]int, len(SS))
		for i, v := range SS {
			S[i], _ = strconv.Atoi(v)
		}

		sort.Ints(S)

		var i, maxLength int
		for {
			if i == len(S) {
				break
			}

			if S[i] < maxLength+1 {
				i++
				continue
			}

			i++
			maxLength++
		}
		fmt.Printf("Case #%d: %d\n", tc+1, maxLength)
	}
}
