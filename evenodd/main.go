package main

import "fmt"

type arr []int

func main() {
	myArr := newArr()
	myArr.print()
}

func newArr() arr {
	newArr := arr{}
	newArr = append(newArr, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	return newArr
}
func (a arr) print() {
	for _, val := range a {
		if val%2 == 0 {
			fmt.Println(val, "is even")
		} else {
			fmt.Println(val, "is odd")
		}
	}
}
