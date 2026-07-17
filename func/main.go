package main

import (
	"time"
)

func closure(tim int) func(...int) int {
	return func(intlist ...int) int {
		time.Sleep(time.Duration(tim) * time.Second)
		var sum int
		for _, item := range intlist {
			sum += item
		}
		return sum
	}
}

// func main() {
// 	t1 := time.Now()
// 	sum := closure(2)(1, 2, 3)
// 	subTime := time.Since(t1)
// 	fmt.Println(sum, subTime)

// }
