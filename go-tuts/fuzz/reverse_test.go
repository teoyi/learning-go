package main 

import (
	"testing"
)

// this is a unit test 
func TestReverse(t *testing.T) {
	testcases := []struct { 
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
        {" ", " "},
        {"!12345", "54321!"},
	}

	for _,tc := range testcases { 
		rev := Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

// limits of unit test is that each input must be added to test by the dev 
// Fuzzing comes up with inputs for the code and may identify edge cases that the test cases the dev comes up with didnt reach . 