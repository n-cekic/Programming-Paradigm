package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

var vektor []int
var n int
var skalar int


func mnozenjeIsecak(l, r, rb int) {
	for i := l; i < r; i++ {
		fmt.Printf("Gorutina %d mnozi el. na poziciji %d\n", rb, i)
		vektor[i] *= skalar
	}
	wg.Done()
}


func mnozenjeOstatak(rb, brojGR int) {
	for i := rb; i < n; i += brojGR {
		fmt.Printf("Gorutina %d mnozi el. na poziciji %d\n", rb, i)
		vektor[i] *= skalar
	}
	wg.Done()
}

func main() {
	fmt.Scanf("%d", &n)
	vektor = make([]int, n)
	
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &vektor[i])
	}
	
	fmt.Println(vektor)
	fmt.Scanf("%d", &skalar)

	brojGR := runtime.NumCPU()
	wg.Add(brojGR)

	kolicnik := n / brojGR
	ostatak := n % brojGR
	
	var duzinaIsecka int

	if ostatak > 0 {
		duzinaIsecka = kolicnik + 1
		ostatak--
	} else {
		duzinaIsecka = kolicnik
	}
	
	for l, r, rg := 0, duzinaIsecka, 0; rg < brojGR; rg++ {
		go mnozenjeIsecak(l, r, rg)
		l = r
		if ostatak <= 0 {
			duzinaIsecka = kolicnik
		} else {
			ostatak--
		}
		r += duzinaIsecka
	}
	wg.Wait()
	fmt.Println(vektor)
}
