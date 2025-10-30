package main

import (
	"fmt"
	"strings"
)

func main() {
	// Since slices are references to the underlying array, it could lead to the garbage collector keeping a large array in memory when sth is referencing it, even though you might only need small parts of it.

	fruits := []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew melon", "imbe", "jackfruit", "kiwi", "lemon", "mango", "nectarine", "orange", "pineapple", "quince", "raspberry", "strawberry", "tomato", "ugli fruit", "vanilla bean", "watermelon", "xigua", "yuzu", "zucchini"}

	// One way to avoid this problem is to simply copy the data you need and use the copied data
	fruitsUsed := make([]string, 10)
	copy(fruitsUsed, fruits[2:6])
	fmt.Println(fruitsUsed)

	var berries []string
	for _, f := range fruits {
		if strings.Contains(f, "berry") {
			berries = append(berries, f)
		}
	}
	fmt.Println(berries)
}
