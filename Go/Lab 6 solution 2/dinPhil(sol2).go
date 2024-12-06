// Dining Philosophers Template Code
// Author: Joseph Kehoe
// Created: 21/10/24
//GPL Licence
// MISSING:
// 1. Readme
// 2. Full licence info.
// 3. Comments
// 4. It can Deadlock!
// Edited by : Qadeer Hussain
// Solution 2: Pick up the right fork first then the left

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func think(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5)) // Generate a random duration between 0 and 4
	time.Sleep(X * time.Second)     //wait random time amount
	fmt.Println("Phil: ", index, "was thinking")
}

func eat(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5)) // Generate a random duration between 0 and 4
	time.Sleep(X * time.Second)     //wait random time amount
	fmt.Println("Phil: ", index, "was eating")
}

func getForks(index int, forks map[int]chan bool) {
	right := (index + 1) % 5
	left := index
	forks[right] <- true // Get Right fork
	forks[left] <- true  // Get Left fork
}

func putForks(index int, forks map[int]chan bool) {
	right := (index + 1) % 5
	left := index
	<-forks[right] // Release right fork
	<-forks[left]  // Release left fork
}

func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool) {
	for {
		think(index)
		getForks(index, forks)
		eat(index)
		putForks(index, forks)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	philCount := 5
	wg.Add(philCount) // wait group for the 5 philosophers

	forks := make(map[int]chan bool)
	for k := 0; k < philCount; k++ {
		forks[k] = make(chan bool, 1) //initialize channel with buffer size 1
	} //set up forks
	for N := 0; N < philCount; N++ {
		go doPhilStuff(N, &wg, forks) // Start philosophers routine
	} //start philosophers
	wg.Wait() //wait here until everyone (10 go routines) is done
} //main
