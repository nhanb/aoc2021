package main

import "fmt"

func main() {
	ar := []int{2, 3}
	for i, _ := range ar {
		ar = append(ar, i)
		fmt.Println(ar)
	}
}
