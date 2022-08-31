package time_package

import (
	"fmt"
	"time"
)

func RunSample() {
	fmt.Println("Time package output:")
	p := fmt.Println

	now := time.Now()
	then := time.Date(
		2025, 11, 17, 20, 34, 58, 651387237, time.UTC,
	)

	p("Now: ", now)
	p("Then: ", then)
	p("Then Year: ", then.Year())
	p("Then Month: ", then.Month())
	p("Then Day: ", then.Day())
	p("Then Hour: ", then.Hour())
	p("Then Minute: ", then.Minute())
	p("Then Second: ", then.Second())
	p("Then Nanosecond: ", then.Nanosecond())
	p("Then Weekday: ", then.Weekday())
	p("Then Location: ", then.Location())

	p("Then Before now: ", then.Before(now))
	p("Then After now: ", then.After(now))
	p("Then Equal now: ", then.Equal(now))

	diff := then.Sub(now)
	p("Diff: ", diff)

	p("diff.Hours(): ", diff.Hours())
	p("diff.Minutes(): ", diff.Minutes())
	p("diff.Seconds(): ", diff.Seconds())
	p("diff.Nanoseconds(): ", diff.Nanoseconds())
	p("now.Add(diff): ", now.Add(diff))
	p("then.Add(-diff): ", then.Add(-diff))
	fmt.Println("")
}
