package main

import (
	"fmt"  
	"math" 
)

type Tacka struct{ X, Y int }

func rastojanjeTacaka(t1, t2 Tacka) float64 {
	return math.Sqrt(float64((t1.X-t2.X)*(t1.X-t2.X) + (t1.Y-t2.Y)*(t1.Y-t2.Y)))
}

func (t Tacka) rastojanje() float64 {
	return math.Sqrt(float64(t.X*t.X + t.Y*t.Y))
}

func main() {

	var X int
	X = 10

	Y := 5

	t1 := Tacka{X, Y}
	t2 := Tacka{-9, 10}

	fmt.Println(t1, t2)

	fmt.Printf("Udaljenost tacaka je %f.\n", rastojanjeTacaka(t1, t2))

	r1 := t1.rastojanje()
	r2 := t2.rastojanje()
	if r1 < r2 {
		fmt.Println("Tacka ", t1, " je bliza koordinatnom pocetku.")
	} else if r1 == r2 {
		fmt.Println("Tacke su podjednako udaljene od koordinatnog pocetka.")
	} else {
		fmt.Println("Tacka ", t2, " je bliza koordinatnom pocetku.")
	}
}
