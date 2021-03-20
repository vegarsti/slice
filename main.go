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
	from, to, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "usage: slice from:to\n")
		os.Exit(1)
	}
	if err := slice(os.Stdin, os.Stdout, from, to); err != nil {
		fmt.Fprintf(os.Stderr, "write: %v\n", err)
		os.Exit(1)
	}
}

func parseArgs(args []string) (int, int, error) {
	var from, to int
	var err error

	if len(args) == 0 {
		return 0, 0, fmt.Errorf("no range specified")
	}
	if len(args) > 1 {
		return 0, 0, fmt.Errorf("only takes one argument")
	}
	arg := args[0]
	fromTo := strings.Split(arg, ":")
	if len(fromTo) > 2 {
		return 0, 0, fmt.Errorf("invalid range format: must be `from:` or `:to` or `from:to`")
	}
	if fromTo[0] == "" {
		fromTo[0] = "0"
	}
	from, err = strconv.Atoi(fromTo[0])
	if err != nil {
		return 0, 0, fmt.Errorf(fmt.Sprintf("invalid range format: %v", err))
	}
	if len(fromTo) == 2 && fromTo[1] != "" {
		to, err = strconv.Atoi(fromTo[1])
		if err != nil {
			return 0, 0, fmt.Errorf(fmt.Sprintf("invalid range format: %v", err))
		}
		if to == 0 {
			return 0, 0, fmt.Errorf("to cannot be 0")
		}
	}
	if from > to && from >= 0 && to > 0 {
		return 0, 0, fmt.Errorf("from must be before to")
	}
	if from < 0 && to > 0 {
		return 0, 0, fmt.Errorf("to must be negative if from is negative")
	}
	return from, to, nil
}

func slice(in io.Reader, out io.Writer, from int, to int) error {
	s := bufio.NewScanner(in)
	w := bufio.NewWriter(out)
	for s.Scan() {
		line := s.Text()
		slicedLine := sliceLine(line, from, to)
		w.WriteString(slicedLine)
		w.WriteString("\n")
	}
	w.Flush()
	if err := s.Err(); err != nil {
		return err
	}
	return nil
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func sliceLine(line string, from int, to int) string {
	// negative from
	if from < 0 {
		from = max(0, len(line)+from)
	}

	// to exceeds line length
	if to == 0 || to > len(line) {
		to = len(line)
	}

	// negative to
	if to < 0 {
		to = max(0, len(line)+to)
	}

	if from >= to {
		return ""
	}
	return line[from:to]
}
