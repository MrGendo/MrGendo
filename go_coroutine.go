package main

import (
	"fmt"
	"program/even"
)

func main() {
	i := even.Feibo(10)
	fmt.Println(i)
}

func mapFUNC(f func(i int) int, s []int) []int {
	for i := range s {
		s[i] = f(s[i])
	}
	return s
}

func chu(i int) int {
	return i / 10
}
