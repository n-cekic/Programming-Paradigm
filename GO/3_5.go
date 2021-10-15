package main

import (
	"fmt"
)

func zbir(x, y int, c chan int) {
	c <- x + y
}
func proizvod(x, y int, c chan int) {
	c <- x * y
}

func main() {
	c := make(chan int, 2)

	go zbir(2, 36, c)
	go zbir(15, 100, c)

	suma1, suma2 := <-c, <-c
	go proizvod(suma1, suma2, c)

	fmt.Printf("Vrednost izraza je %d\n", <-c)
}
