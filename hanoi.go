package main

import "fmt"

func moveDisk(n, from, to, via int) {
	if n <= 0 {
		return
	}
	moveDisk(n-1, from, via, to)
	fmt.Println(from, "->", to)
	moveDisk(n-1, via, to, from)
}

func hanoi(n int) {
	fmt.Println("Number of disks:", n)
	moveDisk(n, 1, 3, 2)
}

func main() {
	hanoi(3)
}
