// Postoji veliko skladište čokolade koje ima 100 hiljada
// čokoladica. Pored skladišta postoje 4 prodavnice u koje može da se smesti najviše 1000 čokoladica.
// Inicijalno svaka ima 100 čokoladica. Prodavnice se snabdevaju periodično iz skladišta na svakih
// 5 sekundi. One uvek uzmu toliko čokoladica da se dopune do maksimalnog kapaciteta. Za svaku
// prodavnicu kreirati zasebnu gorutinu koja obavlja dopunu njenih zaliha i upravlja kupcima. Postoji
// i 100 ljudi od kojih svako na 10 sekundi oseti potrebu za čokoladicom i odlazi do neke nasumično
// odabrane prodavnice. Potom u prodavnici, ako je dostupno, uzme nasumično od 1 do 7 čokoladica
// ili čeka u prodavnici dok se ne dopune zalihe. Za svaku osobu kreirati zasebnu gorutinu koja obavlja
// odlazak u prodavnicu i kupovinu određenog broja čokoladica. Modelovati sistem i prikazivati
// informacije kad god neko uđe u prodavnicu da kupi čokoladice i kad se obavi kupovina

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Covek struct {
	id       int
	cokoSvet *CokoladniSvet
}

type Skladiste struct {
	stanje int
}

type Prodavnica struct {
	id        int
	stanje    int
	kapacitet int
	kupovinaP sync.Mutex
	kupovina  *sync.Cond
	cokoSvet  *CokoladniSvet
}

type CokoladniSvet struct {
	skladiste  Skladiste
	prodavnice [4]Prodavnica
	ljudi      []Covek
	dopunaP    sync.Mutex
}

// funkcija koja inicijalizuje podatke o prodavnici
func napraviProdavnicu(p *Prodavnica, id int, kapacitet int, cokoSvet *CokoladniSvet) {
	p.id = id
	p.stanje = kapacitet / 10
	p.kapacitet = kapacitet
	p.cokoSvet = cokoSvet
	// svaka prodavnica ima svoj red cekanja za kupce
	p.kupovina = sync.NewCond(&p.kupovinaP)
}

// dopuna zaliha cokoladica u prodavnici
func dopuna(p *Prodavnica) {
	for {
		// ako je ispod minimuma za stanje, radi se dopuna
		if p.stanje < p.kapacitet/10 {
			potrebno := p.kapacitet - p.stanje
			// sve uzimaju iz istog skladista - deljeni resurs
			// ulazak u kriticnu sekciju
			p.cokoSvet.dopunaP.Lock()
			if p.cokoSvet.skladiste.stanje >= potrebno {
				p.cokoSvet.skladiste.stanje -= potrebno
				p.stanje = p.kapacitet
			}
			p.cokoSvet.dopunaP.Unlock()
			// ako je obavljena dopuna
			// signaliziramo kupcima koji cekaju na kupovinu
			p.kupovina.Broadcast()
		}
		time.Sleep(time.Second * 5)
	}
}

func kupovina(covek Covek) {
	for {
		// odredjujemo broj cokoladica i iz koje prodavnice kupujemo
		brCoko := rand.Intn(7) + 1
		brP := rand.Intn(4)
		// pozivamo funkciju koja obavlja kupovinu
		kupi(covek.id, brCoko, &covek.cokoSvet.prodavnice[brP])
		time.Sleep(time.Second * 10)
	}
}

func kupi(covek int, brCoko int, p *Prodavnica) {
	// kriticna sekcija - vise kupaca moze biti u jednoj prodavnici
	p.kupovina.L.Lock()
	fmt.Printf("Covek %d zeli da kupi %d cokoladica u prodavnici %d (stanje: %d)\n", covek, brCoko, p.id, p.stanje)
	// ako u prodavnici nema dovoljno cokoladica
	// kupac se stavlja u red cekanja
	for p.stanje < brCoko {
		p.kupovina.Wait()
	}
	// obavlja se kupovina
	p.stanje -= brCoko
	fmt.Printf("Covek %d je kupio %d cokoladica u prodavnici %d (stanje: %d)\n", covek, brCoko, p.id, p.stanje)
	p.cokoSvet.ukupnoCokoladica()
	p.kupovina.L.Unlock()
}

func napraviCokoSvet(cokoSvet *CokoladniSvet, kapacitetS int, kapacitetP int, brLjudi int) {
	cokoSvet.skladiste.stanje = kapacitetS
	cokoSvet.ljudi = make([]Covek, brLjudi)
	for i := 0; i < 4; i++ {
		napraviProdavnicu(&cokoSvet.prodavnice[i], i, kapacitetP, cokoSvet)
		go dopuna(&cokoSvet.prodavnice[i])
	}
	for i := 0; i < brLjudi; i++ {
		cokoSvet.ljudi[i].cokoSvet = cokoSvet
		cokoSvet.ljudi[i].id = i
		go kupovina(cokoSvet.ljudi[i])
	}

}

// ukupan broj cokoladica treba da se smanjuje svakom kupovinom
func (cokoSvet CokoladniSvet) ukupnoCokoladica() {
	stanje := 0
	for _, p := range cokoSvet.prodavnice {
		stanje += p.stanje
	}
	stanje += cokoSvet.skladiste.stanje
	fmt.Printf("Total: %d\n", stanje)

}

func main() {
	cokoSvet := CokoladniSvet{}
	wg.Add(1100)
	napraviCokoSvet(&cokoSvet, 100000, 1000, 100)
	wg.Wait()
}
