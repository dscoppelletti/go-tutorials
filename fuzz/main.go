package main

import (
	"errors"
    "fmt"
    "unicode/utf8"
)

// This function will accept a string, loop over it a byte at a time, and
// return the reversed string at the end.
// func Reverse(s string) string {
//     b := []byte(s)
//     for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
//         b[i], b[j] = b[j], b[i]
//     }
//     return string(b)
// }

// To correct the Reverse function, let’s traverse the string by runes, instead
// of by bytes.
// func Reverse(s string) string {
// 	fmt.Printf("input: %q\n", s)
//     r := []rune(s)
// 	fmt.Printf("runes: %q\n", r)
//     for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
//         r[i], r[j] = r[j], r[i]
//     }
//     return string(r)
// }

// This change will return an error if the input string contains characters
// which are not valid UTF-8.
func Reverse(s string) (string, error) {
    if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r), nil
}

func main() {
    input := "The quick brown fox jumped over the lazy dog"
    rev, revErr := Reverse(input)
    doubleRev, doubleRevErr := Reverse(rev)
    fmt.Printf("original: %q\n", input)
    fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
    fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}
