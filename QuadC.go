package main

import "fmt"

func QuadC(x, y int) {
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 || i == 1 && j == x {
				fmt.Print("A")
			} else if i == y && j == 1 || i == y && j == x {
				fmt.Print("C")
			} else if i == 1 || i == y {
				fmt.Print("B")
			} else if j == 1 || j == x {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
