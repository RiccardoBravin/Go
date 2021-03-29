package main

import (
	"consegna_2/src"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	piatti := [] string{"Anatra","Broccoletti","Calzone","Daino","Eclairs","Fondue","Gelato","Hmm","Insalata","Linguine"}

	orders := make(chan fn.Piatto, 10)
	cooked := make(chan fn.Piatto, 10)

	
	go func(){
		start := time.Now()
		prev := time.Duration(0)
		fmt.Println(prev)
		for{
			now := time.Now().Sub(start).Truncate(time.Second)
			if(now > prev){
				prev = now
				fmt.Println("Sono passati", now)
			}
		}
	}()

	for i := range piatti{
		fn.Ordina(fn.Piatto{piatti[i]}, orders)		
	}	
	close(orders)
	

	fn.Cucina(orders, cooked)

	fn.Consegna(cooked)
	
	




	
	/*
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
	*/

}
