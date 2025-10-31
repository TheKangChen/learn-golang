package main

import "fmt"

func main() {
	// Maps are an unordered collection of key value pairs that grows dynamically in size
	// (hashtables, hashmap, dictionaries, ...)

	// CREATION
	m1 := map[string]int{"apple": 10, "banana": 20}
	var _ map[string]int        // this is a nil map
	var _ = new(map[string]int) // this will return a nilpointer, don't use new

	// To create an empty map, use make
	m2 := make(map[string]string) // make(map[keytype]valtype, initcapacity)

	// value type: func() int
	var mf = map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
	}

	// value type: slices/pointers
	var _ = make(map[int][]int)
	var _ = make(map[int]*[]int)

	// value type: maps
	var _ = make(map[int]map[string]string)

	// ASSIGNING VALUES
	m2["lastname"] = "Bond"
	m2["firstname"] = "James"
	m2["codename"] = "007"

	// Slice of maps
	s := make([]map[int]int, 5)
	for i := range s {
		s[i] = make(map[int]int)
		s[i][i*2] = 10 * i
	}
	s[0][20] = 400
	fmt.Println(s)

	// ACCESSING VALUES
	v := m1["apple"] // existing value
	fmt.Println(v)

	v = m1["kiwi"] // non existing value will return the zero value of the value type
	fmt.Println(v)

	// TESTING KEY EXISTENCE: use `val, ok` pattern
	v1, ok := m1["apple"]
	if ok {
		fmt.Println("Value for apple:", v1)
	}

	if _, ok := m1["mango"]; ok {
		fmt.Println("mango exist")
	} else {
		fmt.Println("mango doesn't exist")
	}

	// DELETING ITEMS
	delete(m1, "banana")
	fmt.Println(m1)

	// FOR RANGE: returns key & value
	// Note: Order is not guaranteed
	for k, f := range mf {
		fmt.Println(k, ":", f())
	}

	// SORTING A MAP
	// To sort the order of a map, copy all the keys to a slice and sort the slice -> use the slice to get the map values
	keys := make([]string, len(m2))
	i := 0
	for k := range m2 {
		keys[i] = k
		i++
	}
	for _, k := range keys {
		fmt.Println(k, ":", m2[k])
	}

	// INVERTING A MAP
	m2i := make(map[string]string, len(m2))
	for k, v := range m2 {
		m2i[v] = k
	}
	fmt.Println(m2i)
}
