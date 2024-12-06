// Student Name: Qadeer Hussain
// Lecture: Joesph Kehoe
// Aim: Implement a barrier into a rendezvous point as before it would execute part b when part a is not finished
package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

//Global variables shared between functions --A BAD IDEA

func WorkWithRendezvous(wg *sync.WaitGroup, barrier *sync.WaitGroup, Num int) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)
	//Rendezvous here

	barrier.Done() // Signal arrival at the rendezvous point (finished part a)
	barrier.Wait() // Wait for threads to rendezvous point before doing part b

	fmt.Println("PartB", Num)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	var barrier sync.WaitGroup
	//barrier := make(chan bool)
	threadCount := 5

	wg.Add(threadCount)
	barrier.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, &barrier, N)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done

}
