package main

import (
	"consegna_1/src"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	clients := [7] string{"Anna","Barbara","Cinzia","Debora","Elena","Filippo","Giorgia"}

	v1 := make(chan fn.Viaggio,1)
	v1 <- fn.Viaggio{
		Meta: "Spagna",
		Prenotati: make([]fn.Cliente,0),
		Minimum: 4,
	}

	v2 := make(chan fn.Viaggio,1)
	v2 <- fn.Viaggio{
		Meta: "Francia",
		Prenotati: make([]fn.Cliente,0),
		Minimum: 2,
	}

	var wg sync.WaitGroup
	

	wg.Add(7)

	for i := 0; i<7; i++{
		cl := fn.Cliente{Nome: clients[i]}
		go fn.Prenota(cl, v1, v2, &wg)
	}

	wg.Wait()
	
	fn.StampaPartecipanti(<-v1,<-v2)

}
