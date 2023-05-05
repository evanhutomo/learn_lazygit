package main

import "fmt"

type fn func(int)

func iniFunc(i int) {
	fmt.Printf("angkanya : %d", i)
}

func test(f fn, val int) {
	f(val)
}

func main() {
	test(iniFunc, 90)
}
