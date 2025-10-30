package main

import (
	"fmt"
)

var gallery []string = make([]string, 0, 10)

// Upload image function
func uploadImage(filename string) {
	if len(gallery) == 10 {
		fmt.Println("Gallery full")
		return
	}
	gallery = append(gallery, filename)
}

// Delete image function
func deleteImage(filename string) {
	if len(gallery) == 0 {
		fmt.Println("No image to remove")
	}
	for i, v := range gallery {
		if v == filename {
			copy(gallery[i:], gallery[i+1:])
			gallery = gallery[:len(gallery)-1]
			fmt.Printf("%s removed\n", filename)
			return
		}
	}
	fmt.Printf("%s not found\n", filename)
}

// List images function
func listImages() {
	if len(gallery) == 0 {
		fmt.Println("Gallery is empty")
		return
	}
	for _, v := range gallery {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	fmt.Println("Testing Image Gallery Management System")
	uploadImage("image1.jpg")
	uploadImage("image2.jpg")
	uploadImage("image3.jpg")
	listImages()
	deleteImage("image2.jpg")
	listImages()
	uploadImage("image4.jpg")
	listImages()
	deleteImage("image1.jpg")
	deleteImage("image2.jpg")
	deleteImage("image3.jpg")
	listImages()
}
