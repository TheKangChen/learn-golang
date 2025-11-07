package main

import (
	"fmt"
	"unsafe"
)

// Public struct
type User struct {
	name  string
	age   int
	email string
}

// Private struct
type point struct {
	x int
	y int
}

// Struct of structs
type rect struct {
	min, max point
}

// Recursive struct
type node struct {
	value int
	next  *node
}

func main() {
	// Struct is a value type
	var u1 User
	u1.name = "Bond"
	u1.age = 20
	u1.email = "jamesbond007@example.com"
	fmt.Println(u1)

	// Initializing
	u2 := User{name: "Kent", age: 200, email: "clarkkent@example.com"}
	fmt.Println(u2)

	u3 := User{name: "Bob"} // The rest of the fields will the the corresponding datatype's zero value
	fmt.Println(u3)

	// Create a pointer to a struct:
	var u4 *User
	u4 = &User{name: "Tim", age: 50, email: "timburton@example.com"}
	fmt.Println(u4)

	p1 := new(point)
	p1.x = 20    //Automatic dereferenced
	(*p1).y = 10 // This is probably technically the more correct way to do it but golang does automatic dereferencing
	fmt.Println(p1)

	// Linked list
	n1 := new(node)
	var n2 *node
	n1.value = 1
	n1.next = &node{value: 2, next: n2}
	n2 = &node{value: 3}
	fmt.Println(n1)

	// Size of a struct is the combined size of the contents
	fmt.Println("Size of point:", unsafe.Sizeof(point{}))
	fmt.Println("Size of rect:", unsafe.Sizeof(rect{}))
}
