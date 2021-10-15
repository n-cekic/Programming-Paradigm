package main

import (
	"fmt"
	"io" // sadrzi EOF
	"os" // sadrzi funkcije za rad sa datotekama
)

func suma(niz []int) int {
	rezultat := 0
	for _, broj := range niz {
		rezultat += broj
	}
	return rezultat
}

func main() {
	// poruku o greski dobijamo preko druge povratne vrednosti
	file, err := os.Open("brojevi.txt")
	// ukoliko je uspesno otvorena datoteka err je nil
	if err != nil {
		fmt.Println("Greska prilikom otvaranja datoteke")
		os.Exit(1)
	}
	var broj int
	var niz []int
	for {
		_, err := fmt.Fscanf(file, "%d\n", &broj)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Lose formatirana datoteka")
			os.Exit(1)
		}
		niz = append(niz, broj)
	}
	fmt.Println(niz)
	fmt.Println(suma(niz))
}
