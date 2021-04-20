package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Gruppo struct {
	nome     string
	nPalline int
}

type Tunnel struct {
	//libero bool
}

func transumanza(g Gruppo, t chan Tunnel, wg *sync.WaitGroup) {
	for g.nPalline > 0 {
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		fmt.Println("mandaPersona", g.nome)
		wg.Add(1)
		mandaPersona(&g, t)
		wg.Done()
		wg.Wait()
	}
}

func mandaPersona(g *Gruppo, t chan Tunnel) {

	//ho riscritto questa funzione perchè ritenevo superfluo il channel c1

	select {
	case <-t: //se riesco a prendere una pallina dal tunnel allora è avvenuto uno scontro
		fmt.Println("La pallina", g.nome, "si è scontrata")
	case t <- Tunnel{}: //se riesco a mettere una pallina nel channel allora:
		select {
		case t <- Tunnel{}: //se riesco a mettere una pallina nel tunnel quando ne avevo già inserita una allora vuol dire che la mia si è scontrata quindi mi fermo
			<-t
			fmt.Println("La pallina", g.nome, "si è scontrata")
		case <-time.After(time.Second): //se riesco ad attendere un secondo allora la mia pallina è arrivata dall'altra parte
			<-t
			g.nPalline-- //tolgo una pallina dal gruppo
			fmt.Println("La pallina in arrivo da", g.nome, "è passata, ne rimangono", g.nPalline)
		}
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	gruppo1 := Gruppo{"destra", 5}
	gruppo2 := Gruppo{"sinistra", 5}

	tunnelChannel := make(chan Tunnel, 1) //al posto di uno struct tunnel basta ovviamente un buleano ma volevo mantenere la presenza del tipo tunnel facendo finta che questo contenga o meno la pallina

	var wg sync.WaitGroup

	go transumanza(gruppo2, tunnelChannel, &wg)
	go transumanza(gruppo1, tunnelChannel, &wg)

	time.Sleep(time.Minute) //si potrebbe ottimizzare per terminare quando i gruppi sono vuoti
}
