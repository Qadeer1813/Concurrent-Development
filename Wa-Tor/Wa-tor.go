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
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"math/rand"
	"time"
)

const (
	NumShark   = 6  // Population of Sharks
	NumFish    = 30 // Population of Fishs
	FishBreed  = 3  // Number of time units that pass before a fish can reproduce
	SharkBreed = 3  // Number of time units that must pass before a shark can reproduce
	Starve     = 4  // Period of time a shark can go without food before dying
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

// Initialize the world of FISH and SHARK :)
func initializeWorld() {
	for i := range Grid {
		Grid[i] = make([]int, GridSize)
	}

	// Randomly place fish in empty cell on the grid
	for i := 0; i < NumFish; i++ {
		x, y := rand.Intn(GridSize), rand.Intn(GridSize)
		// Find empty cell on the grid for the fish
		for Grid[x][y] != EmptyCell {
			x, y = rand.Intn(GridSize), rand.Intn(GridSize)
		}
		Grid[x][y] = Fish
	}

	// Randomly place Shark in empty cell on the grid
	for i := 0; i < NumShark; i++ {
		x, y := rand.Intn(GridSize), rand.Intn(GridSize)
		// Find empty cell on the grid for the shark
		for Grid[x][y] != EmptyCell {
			x, y = rand.Intn(GridSize), rand.Intn(GridSize)
		}
		Grid[x][y] = Shark
	}

}

// Fish Movement
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

	// Move each fish to the first available empty cell
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			if Grid[i][j] == Fish { // Check if this cell contains a fish
				var moved bool
				// Check each direction for movement
				for direction := 0; direction < 4 && !moved; direction++ {
					switch direction {
					case North: // North
						if i > 0 && newGrid[i-1][j] == EmptyCell {
							newGrid[i-1][j] = Fish
							newGrid[i][j] = EmptyCell
							moved = true
						}
					case East: // East
						if j < GridSize-1 && newGrid[i][j+1] == EmptyCell {
							newGrid[i][j+1] = Fish
							newGrid[i][j] = EmptyCell
							moved = true
						}
					case South: // South
						if i < GridSize-1 && newGrid[i+1][j] == EmptyCell {
							newGrid[i+1][j] = Fish
							newGrid[i][j] = EmptyCell
							moved = true
						}
					case West: // West
						if j > 0 && newGrid[i][j-1] == EmptyCell {
							newGrid[i][j-1] = Fish
							newGrid[i][j] = EmptyCell
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

	// Move each shark to the first available empty cell
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			if Grid[i][j] == Shark { // Check if this cell contains a shark
				var moved bool
				// Check each direction for movement
				for direction := 0; direction < 4 && !moved; direction++ {
					switch direction {
					case North: // North
						if i > 0 && newGrid[i-1][j] == EmptyCell {
							newGrid[i-1][j] = Shark
							newGrid[i][j] = EmptyCell
							moved = true
						}
					case East: // East
						if j < GridSize-1 && newGrid[i][j+1] == EmptyCell {
							newGrid[i][j+1] = Shark
							newGrid[i][j] = EmptyCell
							moved = true
						}
					case South: // South
						if i < GridSize-1 && newGrid[i+1][j] == EmptyCell {
							newGrid[i+1][j] = Shark
							newGrid[i][j] = EmptyCell
							moved = true
						}
					case West: // West
						if j > 0 && newGrid[i][j-1] == EmptyCell {
							newGrid[i][j-1] = Shark
							newGrid[i][j] = EmptyCell
							moved = true
						}
					}
				}

				// If no move was made, keep shark in the current cell
				if !moved {
					newGrid[i][j] = Shark
				}
			}
		}
	}
	// Update the grid with new positions
	Grid = newGrid
}

var w fyne.Window

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
	fishMovement()                          // Update fish positions
	sharkMovement()                         // Update shark movement
	w.SetContent(createGrid())              // Refresh UI
	time.AfterFunc(time.Second, updateFunc) // Schedule the next update
}

func main() {
	a := app.New()
	w = a.NewWindow("Wa-Tor Simulation")

	initializeWorld()          // Ensure this function populates the initial state of Grid
	w.SetContent(createGrid()) // Display the initial grid

	updateFunc() // Start the update loop
	w.ShowAndRun()
}
