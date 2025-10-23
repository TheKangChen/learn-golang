package main

import (
	"fmt"
	"time"
)

func main() {
	// Get time
	t := time.Now()
	fmt.Println("Now: ", t)
	fmt.Println("Local Time: ", t.Local())
	fmt.Printf("%4d/%2d/%2d %2d:%2d:%2d:\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	t = time.Now().UTC()
	fmt.Println("UTC: ", t)

	// Get time from timezone
	loc, err := time.LoadLocation("Japan")
	if err != nil {
		fmt.Println("Error loading tz location:", err)
		return
	}
	fmt.Println("Japan time:", time.Now().In(loc))

	// Sleep
	time.Sleep(time.Second * 1)

	// Time delta
	fmt.Println("Elapsed time: ", time.Since(t))
	week := time.Hour * 24 * 7
	fmt.Println("A week from now:", t.Add(week))

	// Formatting time
	t = time.Now()
	fmt.Println("Formatted time:", t.Format("Jan 02 15:04"))
	fmt.Println(t.Format(time.ANSIC))
}
