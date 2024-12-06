//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Kelan Morgan, Qadeer Hussain
// Description:
// A simple barrier implemented using mutex and unbuffered channel
// Issues:
// None I hope
//1. Change mutex to atomic variable
//2. Make it a reusable barrier
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, arrived *atomic.Int64, max int, wg *sync.WaitGroup, theChan chan bool) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	//we wait here until everyone has completed part A
	//Kelan Morgan
	arrived.Add(1)
	if arrived.Load() == int64(max) { //last to arrive -signal others to go
		for range max - 1 { // run for all the other routines to free them up
			theChan <- true
		}
	} else { //not all here yet we wait until signal
		<-theChan
	} //end of if-else
	fmt.Println("Part B", goNum)
	// wait here until everyone has completed part B
	//Qadeer Hussain
	arrived.Add(-1)
	if arrived.Load() == 0 { // last routine arrives here
		for range max - 1 { // run for all the other routines to free them up
			theChan <- true
		}
	} else {
		<-theChan // wait here for last routine
	}
	fmt.Println("Part C", goNum)
	wg.Done()
	return true
} //end-doStuff

// Qadeer and Kelan
func main() {
	totalRoutines := 100
	var arrived atomic.Int64
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these
	theChan := make(chan bool)     //use unbuffered channel in place of semaphore
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &arrived, totalRoutines, &wg, theChan)
	}
	wg.Wait() //wait for everyone to finish before exiting
} //end-main
