package main

import (
	"fmt"
	"math/rand"
)

func suma(niz []int, c chan int) {
	rezultat := 0
	for _, broj := range niz {
		rezultat += broj
	}
	c <- rezultat
}

func main() {
	r := rand.New(rand.NewSource(100))
	var niz [100]int
	var broj int
	for i := 0; i < 100; i++ {
		broj = r.Intn(101)
		niz[i] = broj
	}
	
	c := make(chan int)

	go suma(niz[len(niz)/2:], c)
	rez1, rez2 := <-c, <-c
	fmt.Println(rez1 + rez2)
}
