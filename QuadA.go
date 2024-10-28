package main

import "fmt"

func QuadA(x, y int) {
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && (j == 1 || j == x) || i == y && (j == 1 || j == x) {
				fmt.Print("o")
			} else if i == 1 || i == y {
				fmt.Print("-")
			} else if j == 1 || j == x {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
