// table writes csv type formatted input in tabular format

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
	from, to, err := parseArgs(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "usage: slice from:to\n")
		os.Exit(1)
	}
	if err := write(os.Stdin, os.Stdout, from, to); err != nil {
		fmt.Fprintf(os.Stderr, "write: %v\n", err)
		os.Exit(1)
	}
}

func parseArgs(args []string) (int, int, error) {
	var from, to int
	var err error

	if len(args) == 1 {
		return 0, 0, fmt.Errorf("no cutstring specified")
	}
	if len(args) > 2 {
		return 0, 0, fmt.Errorf("more than one argument")
	}
	cutstring := args[1]
	x := strings.Split(cutstring, ":")
	if len(x) > 2 && x[0] != "" {
		return 0, 0, fmt.Errorf("invalid cutstring format: must be `from:` or `:to` or `from:to`")
	}
	if x[0] == "" {
		x[0] = "0"
	}
	from, err = strconv.Atoi(x[0])
	if err != nil {
		return 0, 0, fmt.Errorf(fmt.Sprintf("invalid cutstring format: %v", err))
	}
	if len(x) == 2 && x[1] != "" {
		to, err = strconv.Atoi(x[1])
		if err != nil {
			return 0, 0, fmt.Errorf(fmt.Sprintf("invalid cutstring format: %v", err))
		}
		if to == 0 {
			return 0, 0, fmt.Errorf("to cannot be 0")
		}
	}
	if from > to && from >= 0 && to >= 0 {
		return 0, 0, fmt.Errorf("from must be before to")
	}
	return from, to, nil
}

func write(in io.Reader, out io.Writer, from int, to int) error {
	s := bufio.NewScanner(in)
	w := bufio.NewWriter(out)
	var fromm, too int
	for s.Scan() {
		line := s.Text()
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

		if from > too {
			w.WriteString("")
		} else {
			w.WriteString(line[fromm:too])
		}
		w.WriteString("\n")
	}
	w.Flush()
	if err := s.Err(); err != nil {
		return err
	}
	return nil
}
