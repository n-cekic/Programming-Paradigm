// Napisati konkurentan program kojim se izračunava proizvod n-dimenzionalnog
// vektora skalarom.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup


var mutex sync.Mutex


var vektor []int
var n int 
var i int 
var skalar int


func mnozenje(rb int) {
	
	var index int 
	for {
		
		mutex.Lock()
		if i == n {
			mutex.Unlock()
			break
		} else {
			index = i
			i += 1
		}
		mutex.Unlock()
		fmt.Printf("Gorutina %d mnozi el. na poziciji %d\n", rb, index)
		vektor[index] *= skalar
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
	for rg := 0; rg < brojGR; rg++ {
		go mnozenje(rg)
	}
	wg.Wait()
	fmt.Println(vektor)
}
