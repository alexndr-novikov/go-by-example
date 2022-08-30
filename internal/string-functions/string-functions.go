package string_functions

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func RunSample() {
	p("String functions package output:")
	str := "lorem ipsum dolor sit amet, consectetur adipiscing elit. curabitur nunc."
	p("Contains: s.Contains(str, \"lorem\"):", s.Contains(str, "lorem"))
	p("Contains: s.Contains(str, \"Lorem\"):", s.Contains(str, "Lorem"))
	p("Count: s.Count(str, \"o\"):", s.Count(str, "o"))
	p("HasSuffix: s.HasSuffix(\"test\", \"st\"):", s.HasSuffix("test", "st"))
	p("HasPrefix: s.HasPrefix(\"test\", \"te\"):", s.HasPrefix("test", "te"))
	p("Index: s.Index(\"test\", \"e\"):", s.Index("test", "e"))
	p("Join: s.Join([]string{\"a\", \"b\", \"c\"}, \"-\"):", s.Join([]string{"a", "b", "c"}, "-"))
	p("Split: s.Split(str, \" \"):", s.Split(str, " "))
	p("ToLower: s.ToLower(str):", s.ToLower(str))
	p("ToUpper: s.ToUpper(str):", s.ToUpper(str))
	p("ToTitle: s.ToTitle(str):", s.ToTitle(str))
	p("Repeat: s.Repeat(\"a\", 5):", s.Repeat("a", 5))
	p("Replace: s.Replace(\"foo\", \"o\", \"0\", 1):", s.Replace("foo", "o", "0", 1))
	p("Replace: s.Replace(\"foo\", \"o\", \"0\", -1):", s.Replace("foo", "o", "0", -1))
	p("")
}
