package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Krug struct{ R int }

func (k Krug) obim() float64 {
	return 2 * math.Pi * float64(k.R)
}

func main() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	for i := 0; i < 10; i++ {
		k := Krug{r.Intn(10) + 1}
		fmt.Printf("Obim kruga poluprecnika %d je %f\n", k.R, k.obim())
	}
}
