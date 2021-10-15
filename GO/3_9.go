// Napisati konkurentan program kojim se proverava koji su elementi niza prosti
// brojevi i ispisuje ih na standardni izlaz. U datoteci brojevi.txt je u prvoj liniji zadat broj
// elemenata, a zatim i elementi niza.

package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

// deljeni podaci
var vektor []int
var n int // dimenzija vektora
var i int // prva slobodna neobradjena pozicija vektora

// funkcija proverava da li je broj prost ili ne
func prost(broj int) bool {
	for delilac := 2; delilac <= int(math.Sqrt(float64(broj))); delilac++ {
		if broj%delilac == 0 {
			return false
		}
	}
	return true
}

func prosti(rb int) {
	var index int
	for {
		// zakljucavanje kriticne sekcije
		mutex.Lock()
		// provera da li se doslo do kraja
		if i == n {
			mutex.Unlock()
			break
		} else {
			index = i
			i += 1
			// otkljucavamo sekciju
			mutex.Unlock()
		}
		// proveravamo da li je broj prost
		if prost(vektor[index]) == true {
			fmt.Printf("Gorutina %d je pronasla prost broj %d\n", rb, vektor[index])
		}
	}
	wg.Done()
}

func main() {
	file, err := os.Open("brojevi.txt")
	if err != nil {
		fmt.Println("Greska prilikom otvaranja datoteke")
		// prekidamo program
		os.Exit(1)
	}
	fmt.Fscanf(file, "%d\n", &n)
	vektor = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(file, "%d\n", &vektor[i])
	}
	i = 0
	fmt.Println(vektor)

	// broj raspolozivih procesora
	brojGR := runtime.NumCPU()

	wg.Add(brojGR)
	for rg := 0; rg < brojGR; rg++ {
		go prosti(rg)
	}
	// glavna gorutina ne sme da zavrsi
	// dok pomocne ne prodju kroz sve brojeve i ispisu ih
	// inace se moze desiti da ne vidimo ispis svih prostih brojeva
	wg.Wait()
}
