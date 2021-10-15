// Napisati konkurentan program kojim se izračunava skalarni proizvod dva n-
// dimenzionalna vektora.

package main

import (
	"fmt"
	"os"      // sadrzi Args - niz argumenata komandne linije
	"strconv" // sadrzi Atoi funkciju
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

// deljeni podaci
var vektor1 []int
var vektor2 []int
var n int      // dimenzija vektora
var i int      // prva slobodna neobradjena pozicija vektora
var skalar int // rezultat

func skalarniP(rb int) {
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
			// azuriramo podatak
			i += 1
			// otkljucavamo sekciju
			mutex.Unlock()
		}
		element := vektor1[index] * vektor2[index]
		fmt.Printf("Gorutina %d mnozi el. na poziciji %d\n", rb, index)
		// azuriramo deljeni podatak - kriticna sekcija
		mutex.Lock()
		skalar += element
		mutex.Unlock()
	}
	wg.Done()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Niste zadali broj gorutina.")
		os.Exit(1)
	}
	brojGR, _ := strconv.Atoi(os.Args[1])
	fmt.Scanf("%d", &n)
	vektor1 = make([]int, n)
	vektor2 = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &vektor1[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &vektor2[i])
	}
	fmt.Println(vektor1, vektor2)

	wg.Add(brojGR)
	for rg := 0; rg < brojGR; rg++ {
		go skalarniP(rg)
	}
	wg.Wait()
	fmt.Println(skalar)
}
