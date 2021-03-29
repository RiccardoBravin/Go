package main

import (
	fn "consegna_2/src"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	piatti := []string{"Anatra", "Broccoletti", "Calzone", "Daino", "Eclairs", "Fondue", "Gelato", "Humus", "Insalata", "Linguine"}

	orders := make(chan fn.Piatto, 10)
	cooked := make(chan fn.Piatto, 10)

	//funzione che stampa il tempo ogni secondo
	go fn.Timer()

	//ordinazione di tutti i piatti disponibili
	for i := range piatti {
		fn.Ordina(fn.Piatto{Nome: piatti[i]}, orders)
	}
	close(orders)

	//invio in cottura dei piatti
	go fn.Cucina(orders, cooked)

	//consegna dei piatti
	fn.Consegna(cooked)

}
