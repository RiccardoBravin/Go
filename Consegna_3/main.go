package main

import (
	"consegna_3/src"
	"math/rand"
	"sync"
	"time"
)


func main() {
	fn.Timer()
	
	rand.Seed(time.Now().UnixNano())

	nomi := [] string{"Aldo","Giovanni","Giacomo"}


	//creazione e deposito degli strumenti
	martelli := make(chan fn.Martello, 1)
	martelli <- fn.Martello{}

	trapani := make(chan fn.Trapano, 2)
	trapani <- fn.Trapano{}
	trapani <- fn.Trapano{}
	
	cacciaviti := make(chan fn.Cacciavite, 1)
	cacciaviti <- fn.Cacciavite{}



	var wg sync.WaitGroup //waitgroup per gli operai

	wg.Add(3)

	for _,i := range nomi{
		go fn.Opera(fn.Operaio{Nome: i}, martelli, cacciaviti, trapani, &wg)
	}
	

	wg.Wait()
	wg.Wait()
	wg.Wait()

}
