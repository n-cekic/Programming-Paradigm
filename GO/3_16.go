// Železnica nudi usluge prevoza robe. Vozovi kreću
// iz date stanice u nasumičnu sledeću samo ako su ili potpuno puni, ili potpuno prazni. Tada u
// slučaju da su potpuno puni, pokušavaju da istovare svu svoju robu u prvoj sledećoj stanici, inače
// pokušavaju da napune svoje vagone do punog kapaciteta. Postoji veza između svaka dva grada
// i pruga ima dva koloseka, odnosno, pruga nije deljeni resurs. Sa druge strane, stanice su deljeni
// resursi, iz razloga što u datom momentu više vozova može čekati na robu u njoj. Potrebno je kreirati
// četiri stanice i šest vozova koje inicijalno treba postaviti na različite polazne stanice. Vozovi će
// imati svoj kapacitet proizvoljno određen u opsegu od 5 do 10 tona. Takođe, inicijalno su svi prazni,
// dakle, moraće da uzmu robu iz polazne stanice. Za svaki voz kreirati zasebnu gorutinu koja obavlja
// prevoz robe. Stanice imaju proizvoljno određen kapacitet u opsegu od 10 do 15 tona i inicijalno su
// im skladišta puna. Modelovati sistem i prikazivati informacije kad god neki voz pristigne u stanicu
// i obavi preuzimanje ili predaju robe.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Voz struct {
	id        int
	kapacitet int
	pun       bool
	stanica   int
	zeleznica *Zeleznica
}

type Stanica struct {
	id              int
	stanje          int
	kapacitet       int
	katanac1        sync.Mutex
	katanac2        sync.Mutex
	preuzimanjeRobe *sync.Cond
	predajaRobe     *sync.Cond
	zeleznica       *Zeleznica
}

type Zeleznica struct {
	stanice [4]Stanica
	vozovi  [6]Voz
}

func prevozRobe(v *Voz) {
	for {
		if v.pun {
			prodaja(v, &v.zeleznica.stanice[v.stanica])
			// azuriramo u koju stanicu sada ide voz nakon razmene robe
			// na primer, svaki voz ide u narednu stanicu
			v.stanica = (v.stanica + 1) % 4
			time.Sleep(time.Second * 20)
		} else {
			kupovina(v, &v.zeleznica.stanice[v.stanica])
			v.stanica = (v.stanica + 1) % 4
			time.Sleep(time.Second * 20)
		}

	}

}

func prodaja(v *Voz, s *Stanica) {
	// voz je pun i pokusava da preda robu u stanici
	s.preuzimanjeRobe.L.Lock()
	fmt.Printf("Voz %d kapaciteta %d je stigao pun u stanicu %d (stanje %d)\n", v.id, v.kapacitet, s.id, s.stanje)
	for v.kapacitet+s.stanje > s.kapacitet {
		s.preuzimanjeRobe.Wait()
	}
	s.stanje += v.kapacitet
	v.pun = false
	fmt.Printf("Voz %d je predao %d tona robe u stanici %d (stanje %d)\n", v.id, v.kapacitet, s.id, s.stanje)
	s.preuzimanjeRobe.L.Unlock()
	// obavestavamo vozove koji cekaju da preuzmu robu da se promenilo stanje
	s.predajaRobe.Broadcast()
}

func kupovina(v *Voz, s *Stanica) {
	// voz je prazan i pokusava da preuzme robu iz stanice
	s.predajaRobe.L.Lock()
	fmt.Printf("Voz %d kapaciteta %d je stigao prazan u stanicu %d (stanje %d)\n", v.id, v.kapacitet, s.id, s.stanje)
	for s.stanje < v.kapacitet {
		s.predajaRobe.Wait()
	}
	s.stanje -= v.kapacitet
	v.pun = true
	fmt.Printf("Voz %d je preuzeo %d tona robe iz stanice %d (stanje %d)\n", v.id, v.kapacitet, s.id, s.stanje)
	s.predajaRobe.L.Unlock()
	// obavestavamo vozove koji cekaju da predaju robu da se oslobodio deo skladista
	s.preuzimanjeRobe.Broadcast()
}

func napraviZeleznicu(zeleznica *Zeleznica) {
	for i := 0; i < 4; i++ {
		zeleznica.stanice[i].zeleznica = zeleznica
		kapacitet := rand.Intn(6) + 10
		zeleznica.stanice[i].kapacitet = kapacitet
		zeleznica.stanice[i].stanje = kapacitet
		zeleznica.stanice[i].id = i
		zeleznica.stanice[i].predajaRobe = sync.NewCond(&zeleznica.stanice[i].katanac1)
		zeleznica.stanice[i].preuzimanjeRobe = sync.NewCond(&zeleznica.stanice[i].katanac2)
	}
	for i := 0; i < 6; i++ {
		kapacitet := rand.Intn(6) + 5
		zeleznica.vozovi[i].zeleznica = zeleznica
		zeleznica.vozovi[i].kapacitet = kapacitet
		zeleznica.vozovi[i].id = i
		zeleznica.vozovi[i].pun = false
		// rasporedjujemo vozove na razlicite stanice
		zeleznica.vozovi[i].stanica = i % 4
		go prevozRobe(&zeleznica.vozovi[i])
	}
}

func main() {
	zeleznica := Zeleznica{}
	wg.Add(6)
	napraviZeleznicu(&zeleznica)
	wg.Wait()
}
