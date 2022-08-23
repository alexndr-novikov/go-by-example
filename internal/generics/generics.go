package generics

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
		fmt.Println(k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}
func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func RunSample() {
	fmt.Println("Generics package output:")
	var m = map[int]string{1: "One", 2: "Two", 3: "Three"}
	fmt.Println(MapKeys(m))
	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(100)
	lst.Push(1000)
	fmt.Println("list:", lst.GetAll())

	lst2 := List[string]{}
	lst2.Push("a")
	lst2.Push("b")
	lst2.Push("c")
	fmt.Println("list:", lst2.GetAll())

	fmt.Println("")
}
