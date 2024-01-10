/*
 * https://leetcode.com/problems/reverse-vowels-of-a-string/
 */
package main

import (
	"fmt"
)

var s = string("lwowptepppl")

func main() {
	fmt.Println(mycode(s))
	fmt.Println(withBetterSyntax(s))
}

// This is effecient in looping but is slow because of strings and map usage
func mycode(s string) string {
	vowelMap := make(map[uint8]bool)
	vowelMap['a'] = true
	vowelMap['e'] = true
	vowelMap['i'] = true
	vowelMap['o'] = true
	vowelMap['u'] = true
	vowelMap['A'] = true
	vowelMap['E'] = true
	vowelMap['I'] = true
	vowelMap['O'] = true
	vowelMap['U'] = true

	leftI := 0
	rightI := len(s) - 1
	i := 0
	for leftI < rightI {
		i++
		_, leftV := vowelMap[s[leftI]]
		_, rightV := vowelMap[s[rightI]]
		if leftV && rightV {
			temp := s[leftI]
			s = s[:leftI] + string(s[rightI]) + s[leftI+1:]
			s = s[:rightI] + string(temp) + s[rightI+1:]
			leftI++
			rightI--
		}
		if !leftV {
			leftI++
		}
		if !rightV {
			rightI--
		}
	}
	fmt.Println(i)
	return s
}

// Note it's ineffecient in looping, but the use of ruines and no maps makes it much faster.
func withBetterSyntax(s string) string {
	lo, hi := 0, len(s)-1
	runes := make([]rune, len(s))
	for i, r := range s {
		runes[i] = r
	}
	i := 0
	for lo < hi {
		i++
		if !isVowel(runes[hi]) {
			hi--
			continue
		}

		if !isVowel(runes[lo]) {
			lo++
			continue
		}

		runes[lo], runes[hi] = runes[hi], runes[lo]
		lo++
		hi--
	}
	fmt.Println(i)
	return string(runes)
}

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}
