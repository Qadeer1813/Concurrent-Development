// --------------------------------------------
// Author: Qadeer Hussain (C00270632@setu.ie)
// Created on 04/11/2024
// Modified by: Qadeer Hussain
// Due: 06/12/2024
// Lecture: Joesph Kehoe
// Wa Tor
// --------------------------------------------
// Package main
package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	NumShark   = 30 // Population of Sharks
	NumFish    = 35 // Population of Fishs
	FishBreed  = 7  // Number of time units that pass before a fish can reproduce
	SharkBreed = 9  // Number of time units that must pass before a shark can reproduce
	Starve     = 7  // Period of time a shark can go without food before dying
	GridSize   = 30 // Size of World
	Threads    = 1  // Temp No of Threads
)

const (
	EmptyCell = 0
	Fish      = 1
	Shark     = 2
)

const (
	North = 0
	East  = 1
	South = 2
	West  = 3
)

var Grid = make([][]int, GridSize)
var simulationRunning = false
var fishChronons = make([][]int, GridSize)  // Tracks the number of chronons for each fish
var sharkChronons = make([][]int, GridSize) // Tracks the number of chronons for each shark's breeding cycle
var sharkStarve = make([][]int, GridSize)   // Tracks the number of chronons since last eaten fish

// Initialize an empty grid
func EmptyGrid() {
	for i := range Grid {
		Grid[i] = make([]int, GridSize)
	}
}

// Check if the system supports the required number of threads
func CheckThreads() {
	availableThreads := runtime.NumCPU()
	if availableThreads < Threads {
		fmt.Printf("Warning: The system supports only %d threads, but %d are requested.\n", availableThreads, Threads)
	} else {
		fmt.Printf("System supports %d threads. Simulation will proceed.\n", availableThreads)
	}
}

// Initialize the world of FISH and SHARK :)
func InitializeWorld() {
	for i := range Grid {
		for j := range Grid[i] {
			Grid[i][j] = EmptyCell
		}
	}
	InitializeFishChronons() // Reset fish reproduction chronons
	InitializeSharksStat()   // Reset Shark reproduction chronons

	// Randomly place fish on the grid
	for i := 0; i < NumFish; i++ {
		x, y := rand.Intn(GridSize), rand.Intn(GridSize)
		// Find empty cell on the grid for the fish
		for Grid[x][y] != EmptyCell {
			x, y = rand.Intn(GridSize), rand.Intn(GridSize)
		}
		Grid[x][y] = Fish
	}

	// Randomly place Shark on the grid
	for i := 0; i < NumShark; i++ {
		x, y := rand.Intn(GridSize), rand.Intn(GridSize)
		// Find empty cell on the grid for the shark
		for Grid[x][y] != EmptyCell {
			x, y = rand.Intn(GridSize), rand.Intn(GridSize)
		}
		Grid[x][y] = Shark
	}

}

// Initialize fishChronons grid
func InitializeFishChronons() {
	for i := range fishChronons {
		fishChronons[i] = make([]int, GridSize)
	}
}

// Initialize shark stats(chronons and starvation) grid
func InitializeSharksStat() {
	for i := range sharkChronons {
		sharkChronons[i] = make([]int, GridSize)
		sharkStarve[i] = make([]int, GridSize)
	}
}

// Generate random list of directions for the sharks and fish to move
func RandomDirection() []int {
	directions := []int{North, East, South, West}
	rand.Shuffle(len(directions), func(i, j int) {
		directions[i], directions[j] = directions[j], directions[i]
	})
	return directions
}

// Execute the fish movement
func ExecuteFishMove(fromI, fromJ, toI, toJ int, newGrid *[][]int) {
	// Move fish and reproduce if can
	if fishChronons[fromI][fromJ] >= FishBreed {
		(*newGrid)[fromI][fromJ] = Fish // Fish reproduces
		fishChronons[fromI][fromJ] = 0  // Reset breeding time
	} else {
		(*newGrid)[fromI][fromJ] = EmptyCell // Fish moves without reproducing
	}
	(*newGrid)[toI][toJ] = Fish
	fishChronons[toI][toJ] = fishChronons[fromI][fromJ] // Increment breeding chronon
}

