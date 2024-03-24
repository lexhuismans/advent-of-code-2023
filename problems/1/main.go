package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var total int = 0

	// Loop over digits
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var txt string = scanner.Text()
		txt = strings.Replace(txt, "one", "o1e", -1)
		txt = strings.Replace(txt, "two", "t2o", -1)
		txt = strings.Replace(txt, "three", "t3e", -1)
		txt = strings.Replace(txt, "four", "4", -1)
		txt = strings.Replace(txt, "five", "5e", -1)
		txt = strings.Replace(txt, "six", "6", -1)
		txt = strings.Replace(txt, "seven", "7n", -1)
		txt = strings.Replace(txt, "eight", "e8t", -1)
		txt = strings.Replace(txt, "nine", "n9e", -1)

		var nb [2]rune
		// First digit
		for _, char := range txt {
			if unicode.IsDigit(char) {
				nb[0] = char
				break
			}
		}

		// Last digit
		rs := []rune(txt)
		for i := len(rs) - 1; i >= 0; i-- {
			if unicode.IsDigit(rs[i]) {
				nb[1] = rs[i]
				break
			}
		}

		num, err := strconv.Atoi(string(nb[0]) + string(nb[1]))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		time.Sleep(5 * time.Millisecond)

		total += num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result: ", total)
}
