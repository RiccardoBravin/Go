package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Gruppo struct {
    nome string
    nPalline int
}

type Tunnel struct {
    libero bool
}

func transumanza(g Gruppo, t chan Tunnel, c1 chan int, wg *sync.WaitGroup){
    for g.nPalline > 0{
        time.Sleep(time.Duration(rand.Intn(2))*time.Second)
        mandaPersona(&g, t, c1, wg)
    }
}

func mandaPersona(g *Gruppo, t chan Tunnel, c1 chan int, wg *sync.WaitGroup){
    fmt.Println("manda persona",g.nome)
    tunnel := <- t
    if tunnel.libero {
        tunnel.libero = false
        t <- tunnel
        select{
		    case <- c1:
			    fmt.Println("Scontro ", g.nome)
                wg.Done()
                // GO si arrabbia se non usate una varabile...
		    case <-time.After(time.Second):
                tunnel := <- t
                tunnel.libero = true
                t <- tunnel
		        fmt.Println("La pallina in arrivo da", g.nome, "è passata")
                g.nPalline = g.nPalline - 1
                fmt.Println("Rimangono ", g.nPalline, " nel gruppo ", g.nome)
	    }
        
    } else{

        
        tunnel.libero = true
        t <- tunnel
        c1 <- 1
        wg.Wait()
    }
    
}

func main() {
    rand.Seed(time.Now().UnixNano())
    gruppo1 := Gruppo{"destra",5}
    gruppo2 := Gruppo{"sinistra",5}
    
    c1 := make(chan int,1)

    var wg sync.WaitGroup
    wg.Add(0)
    //wg.Done()
    
    tunnelChannel := make(chan Tunnel, 1)
    tunnel := Tunnel{true}
    tunnelChannel <- tunnel
    
    go transumanza(gruppo1,tunnelChannel,c1, &wg)
    go transumanza(gruppo2,tunnelChannel,c1, &wg)
    
	time.Sleep(time.Minute)
}
	
	
	
	
	
	
	
	
	
	
	