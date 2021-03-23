package main

import (
	"fmt"
	"testing"
)

func TestParseArgsOK(t *testing.T) {
	tcs := []struct {
		slice string
		from  int
		to    int
	}{
		{":", 0, 0},
		{":1", 0, 1},
		{"1:", 1, 0},
		{":-1", 0, -1},
		{"-2:-1", -2, -1},
		{"1:-1", 1, -1},
	}
	for _, tc := range tcs {
		t.Run(tc.slice, func(t *testing.T) {
			from, to, err := parseArgs([]string{tc.slice})
			if err != nil {
				t.Fatalf("got error: '%v'", err)
			}
			if from != tc.from {
				t.Errorf("from is %d, expected %d", from, tc.from)
			}
			if to != tc.to {
				t.Errorf("to is %d, expected %d", to, tc.to)
			}
		})
	}
}

func TestParseArgsError(t *testing.T) {
	tcs := []struct {
		slice string
	}{
		{"abc"},
		{"abc:"},
		{"abc:123"},
		{":abc"},
		{"1:2:"},
		{":2:"},
		{"1:2:"},
		{":2:3"},
		{"1:2:3"},
		{"-2:1"},
		{"-1:-2"},
		{"0:"},
		{":0"},
	}
	for _, tc := range tcs {
		t.Run(tc.slice, func(t *testing.T) {
			if _, _, err := parseArgs([]string{tc.slice}); err == nil {
				t.Error("expected error, but was nil")
			}
		})
	}
}

func TestSliceLine(t *testing.T) {
	tcs := []struct {
		line  string
		slice string
		from  int
		to    int
	}{
		{"", "", 0, 0},
		{"a", "a", -1, 0},
		{"a", "a", -2, 0},
		{"a", "a", -2, 0},
		{"ab", "b", -1, 0},
		{"ab", "a", 0, 1},
		{"ab", "ab", 0, 2},
		{"ab", "a", 0, -1},
		{"ab", "", 0, -2},
		{"ab", "", 0, -3},
	}
	for _, tc := range tcs {
		description := fmt.Sprintf("%d:%d on %s", tc.from, tc.to, tc.line)
		t.Run(description, func(t *testing.T) {
			if s := sliceLine(tc.line, tc.from, tc.to); s != tc.slice {
				t.Errorf("expected '%s', but was '%s'", tc.slice, s)
			}
		})
	}
}
