package main

import (
	"math/rand"
)

const (
	NB_ROW = 30
	NB_COL = 60
)

// Une cellule visité ou pas avec des murs aux 4 pôles
type Cell struct {
	visited bool
	nWall   bool
	sWall   bool
	oWall   bool
	eWall   bool
	row     int
	col     int
}

func (c Cell) allWalls() bool {
	return c.nWall && c.sWall && c.oWall && c.eWall
}

var visitedCells int
var totalCells int
var currentCell Cell
var cellStacks []Cell

func genGrid() (grid [NB_ROW][NB_COL]Cell) {
	visitedCells = 0
	totalCells = 0
	for r := 1; r < NB_ROW-1; r++ {
		for c := 1; c < NB_COL-1; c++ {
			cell := Cell{false, true, true, true, true, r, c}
			grid[r][c] = cell
		}
	}

	totalCells = (NB_COL - 2) * (NB_ROW - 2)

	currentCell = grid[1][1]
	currentCell.visited = true
	visitedCells += 1

	return grid
}

func findNeighbors(c Cell, grid [NB_ROW][NB_COL]Cell) (neighbors []Cell) {
	if c.row-1 >= 0 && c.row+1 < NB_ROW && c.col-1 >= 0 && c.col+1 < NB_COL {
		neigh := []Cell{grid[c.row-1][c.col],
			grid[c.row][c.col-1], grid[c.row][c.col+1],
			grid[c.row+1][c.col]}
		for _, c := range neigh {
			if c.allWalls() {
				neighbors = append(neighbors, c)
			}
		}
		return neighbors
	}
	return
}

func next(grid [NB_ROW][NB_COL]Cell) (newGrid [NB_ROW][NB_COL]Cell) {
	neighbors := findNeighbors(currentCell, grid)
	// Choisis un voisin aléatoire ou pas
	var neighbor Cell
	if len(neighbors) >= 1 {
		rand := randInt(0, len(neighbors))
		neighbor = neighbors[rand]
		// Brise les murs des cellules entre la courante et le voisin
		if neighbor.row == currentCell.row+1 {
			//println("Au dessous")
			grid[currentCell.row][currentCell.col].sWall = false
			grid[neighbor.row][neighbor.col].nWall = false
		} else if neighbor.row == currentCell.row-1 {
			//println("Au dessus")
			grid[currentCell.row][currentCell.col].nWall = false
			grid[neighbor.row][neighbor.col].sWall = false
		} else if neighbor.col == currentCell.col+1 {
			//println("A droite")
			grid[currentCell.row][currentCell.col].eWall = false
			grid[neighbor.row][neighbor.col].oWall = false
		} else if neighbor.col == currentCell.col-1 {
			//println("Au gauche")
			grid[currentCell.row][currentCell.col].oWall = false
			grid[neighbor.row][neighbor.col].eWall = false
		}

		// On ajoute la cellule courrante dans le stock
		cellStacks = append(cellStacks, currentCell)
		currentCell = neighbor
		visitedCells += 1
	} else {
		// Pop d'une liste
		neighbor, cellStacks = cellStacks[len(cellStacks)-1], cellStacks[:len(cellStacks)-1]
		currentCell = neighbor
	}
	newGrid = grid
	return newGrid
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