// Execute the shark movement
func ExecuteSharkMove(fromI, fromJ, toI, toJ int, newGrid *[][]int, eatenFish bool) {
	// Move shark and reproduce if can
	if sharkChronons[fromI][fromJ] >= SharkBreed {
		(*newGrid)[fromI][fromJ] = Shark // Shark reproduces
		sharkChronons[fromI][fromJ] = 0  // Reset breeding timer
	} else {
		(*newGrid)[fromI][fromJ] = EmptyCell // Shark moves without reproducing
	}

	(*newGrid)[toI][toJ] = Shark
	sharkChronons[toI][toJ] = sharkChronons[fromI][fromJ] + 1 // Increment breeding chronon

	if eatenFish {
		sharkStarve[toI][toJ] = 0 // Reset hunger if shark has eaten
	} else {
		sharkStarve[toI][toJ] = sharkStarve[fromI][fromJ] + 1 // Increment hunger
	}

	// Check if the shark starves
	if sharkStarve[toI][toJ] >= Starve {
		(*newGrid)[toI][toJ] = EmptyCell // Shark dies of starvation
	}
}

// Concurrent processing of the grid
func Concurrent(grid [][]int, processFunc func(int, int, *[][]int), newGrid *[][]int) {
	var wg sync.WaitGroup // Wait group for goroutines
	var mutex sync.Mutex  // Mutex prevent race cond

	// Calculate sub-grid size each thread will be processing
	subGrid := GridSize / Threads

	for threadID := 0; threadID < Threads; threadID++ {
		wg.Add(1) // Waitgroup counter for each thread
		go func(tID int) {
			defer wg.Done() // Decrement the WaitGroup counter when this goroutine finishes

			// Calculate start and end rows for this thread
			startRow := tID * subGrid
			endRow := startRow + subGrid
			// Last thread covers any remaining rows
			if tID == Threads-1 {
				endRow = GridSize
			}

			// Process rows for this thread
			for i := startRow; i < endRow; i++ {
				for j := 0; j < GridSize; j++ {
					if grid[i][j] != EmptyCell { // Only process non-empty cells
						mutex.Lock()               // Lock the shared grid
						processFunc(i, j, newGrid) // Apply the processing function
						mutex.Unlock()             // Unlock shared grid
					}
				}
			}
		}(threadID) // Pass the thread ID to the goroutine
	}

	// Wait for all go routines to complete
	wg.Wait()
}

// Fish Movement and Reproduction
func FishMovement() {
	// Create a temporary grid to store new positions
	newGrid := make([][]int, GridSize)
	for i := range newGrid {
		newGrid[i] = make([]int, GridSize)
	}

	// Copy fish and shark positions to the temporary grid
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			newGrid[i][j] = Grid[i][j]
		}
	}

	//// Attempt to move or reproduce fish
	//for i := 0; i < GridSize; i++ {
	//	for j := 0; j < GridSize; j++ {
	// Use the Concurrent function to handle fish movement and reproduction
	Concurrent(Grid, func(i, j int, newGrid *[][]int) {
		if Grid[i][j] == Fish {
			fishChronons[i][j]++
			// Get a random list of movement directions
			directions := RandomDirection()
			var moved bool
			for _, direction := range directions {
				if moved {
					break
				}
				// Move the fish in the given direction
				switch direction {
				case North: // North
					if i > 0 && (*newGrid)[i-1][j] == EmptyCell {
						ExecuteFishMove(i, j, i-1, j, newGrid)
						moved = true
					}
				case East: // East
					if j < GridSize-1 && (*newGrid)[i][j+1] == EmptyCell {
						ExecuteFishMove(i, j, i, j+1, newGrid)
						moved = true
					}
				case South: // South
					if i < GridSize-1 && (*newGrid)[i+1][j] == EmptyCell {
						ExecuteFishMove(i, j, i+1, j, newGrid)
						moved = true
					}
				case West: // West
					if j > 0 && (*newGrid)[i][j-1] == EmptyCell {
						ExecuteFishMove(i, j, i, j-1, newGrid)
						moved = true
					}
				}
			}

			// If no move was made, keep fish in the current cell
			if !moved {
				(*newGrid)[i][j] = Fish
			}
		}
	}, &newGrid)
	// Update the grid with new positions
	Grid = newGrid
}

