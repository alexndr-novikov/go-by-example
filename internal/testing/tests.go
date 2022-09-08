package main

import (
	"fmt"
	"regexp"
	"strings"
)

func IsSubsequence(s string, t string) bool {
	reg := regexp.MustCompile(".*" + strings.Join(strings.Split(s, ""), ".*") + ".*")
	return reg.MatchString(t)
}
func IsSubsequenceImproved(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return len(s) == i
}

func main() {
	fmt.Println("Testing package output:")
	fmt.Println(IsSubsequence("abc", "ahbgdc"))
	fmt.Println(IsSubsequence("axc", "ahbgdc"))
	fmt.Println(IsSubsequenceImproved("abc", "ahbgdc"))
	fmt.Println(IsSubsequenceImproved("axc", "ahbgdc"))
	fmt.Println("")
}
