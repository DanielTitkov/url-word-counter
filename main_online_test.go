// +build online

package main

import "testing"

func TestCountTokenAtURL(t *testing.T) {
	// WARNING! This test will fail if executed offline
	// Though w3 pages are unlikely to change it still may happen
	// so the test is not 100% stable
	// https://wordcounter.net/website-word-count was used to check word count
	testcases := []struct {
		u    string
		t    string
		outp int
	}{
		{"https://www.w3.org/TR/2003/REC-PNG-20031110/", "commence", 1},
		{"https://www.w3.org/TR/2015/REC-webmessaging-20150519/", "strawberry", 0},
	}

	for _, tc := range testcases {
		if res, err := countTokenAtURL(tc.u, tc.t); err == nil {
			if res != tc.outp {
				t.Errorf("expected %d, got %d on testcase '%v'", tc.outp, res, tc)
			}
		} else {
			t.Errorf("error occured on testcase '%v': %v", tc, err)
		}
	}
}
