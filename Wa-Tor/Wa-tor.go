// --------------------------------------------
// Author: Qadeer Hussain (C00270632@setu.ie)
// Created on 04/11/2024
// Modified by: Qadeer Hussain
// Due: 06/12/2024
// Lecture: Joesph Kehoe
// Wa Tor
// --------------------------------------------

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math/rand"
	"time"
)

const (
	NumShark   = 10 // Population of Sharks
	NumFish    = 10 // Population of Fishs
	FishBreed  = 5  // Number of time units that pass before a fish can reproduce
	SharkBreed = 7  // Number of time units that must pass before a shark can reproduce
	Starve     = 5  // Period of time a shark can go without food before dying
	GridSize   = 10 // Size of World
	Threads    = 4  // Temp No of Threads
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

// Initialize the world of FISH and SHARK :)
func initializeWorld() {
	for i := range Grid {
		for j := range Grid[i] {
			Grid[i][j] = EmptyCell
		}
	}
	initializeFishChronons() // Reset fish reproduction chronons
	initializeSharksStat()   // Reset Shark reproduction chronons

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
func initializeFishChronons() {
	for i := range fishChronons {
		fishChronons[i] = make([]int, GridSize)
	}
}

// Initialize shark stats(chronons and starvation) grid
func initializeSharksStat() {
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
func executeFishMove(fromI, fromJ, toI, toJ int, newGrid *[][]int) {
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
func executeSharkMove(fromI, fromJ, toI, toJ int, newGrid *[][]int, eatenFish bool) {
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

// Fish Movement and Reproduction
func fishMovement() {
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

	// Attempt to move or reproduce fish
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			if Grid[i][j] == Fish {
				fishChronons[i][j]++
				directions := RandomDirection()
				var moved bool
				for _, direction := range directions {
					if moved {
						break
					}
					switch direction {
					case North: // North
						if i > 0 && newGrid[i-1][j] == EmptyCell {
							executeFishMove(i, j, i-1, j, &newGrid)
							moved = true
						}
					case East: // East
						if j < GridSize-1 && newGrid[i][j+1] == EmptyCell {
							executeFishMove(i, j, i, j+1, &newGrid)
							moved = true
						}
					case South: // South
						if i < GridSize-1 && newGrid[i+1][j] == EmptyCell {
							executeFishMove(i, j, i+1, j, &newGrid)
							moved = true
						}
					case West: // West
						if j > 0 && newGrid[i][j-1] == EmptyCell {
							executeFishMove(i, j, i, j-1, &newGrid)
							moved = true
						}
					}
				}

				// If no move was made, keep fish in the current cell
				if !moved {
					newGrid[i][j] = Fish
				}
			}
		}
	}
	// Update the grid with new positions
	Grid = newGrid
}

// Shark Movement
func sharkMovement() {
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
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			if Grid[i][j] == Shark {
				directions := RandomDirection()
				var moved bool
				// First try to find and eat fish
				for _, direction := range directions {
					if moved {
						break
					}
					switch direction {
					case North:
						if i > 0 && newGrid[i-1][j] == Fish {
							executeSharkMove(i, j, i-1, j, &newGrid, true)
							moved = true
						}
					case East:
						if j < GridSize-1 && newGrid[i][j+1] == Fish {
							executeSharkMove(i, j, i, j+1, &newGrid, true)
							moved = true
						}
					case South:
						if i < GridSize-1 && newGrid[i+1][j] == Fish {
							executeSharkMove(i, j, i+1, j, &newGrid, true)
							moved = true
						}
					case West:
						if j > 0 && newGrid[i][j-1] == Fish {
							executeSharkMove(i, j, i, j-1, &newGrid, true)
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
							if i > 0 && newGrid[i-1][j] == EmptyCell {
								executeSharkMove(i, j, i-1, j, &newGrid, false)
								moved = true
							}
						case East:
							if j < GridSize-1 && newGrid[i][j+1] == EmptyCell {
								executeSharkMove(i, j, i, j+1, &newGrid, false)
								moved = true
							}
						case South:
							if i < GridSize-1 && newGrid[i+1][j] == EmptyCell {
								executeSharkMove(i, j, i+1, j, &newGrid, false)
								moved = true
							}
						case West:
							if j > 0 && newGrid[i][j-1] == EmptyCell {
								executeSharkMove(i, j, i, j-1, &newGrid, false)
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
						newGrid[i][j] = EmptyCell // Shark dies of starvation
					}
				}
			}
		}
	}
	// Update the grid with new positions
	Grid = newGrid
}

var w fyne.Window

// Create grid displayed using Fyne Library
func createGrid() *fyne.Container {
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
			rect.SetMinSize(fyne.NewSize(20, 20))
			grid.Add(rect)
		}
	}
	return grid
}

func updateFunc() {
	if !simulationRunning {
		return
	}
	fishMovement()                  // Update fish positions
	sharkMovement()                 // Update shark movement
	w.SetContent(container.NewVBox( // Refresh UI with Start button
		createGrid(),
		widget.NewButton("Start", func() {
			if !simulationRunning {
				simulationRunning = true
				initializeWorld()
				updateFunc()
			}
		}),
	))
	time.AfterFunc(time.Second, updateFunc) // Schedule the next update
}

func main() {
	a := app.New()
	w = a.NewWindow("Wa-Tor Simulation")

	EmptyGrid() // Start with empty grid

	// Add the Start button to the UI
	startButton := widget.NewButton("Start", func() {
		if !simulationRunning {
			simulationRunning = true
			initializeWorld()
			updateFunc()
		}
	})

	w.SetContent(container.NewVBox(
		createGrid(),
		startButton,
	))
	w.ShowAndRun()
}
