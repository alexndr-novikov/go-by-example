package main

import (
	"fmt"
	"testing"
)

// run as  go test -bench=. -benchtime 5s -benchmem

func BenchmarkSubsequenceRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsSubsequence("abc", "ahbgdc")
		IsSubsequence("axc", "ahbgdc")
		IsSubsequence("axcgle", "Ifindthatbenchmarksresultsaredifficulttoread")
	}
}

func BenchmarkSubsequencePlain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsSubsequenceImproved("abc", "ahbgdc")
		IsSubsequenceImproved("axc", "ahbgdc")
		IsSubsequenceImproved("axcgle", "Ifindthatbenchmarksresultsaredifficulttoread")
	}
}

func TestSubsequencePlain(t *testing.T) {
	var tests = []struct {
		a, b string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
		{"axcgle", "Ifindthatbenchmarksresultsaredifficulttoread", false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s", tt.a, tt.b), func(t *testing.T) {
			ans := IsSubsequenceImproved(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}

func TestSubsequenceRegex(t *testing.T) {
	var tests = []struct {
		a, b string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
		{"axcgle", "Ifindthatbenchmarksresultsaredifficulttoread", false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s", tt.a, tt.b), func(t *testing.T) {
			ans := IsSubsequence(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}
