package methods

import "fmt"

type user struct {
	name string
	age  int
}

func (u user) ageInMonths() int {
	return u.age * 12
}

func (u *user) happyBirthday(times int) {
	u.age = u.age + times
	if u.age > 35 {
		u.name = fmt.Sprintf("Mr. %s", u.name)
	}
}

func RunSample() {
	fmt.Println("Methods package output:")
	user1 := user{"John", 31}
	fmt.Println("Name:", user1.name)
	fmt.Println("Age:", user1.age)
	fmt.Println("Age in months:", user1.ageInMonths())
	fmt.Println("5 years later....")
	user1.happyBirthday(5)
	fmt.Println("Name:", user1.name)
	fmt.Println("Age:", user1.age)
	fmt.Println("Age in months:", user1.ageInMonths())
	fmt.Println("")
}
