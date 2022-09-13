package regular_expressions

import (
	"bytes"
	"fmt"
	"regexp"
)

func RunSample() {
	p := fmt.Println

	p("regular-expressions package output:")
	match, _ := regexp.MatchString("[ph]*", "php")
	p(match)

	r, _ := regexp.Compile("[1-3][2-5][a-z]*")

	p(r.MatchString("a13aaa"))
	p(r.MatchString("a11aaa"))

	p(r.FindString("a13aaa"))
	p("idx", r.FindStringIndex("a13aaa"))
	p("FindStringSubmatch", r.FindStringSubmatch("a13aaa"))
	p("FindAllString", r.FindAllString("a13aa aa13aaa", -1))
	p("FindAllString", r.FindAllStringSubmatchIndex("a13aa aa13aaa", -1))
	p("Replace", r.ReplaceAllString("a13aa aa13aaa", "<censored>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	p(string(out))

	p("")
}
