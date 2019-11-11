package main

import (
	"testing"
)

func TestCountTokenEntries(t *testing.T) {
	testcases := []struct {
		s    string
		t    string
		outp int
	}{
		{"gogogo", "go", 3},
		{"foobar", "go", 0},
		{"GOgoGo", "go", 1},
		{"GOgoGoogggOGOO", "[Gg][Oo]", 5},
		{"Gogo Go Gogo", `\sGo\s`, 1},
	}

	for _, tc := range testcases {
		if res, err := countTokenEntries(tc.s, tc.t); err == nil {
			if res != tc.outp {
				t.Errorf("expected %d, got %d on testcase '%v'", tc.outp, res, tc)
			}
		} else {
			t.Errorf("error occured on testcase '%v': %v", tc, err)
		}
	}
}

func TestCountTokenEntriesError(t *testing.T) {
	if _, err := countTokenEntries("foo", "\\\324343"); err == nil {
		t.Errorf("expected parsing regexp error, got nil")
	}
}
