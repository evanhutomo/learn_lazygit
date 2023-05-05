package main

import (
	"fmt"
	"os"

	"github.com/evilsocket/islazy/tui"
)

var (
	nama       = "Veloz"
	jumlahUnit = 1
)

type fn func(string, int)

func mobil(n string, i int) {
	fmt.Printf(tui.Bold("Mobil: "))
	fmt.Printf("%s\n", n)
	fmt.Printf(tui.Bold("Unit: "))
	fmt.Printf("%d\n", i)

	cols := []string{
		"City",
		"Temp",
	}

	rows := [][]string{
		{"Roma", "16 C"},
		{"Malang", "25 C"},
		{"Hamamatsu", "16 C"},
	}

	tui.Table(os.Stdout, cols, rows)
}

func test(f fn, a string, val int) {
	f(a, val)
}

func main() {
	test(mobil, nama, jumlahUnit)
}
