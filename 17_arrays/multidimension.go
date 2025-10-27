package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 2-D Array: [row][col]Type
	var array2d [2][3]string 
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			array2d[i][j] = strconv.Itoa(i) + "," + strconv.Itoa(j)
			fmt.Print(array2d[i][j], " ")
		}
		fmt.Print("\n")
	}
	var _ [2][3][4]int // 3d array: [row][col][depth]
}
