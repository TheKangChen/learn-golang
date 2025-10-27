package main

import "fmt"

func main() {
	// 2d slice: [row][col]Type
	slice2d := [][]int{}

	row1 := []int{11, 12, 13}
	row2 := []int{21, 22, 23}

	slice2d = append(slice2d, row1)
	slice2d = append(slice2d, row2)

	for i := 0; i < 2; i++ {
		fmt.Printf("Row %d\n", i+1)
		fmt.Println(slice2d[i])
	}

	fmt.Println("First element")
	fmt.Println(slice2d[0][0])

	fmt.Println("Values")
	fmt.Println(slice2d)
	var _ [][][]int // 3d slice: [row][col][depth]
}
