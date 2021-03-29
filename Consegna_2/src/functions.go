package fn

import (
	"fmt"
	"math/rand"
	"time"
)

//Piatto contiene il nome del piatto
type Piatto struct {
	Nome string
}

//Cameriere contiene il nome del cameriere e il piatto che sta trasportando
type Cameriere struct {
	Nome string
}

func Ordina(p Piatto, orders chan Piatto) {
	orders <- p
}

//Cucina manda in cottura i piatti ordinati. La cottura di ogni piatto impiega tra i 4 e i 6 secondi
func Cucina(orders chan Piatto, cooked chan Piatto) {
	gasLock := make(chan bool, 3) //channel per il lock dei fornelli

	gasLock <- true
	gasLock <- true
	gasLock <- true

	for plate := range orders {
		go func(p Piatto) { //funzione che effettua la cottura del piatto ordinato passato come parametro se ci sono fornelli disponibili
			<-gasLock
			fmt.Println("Il piatto", p, "sta venendo cucinato")
			time.Sleep(time.Second * time.Duration(rand.Intn(3)+4))
			fmt.Println("Il piatto", p, "ha terminato la cottura")
			gasLock <- true
			cooked <- p
		}(plate)
	}

}

//Consegna definisce tre camerieri e gestisce la consegna (in tre secondi) dei piatti nel channel passato per parametro. Termina se non viene consegnato un piatto da più di 15 secondi dato che tempo di cottura più consegna è inferiore
func Consegna(cooked chan Piatto) {
	waiters := make(chan Cameriere, 2) //definizione dei camerieri
	waiters <- Cameriere{Nome: "Stanlio"}
	waiters <- Cameriere{Nome: "Ollio"}

	for {
		select {
		case p := <-cooked: //se ci sono piatti da consegnare allora un cameriere se libero lo porta
			go func(p Piatto) { //funzione che gestisce la consegna in parallelo
				w := <-waiters
				fmt.Println("Il piatto", p, "sta venendo consegnato da", w.Nome)
				time.Sleep(time.Second * 3)
				fmt.Println("Il piatto", p, "è stato consegnato da", w.Nome)
				waiters <- w
			}(p)

		case <-time.After(15 * time.Second):
			return
		}

	}

}

//Timer è una funzione che stampa il tempo ogni secondo (lanciare con go)
func Timer() {
	start := time.Now()      //tempo inizio programma
	prev := time.Duration(0) //valore di tempo precedente
	fmt.Println(prev)
	for {
		now := time.Since(start).Truncate(time.Second) //effettua la differenza tra il tempo attuale e quello di inzio troncando ai secondi
		if now > prev {                                //se è passato un secondo allora stampa e aggiorna prev
			prev = now
			fmt.Println("Sono passati", now)
		}
	}
}
