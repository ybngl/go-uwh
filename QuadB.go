package main

import "fmt"

func QuadB(x, y int) {
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 || i == y && j == x && x > 1 && y > 1 {
				fmt.Print("/")
			} else if i == 1 && j == x || i == y && j == 1 {
				fmt.Print("\\")
			} else if i == 1 || i == y {
				fmt.Print("*")
			} else if j == 1 || j == x {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
