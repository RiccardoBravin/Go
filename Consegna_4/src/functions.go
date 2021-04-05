package fn

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Operaio struct {
	Nome string
}

type Martello struct {
}

type Cacciavite struct {
}

type Trapano struct {
}


func Opera(o Operaio, m chan Martello, c chan Cacciavite, t chan Trapano, wg *sync.WaitGroup){

	select{
	case obj1 := <- t: //se riesco a prendere il trapano allora potrò cercare di prendere cacciavite o martello (t, , )
		obj1.use(o)
		t <- obj1

		select{//ho preso il trapano quindi vedo cosa posso prendere tra cacciavire e martello e completo con l'ultimo di conseguenza 
		case obj1 := <- c: //prendo un cacciavite (t,c, )
			obj1.use(o)
			c <- obj1

			obj2 := <- m //attendo un martello (t,c,m)
			obj2.use(o)
			m <- obj2

			wg.Done() //lavoro terminato
		
		case obj1 := <- m: //prendo un martello (t,m, )
			obj1.use(o)
			m <- obj1

			obj2 := <- c //attendo un cacciavite (t,m,c)
			obj2.use(o)
			c <- obj2
			
			wg.Done() //lavoro terminato

		}

	case obj1 := <- m : //se riesco a prendere subito il martello allora dovrò aspettare nell'ordine trapano e poi cacciavite (m, , )
		obj1.use(o)
		m <- obj1

		obj2 := <- t //attendo un trapano (m,t, )
		obj2.use(o)
		t <- obj2
		
		obj3 := <- c //attendo un cacciavite (m,t,c)
		obj3.use(o)
		c <- obj3

		wg.Done() //lavoro terminato

	}
}

//non trovo un modo per scrivere un unica funzione use, esiste?
func (obj *Martello)use(o Operaio){
	fmt.Println("L'operaio", o.Nome, "ha preso il martello")
	time.Sleep(time.Second * time.Duration(rand.Intn(2)+1))
	fmt.Println("L'operaio", o.Nome, "ha terminato di usare il martello")
}

func (obj *Cacciavite)use(o Operaio){
	fmt.Println("L'operaio", o.Nome, "ha preso il cacciavite")
	time.Sleep(time.Second * time.Duration(rand.Intn(2)+1))
	fmt.Println("L'operaio", o.Nome, "ha terminato di usare il cacciavite")
}

func (obj *Trapano)use(o Operaio){
	fmt.Println("L'operaio", o.Nome, "ha preso il trapano")
	time.Sleep(time.Second * time.Duration(rand.Intn(2)+1))
	fmt.Println("L'operaio", o.Nome, "ha terminato di usare il trapano")
}

func Timer(){
	start := time.Now()
	prev := time.Duration(0)
	fmt.Println(prev)
	go func(){
		for{
			now := time.Since(start)
			if(now > prev + time.Second){
				prev = now
				fmt.Println("Sono passati", now.Truncate(time.Second))
			}
		}
	}()
}

