package random

import (
	"fmt"
	"math/rand"
	"time"
)

func RunSample() {
	p := fmt.Println
	p("Random package output:")
	p(rand.Intn(100), ",", rand.Intn(100))
	p(rand.Intn(100), ",", rand.Intn(100))
	p()
	p(rand.Float64())
	s1 := rand.NewSource(time.Now().UnixNano())
	t1 := rand.New(s1)
	p(t1.Intn(100))
	p(t1.Intn(100))

	p()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	p(r2.Intn(100), ",")
	p(r2.Intn(100))
	p()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	p(r3.Intn(100), ",")
	p(r3.Intn(100))

	p("")
}
