// Stefan Nilsson 2013-03-13

/*
*Vad händer om man byter plats på satserna wg.Wait() och close(ch) i slutet av main-funktionen?
-ch kommer att stängas innan producernterna har hunnit producera alla strängar. Detta resulterar i att Producer kommer att
-försöka skicka data till en stängd channel.
*Vad händer om man flyttar close(ch) från main-funktionen och i stället stänger kanalen i slutet av funktionen Produce?
-en producent kommer hinna producera klart alla sina strängar och sedan stänga ch. När då de andra producernterna ska
-skicka sina strängar kommer de försöka skicka till en stängd kanal.
-Självklart kommer en hel del strängar hinna bli konsumerade innan detta sker.
*Vad händer om man tar bort satsen close(ch) helt och hållet?
-ingenting eftersom programmet själv kommer stänga kanalen då det inte finns fler värden att hämta i den
*Vad händer om man ökar antalet konsumenter från 2 till 4?
-en konsument kommer nu att konsumera 32/4 strängar istället för 32/2 strängar. Programmet fungerar som förväntat.
*Kan man vara säker på att alla strängar blir utskrivna innan programmet stannar?
-nej, trådarna synkroniserar fram till att sista producertråden lyckats skicka på kanalen till en konsument. Det kan hända
-att det direkt blir producentens tur att köra så att wg.Done() anropas så att programmet till slut terminerar. Då kan det
-hända att konsumenten inte hinner konsumera ordentligt. 
*/


// This is a testbed to help you understand channels better.
package many2many

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func Main() {
	// Use different random numbers each time this program is executed.
	rand.Seed(time.Now().Unix())

	const strings = 32
	const producers = 4
	const consumers = 2

	before := time.Now()
	ch := make(chan string)
	wgp := new(sync.WaitGroup)
	wgp.Add(producers)
	wgp2 := new(sync.WaitGroup)
	wgp2.Add(consumers)
	for i := 0; i < producers; i++ {
		go Produce("p" + strconv.Itoa(i), strings/producers, ch, wgp)
	}
	for i := 0; i < consumers; i++ {
		go Consume("c" + strconv.Itoa(i), ch, wgp2)
	}
	wgp.Wait() // Wait for all producers to finish.
	close(ch)
	wgp2.Wait() // Wait for all consumers to finish.
	fmt.Println("time:", time.Now().Sub(before))
}

// Produce sends n different strings on the channel and notifies wg when done.
func Produce(id string, n int, ch chan <- string, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		RandomSleep(100) // Simulate time to produce data.
		ch <- id + ":" + strconv.Itoa(i)
		//fmt.Println("p")
	}
	wg.Done()
}

// Consume prints strings received from the channel until the channel is closed.
func Consume(id string, ch <-chan string, wg *sync.WaitGroup) {
	for s := range ch {
		fmt.Println(id, "received", s)
		RandomSleep(100) // Simulate time to consume data.
	}
	wg.Done()
}

// RandomSleep waits for x ms, where x is a random number
func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n))*time.Millisecond)
}
