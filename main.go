// table writes csv type formatted input in tabular format

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func die(s string) {
	fmt.Fprintf(os.Stderr, s)
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "usage: slice from:to\n")
	os.Exit(1)
}

func main() {
	if len(os.Args) == 1 {
		die("no cutstring specified")
	}
	if len(os.Args) > 2 {
		die("more than one argument")
	}
	cutstring := os.Args[1]
	x := strings.Split(cutstring, ":")
	if len(x) > 2 && x[0] != "" {
		die("invalid cutstring format: must be `from:` or `:to` or `from:to`")
	}
	var from, to int
	var err error
	if x[0] == "" {
		x[0] = "0"
	}
	from, err = strconv.Atoi(x[0])
	if err != nil {
		die(fmt.Sprintf("invalid cutstring format: %v", err))
	}
	if len(x) == 2 && x[1] != "" {
		to, err = strconv.Atoi(x[1])
		if err != nil {
			die(fmt.Sprintf("invalid cutstring format: %v", err))
		}
		if to == 0 {
			die("to cannot be 0")
		}
	}
	if from > to && from >= 0 && to >= 0 {
		die("from must be before to")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var fromm, too int
		too = to
		fromm = from

		// handle negative slices
		if fromm < 0 {
			fromm = len(line) + fromm
		}
		if too < 0 {
			too = len(line) + too
		}

		// handle variable length
		if too == 0 || too > len(line) {
			too = len(line)
		}

		// if from > len(line) {
		if from > too {
			fmt.Println()
			continue
		}

		// print sliced line
		fmt.Println(line[fromm:too])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