// Shark Movement
func SharkMovement() {
	// Create a temporary grid to store new positions
	newGrid := make([][]int, GridSize)
	for i := range newGrid {
		newGrid[i] = make([]int, GridSize)
	}

	// Copy fish and shark positions to the temporary grid
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			newGrid[i][j] = Grid[i][j]
		}
	}

	// Attempt to move or reproduce the sharks
	//for i := 0; i < GridSize; i++ {
	//	for j := 0; j < GridSize; j++ {
	// Use the Concurrent function to handle shark movement and reproduction
	Concurrent(Grid, func(i, j int, newGrid *[][]int) {
		if Grid[i][j] == Shark {
			// Get a random list of movement directions
			directions := RandomDirection()
			var moved bool
			// First try to find and eat fish
			for _, direction := range directions {
				if moved {
					break
				}
				switch direction {
				case North:
					if i > 0 && (*newGrid)[i-1][j] == Fish {
						ExecuteSharkMove(i, j, i-1, j, newGrid, true)
						moved = true
					}
				case East:
					if j < GridSize-1 && (*newGrid)[i][j+1] == Fish {
						ExecuteSharkMove(i, j, i, j+1, newGrid, true)
						moved = true
					}
				case South:
					if i < GridSize-1 && (*newGrid)[i+1][j] == Fish {
						ExecuteSharkMove(i, j, i+1, j, newGrid, true)
						moved = true
					}
				case West:
					if j > 0 && (*newGrid)[i][j-1] == Fish {
						ExecuteSharkMove(i, j, i, j-1, newGrid, true)
						moved = true
					}
				}
			}
			// If no fish was eaten, try to move to an empty cell
			if !moved {
				for _, direction := range directions {
					if moved {
						break
					}
					switch direction {
					case North:
						if i > 0 && (*newGrid)[i-1][j] == EmptyCell {
							ExecuteSharkMove(i, j, i-1, j, newGrid, false)
							moved = true
						}
					case East:
						if j < GridSize-1 && (*newGrid)[i][j+1] == EmptyCell {
							ExecuteSharkMove(i, j, i, j+1, newGrid, false)
							moved = true
						}
					case South:
						if i < GridSize-1 && (*newGrid)[i+1][j] == EmptyCell {
							ExecuteSharkMove(i, j, i+1, j, newGrid, false)
							moved = true
						}
					case West:
						if j > 0 && (*newGrid)[i][j-1] == EmptyCell {
							ExecuteSharkMove(i, j, i, j-1, newGrid, false)
							moved = true
						}
					}
				}
			}
			// If no move was made, increment sharks chronons
			if !moved {
				sharkChronons[i][j]++ // Increment breeding chronon
				sharkStarve[i][j]++   // Increment hunger chronon
				if sharkStarve[i][j] >= Starve {
					(*newGrid)[i][j] = EmptyCell // Shark dies of starvation
				}
			}
		}
	}, &newGrid)
	// Update the grid with new positions
	Grid = newGrid
}

var w fyne.Window

// Create grid displayed using Fyne Library
func CreateGrid() *fyne.Container {
	grid := container.NewGridWithColumns(GridSize)
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			var cellColor color.Color
			switch Grid[i][j] {
			case Fish: // Fish
				cellColor = color.RGBA{R: 0, G: 255, B: 0, A: 255}
			case Shark: // Shark
				cellColor = color.RGBA{R: 255, G: 0, B: 0, A: 255}
			default: // Empty
				cellColor = color.Gray{Y: 180}
			}
			rect := canvas.NewRectangle(cellColor)
			rect.SetMinSize(fyne.NewSize(10, 10))
			grid.Add(rect)
		}
	}
	return grid
}

const MaxIterations = 250 // Maximum number of iterations

func UpdateFunc() {
	iteration := 0          // Counter to track iterations
	startTime := time.Now() // Record the start time of the simulation

	for simulationRunning && iteration < MaxIterations {
		FishMovement()  // Update fish positions
		SharkMovement() // Update shark movement
		iteration++     // Increment the iteration counter

		w.SetContent(container.NewVBox( // Refresh UI with updated grid and button
			CreateGrid(),
			widget.NewButton("Stop", func() {
				simulationRunning = false // Stop the simulation
			}),
		))
	}

	// Stop the simulation after reaching the maximum iterations
	if iteration >= MaxIterations {
		simulationRunning = false
		elapsedTime := time.Since(startTime) // Calculate elapsed time
		fmt.Printf("Simulation completed in %s after %d iterations.\n", elapsedTime, MaxIterations)
	}
}

func main() {
	a := app.New()
	w = a.NewWindow("Wa-Tor Simulation")

	CheckThreads() // Function check if user machine has enough threads

	EmptyGrid() // Start with empty grid

	// Add the Start button to the UI
	startButton := widget.NewButton("Start", func() {
		if !simulationRunning {
			simulationRunning = true
			InitializeWorld()
			go UpdateFunc()
		}
	})

	w.SetContent(container.NewVBox(
		CreateGrid(),
		startButton,
	))
	w.ShowAndRun()
}
