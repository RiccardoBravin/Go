package fn

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//	↓	Commento standard ad una funzione da esportare [Nome Descrizione]

//One prints a random number
func One() {
	fmt.Println("Random number:", rand.Intn(100)) //stampa di un numero intero
}

//Two declares and prints different kind of variables
func Two() {
	var a, b, c bool            //dichiarazione
	a, b, c = true, true, false //assegnamento
	var i int = 8               //dichiarazione e assegnamento
	s := "STRINGA"              //dichiarazione e assegnamento di variabile con gestione automatica del tipo
	const PI = 3.14
	fmt.Println("Creazione e stampa di variabili:", a, b, c, i, s, PI)
}

//Three prints and uses slices and arrays
func Three() {
	var a [2]int
	a[0] = 42
	a[1] = 69
	fmt.Println("Un array viene stampato come:", a)               //è in grado di stampare tutto un array
	fmt.Println("Un elemento dello slice è:\n", a[0], "\n", a[1]) // stampa gli elementi singolarmente

	sl := []string{"A", "caval", "donato", "non", "si", "guarda", "in", "culo"} //capienza dello slice autodeterminata dalla grandezza dei valori assegnati
	fmt.Println("Uno slice intero viene stampato come: ", sl)                   //è in grado di stampare tutto uno slice
	fmt.Println("Un elemento dello slice è:", sl[1])
	sl = sl[0:7] //sl diventa uguale a sl tra 0 e 8 escluso
	fmt.Println("Elimino l'ultimo valore", sl)
	sl = append(sl, "bocca") //con append ingrandisco lo slice (effettuo un ridimensionamento dell'array)
	fmt.Println("Aggiungo il valore corretto", sl)

	for _, i := range sl { //mettendo il _, eseguo il for finche ci sono dati ma senza il contatore
		fmt.Println(i)
	}
}

//Four takes two numbers and returns the max and min in this order
func Four(x float64, y float64) (float64, float64) {
	M := math.Max(x, y)
	m := math.Min(x, y)
	return M, m
}

//Book contains a title and an author as string
type Book struct {
	Title  string //i nomi in maiuscolo sono quelli accessibili esternamente a questo file
	Author string
}

//Five changes a book title
func Five(b *Book, title string) {
	b.Title = title
}

//Six prints things throught for loop and while
func Six() {

	fmt.Println("Stampa tramite for")
	for i := 0; i < 10; i++ { //for loop
		fmt.Print(i, " ,")
	}

	fmt.Println("\nStampa tramite while")
	num := 0
	for num < 10 {
		num = rand.Intn(10) + 1
		fmt.Print(num, ",")
	}

	fmt.Println()
}

//Seven tells you if a number is positive or negative
func Seven(x int) bool {
	if x >= 0 {
		return true
	} else if x < 0 {
		return false
	}
	return false
}

//Eight prints with defer
func Eight() {
	defer fmt.Println("Wrote before printed later") //defer posticipa l'esecuzione di un comando al momento in cui la funzione torna un valore
	fmt.Println("Wrote after printed before")
}

//Nine shows the threading
func Nine() {
	ch := make(chan float64)
	go nine(ch)
	start := time.Now()
	fmt.Print("Ricevuto il primo numero ", <-ch)
	t := time.Now()
	fmt.Println(" in", t.Sub(start))

	for i := 5.0; i < 10; { //perdo tempo facendo operazioni
		i = math.Sqrt(i * i * i)
	}

	start = time.Now()
	fmt.Print("Ricevuto il secondo numero ", <-ch)
	t = time.Now()
	fmt.Println(" in", t.Sub(start))

}

func nine(c chan float64) {
	var i float64 = 5
	for i < 10000 { //perdo tempo facendo operazioni
		i = math.Sqrt(i * i * i)
	}
	c <- i

	i = 15.0

	for i < 1000000 {
		i = math.Pow(i, 1.1)
	}
	c <- i
	close(c)

}
