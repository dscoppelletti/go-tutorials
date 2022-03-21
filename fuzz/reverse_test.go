package main

import (
	"testing"
	"unicode/utf8"
)

// This simple test will assert that the listed input strings will be correctly
// reversed.
// The unit test has limitations, namely that each input must be added to the
// test by the developer.
// func TestReverse(t *testing.T) {
//     testcases := []struct {
//         in, want string
//     } {
//         { "Hello, world", "dlrow ,olleH" },
//         { " ", " " },
//         { "!12345", "54321!" },
//     }
//     for _, tc := range testcases {
//         rev := Reverse(tc.in)
//         if rev != tc.want {
//                 t.Errorf("Reverse: %q, want %q", rev, tc.want)
//         }
//     }
// }

// One benefit of fuzzing is that it comes up with inputs for your code, and may
// identify edge cases that the test cases you came up with didn’t reach.
//
// Fuzzing has a few limitations as well. In your unit test, you could predict
// the expected output of the Reverse function, and verify that the actual
// output met those expectations.
// For example, in the test case Reverse("Hello, world") the unit test specifies
// the return as "dlrow ,olleH".
// When fuzzing, you can’t predict the expected output, since you don’t have
// control over the inputs.
// However, there are a few properties of the Reverse function that you can
// verify in a fuzz test. The two properties being checked in this fuzz test
// are:
// 1) Reversing a string twice preserves the original value
// 2) The reversed string preserves its state as valid UTF-8.
func FuzzReverse(f *testing.F) {
    testcases := []string { "Hello, world", " ", "!12345" }
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
    // f.Fuzz(func(t *testing.T, orig string) {
    //     rev := Reverse(orig)
    //     doubleRev := Reverse(rev)
	// 	// This t.Logf line will print to the command line if an error occurs,
	// 	// or if executing the test with -v, which can help you debug this
	// 	// particular issue.
	// 	t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d",
	// 		utf8.RuneCountInString(orig), utf8.RuneCountInString(rev),
	// 		utf8.RuneCountInString(doubleRev))
    //     if orig != doubleRev {
    //         t.Errorf("Before: %q, after: %q", orig, doubleRev)
    //     }
    //     if utf8.ValidString(orig) && !utf8.ValidString(rev) {
    //         t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
    //     }
    // })
	f.Fuzz(func(t *testing.T, orig string) {
        rev, err1 := Reverse(orig)
        if err1 != nil {
            return
        }
        doubleRev, err2 := Reverse(rev)
        if err2 != nil {
             return
        }
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}
