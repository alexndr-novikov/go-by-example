package structs

import "fmt"

type user struct {
	email string
	id    int
}

func newUser(email string) *user {
	u := user{email: email}
	u.id = 1
	return &u
}

func Structs() {
	fmt.Println("Structs package output:")
	fmt.Println(user{"test@example.com", 2})
	fmt.Println(user{email: "test@example.com", id: 3})
	fmt.Println(user{id: 4, email: "test@example.com"})
	fmt.Println(user{id: 5})
	fmt.Println(&user{email: "addr@example.com", id: 6})
	fmt.Println(newUser("hello@world.com"))

	user1 := newUser("new@user.com")
	fmt.Println(user1)
	fmt.Println(*user1)
	fmt.Println(user1.id)
	user1.id = 100
	fmt.Println(user1.id)
	fmt.Println("")
	user2 := user{"test@example.com", 2}
	user2.id = 3
	fmt.Println(user2.id)
}
