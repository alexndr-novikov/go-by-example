package time_formatting

import (
	"fmt"
	"time"
)

func RunSample() {

	p := fmt.Println
	p("time-formatting package output:")

	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1, e)

	p(t.Format(time.Kitchen))

	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)

	p("")
}
