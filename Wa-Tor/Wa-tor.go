// --------------------------------------------
// Author: Qadeer Hussain (C00270632@setu.ie)
// Created on 04/11/2024
// Modified by: Qadeer Hussain
// Due: 29/11/2024
// Lecture: Joesph Kehoe
// Wa Tor
// --------------------------------------------

package main

import (
	"fmt"
	"math/rand"
)

const (
	NumShark   = 6  // Population of Sharks
	NumFish    = 30 // Population of Fishs
	FishBreed  = 3  // Number of time units that pass before a fish can reproduce
	SharkBreed = 3  // Number of time units that must pass before a shark can reproduce
	Starve     = 4  // Period of time a shark can go without food before dying
	GridSize   = 20 // Size of World
	Threads    = 4  // Temp No of Threads
)

var Grid = make([][]int, GridSize)

// Initialize the world of FISH and SHARK :)
func initializeWorld() {
	for i := range Grid {
		Grid[i] = make([]int, GridSize)
	}

	// Randomly place fish in empty cell on the grid
	for i := 0; i < NumFish; i++ {
		x, y := rand.Intn(GridSize), rand.Intn(GridSize)
		// Find empty cell on the grid for the fish
		for Grid[x][y] != 0 {
			x, y = rand.Intn(GridSize), rand.Intn(GridSize)
		}
		Grid[x][y] = 1 // 1 represents Fish
	}

	// Randomly place Shark in empty cell on the grid
	for i := 0; i < NumShark; i++ {
		x, y := rand.Intn(GridSize), rand.Intn(GridSize)
		// Find empty cell on the grid for the shark
		for Grid[x][y] != 0 {
			x, y = rand.Intn(GridSize), rand.Intn(GridSize)
		}
		Grid[x][y] = 2 // 2 represents Shark
	}

}

// Fish Movement
func fishMovement() {

}

// Shark Movement
func sharkMovement() {

}

func printGrid() {
	for _, row := range Grid {
		for _, cell := range row {
			switch cell {
			case 0: // 0 represents an empty cell
				fmt.Print(" - ") // Print - for empty cell

			case 1: // 1 represents Fish
				fmt.Print(" F ") // Print Fish in a cell

			case 2: // 2 represents Shark
				fmt.Print(" S ") // Print Cell in a cell
			}
		}
		fmt.Println()
	}
}

func main() {
	//for {
	//	initializeWorld()
	//	printGrid()
	//	time.Sleep(2 * time.Second)
	//}
	//fishMovement()
	//sharkMovement()
	initializeWorld()
	printGrid()

}
