// AvdentOfCode project main.go
package main

import (
	"flag"
	"fmt"
)

func main() {
	dayOneCode := ""
	flag.StringVar(&dayOneCode, "code", "", "Enter the code for day one")
	flag.Parse()
	if len(dayOneCode) > 0 {
		floor, pos := dayOne(dayOneCode)
		fmt.Println("Day one solutions floor =", floor, "pos when first in basement =", pos)
	}
}

func dayOne(code string) (floor int, pos int) {
	pos = -1
	for i, v := range code {
		switch v {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 && pos == -1 {
			pos = i + 1
		}
	}
	return
}
