package fn

import (
	"fmt"
	"math/rand"
	"sync"
)

//Viaggio contiene il nome della meta e uno slice contenente i clienti
type Viaggio struct{
	Meta string
	Prenotati []Cliente
	Minimum int
}


//NewViaggio crea un nuovo viaggio a partire 
func NewViaggio(_meta string, _minimum int)(Viaggio){
	return Viaggio{
		Meta: _meta,
		Prenotati: make([]Cliente,0),
		Minimum: _minimum,
	}
}

//Cliente è una struttura contenente i dati di un determinato cliente
type Cliente struct{
	Nome string

}

//Prenota effettua una prenotazione
func Prenota(c Cliente, v1 chan Viaggio, v2 chan Viaggio, wg *sync.WaitGroup){
	
	if rand.Intn(2) == 0{
		aux := <- v1
		aux.Prenotati = append(aux.Prenotati, c)
		v1 <- aux
	}else{
		aux := <- v2
		aux.Prenotati = append(aux.Prenotati, c)
		v2 <- aux
	}
	wg.Done()

}

//StampaPartecipanti stampa i viaggi prenotati e i clienti che parteciperanno
func StampaPartecipanti(v1 Viaggio, v2 Viaggio){
	if len(v1.Prenotati) >= v1.Minimum{
		fmt.Println("Il viaggio per la", v1.Meta, "partirà con", v1.Prenotati)	
	}
	if len(v2.Prenotati) >= v2.Minimum{
		fmt.Println("Il viaggio per la", v2.Meta, "partirà con", v2.Prenotati)	
	}
}

