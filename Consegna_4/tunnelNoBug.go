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

func transumanza(g *Gruppo, t chan bool, wg *sync.WaitGroup) {

	time.Sleep(time.Duration(rand.Intn(2)) * time.Second * 5 / 4)
	fmt.Println("mandaPersona", g.nome)
	mandaPersona(g, t)
	wg.Done()

}

func mandaPersona(g *Gruppo, t chan bool) {

	select {
	case t <- true: //se riesco a mettere una pallina nel channel allora:
		select {
		case t <- true: //se riesco a mettere una pallina nel tunnel quando ne avevo già inserita una allora vuol dire che la mia si è scontrata quindi mi fermo
			<-t
			fmt.Println("La pallina", g.nome, "si è scontrata")
		case <-time.After(time.Second): //se riesco ad attendere un secondo allora la mia pallina è arrivata dall'altra parte
			<-t
			g.nPalline-- //tolgo una pallina dal gruppo
			fmt.Println("La pallina in arrivo da", g.nome, "è passata, ne rimangono", g.nPalline)
		}
	default:
		<-t
		fmt.Println("La pallina", g.nome, "si è scontrata")
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	gruppo1 := Gruppo{"destra", 5}
	gruppo2 := Gruppo{"sinistra", 5}

	tunnelChannel := make(chan bool, 1) //al posto di uno struct tunnel basta ovviamente un buleano ma volevo mantenere la presenza del tipo tunnel facendo finta che questo contenga o meno la pallina

	var wg sync.WaitGroup

	for gruppo1.nPalline > 0 && gruppo2.nPalline > 0 {
		wg.Add(2)
		go transumanza(&gruppo2, tunnelChannel, &wg)
		go transumanza(&gruppo1, tunnelChannel, &wg)
		wg.Wait()

	}
	for gruppo1.nPalline > 0 {
		wg.Add(1)
		transumanza(&gruppo1, tunnelChannel, &wg)
		wg.Wait()

	}
	for gruppo2.nPalline > 0 {
		wg.Add(1)
		transumanza(&gruppo2, tunnelChannel, &wg)
		wg.Wait()

	}

}
