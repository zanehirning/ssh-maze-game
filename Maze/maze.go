package main

import (
    "math/rand"
    "fmt"
    "slices"
)

// Cell structure
type Cell struct {
    Row int
    Col int
    leftNeighbor *Cell
    rightNeighbor *Cell
    topNeighbor *Cell
    bottomNeighbor *Cell
    Visited bool
}

func (cell *Cell) hasConnection() bool {
    return cell.leftNeighbor != nil || cell.rightNeighbor != nil || cell.topNeighbor != nil || cell.bottomNeighbor != nil
}

func maze(dim int) {
    var maze [][]*Cell
    var frontier []*Cell
    var cells [][]*Cell

    for i := 0; i < dim; i++ {
        row := make([]*Cell, dim)
        maze = append(maze, row)
        cells = append(cells, row)
        for j := 0; j < dim; j++ {
            c := &Cell{
                Row: i,
                Col: j,
                Visited: false,
                leftNeighbor: nil,
                rightNeighbor: nil,
                topNeighbor: nil,
                bottomNeighbor: nil,
            }
            maze[i][j] = c
            cells[i][j] = c
        }
    }
    primsGeneration(maze, cells, frontier)
}

func primsGeneration(maze [][]*Cell, cells [][]*Cell, frontier []*Cell) {
    // initialize frontier with first two cells
    frontier = append(frontier, maze[1][0])
    frontier = append(frontier, maze[0][1])

    rand := rand.New(rand.NewSource(rand.Int63()))
    
    maze[0][0] = cells[0][0]
    for len(frontier) > 0 {
        // pick random cell from frontier
        randomIndex := rand.Intn(len(frontier))
        randomFrontierCell := frontier[randomIndex]
        //remove from frontier
        frontier = append(frontier[:randomIndex], frontier[randomIndex+1:]...)
        maze[randomFrontierCell.Row][randomFrontierCell.Col] = randomFrontierCell

        // mark cell as visited
        randomFrontierCell.Visited = true

        // get neighbors
        neighbors := getNeighborsInMaze(randomFrontierCell, maze)
        chosenNeighbor := neighbors[rand.Intn(len(neighbors))]
        setNeighbors(randomFrontierCell, chosenNeighbor)

        cellsInMaze := getCellsInMaze(maze)
        for _, neighbor := range neighbors {
            if (!slices.Contains(frontier, neighbor) && !slices.Contains(cellsInMaze, neighbor)) {
                frontier = append(frontier, neighbor)
            }
        }
    }
}

func getCellsInMaze(maze [][]*Cell) []*Cell {
    var cells []*Cell
    for i := 0; i < len(maze); i++ {
        for j := 0; j < len(maze[0]); j++ {
            cell := maze[i][j]
            if cell.hasConnection() {
                cells = append(cells, cell)
            }
        }
    }
    return cells
}

func setNeighbors(cell *Cell, neighbor *Cell) {
    if cell.Row == neighbor.Row {
        if cell.Col < neighbor.Col {
            cell.rightNeighbor = neighbor
            neighbor.leftNeighbor = cell
        } else {
            cell.leftNeighbor = neighbor
            neighbor.rightNeighbor = cell
        }
    } else {
        if cell.Row < neighbor.Row {
            cell.bottomNeighbor = neighbor
            neighbor.topNeighbor = cell
        } else {
            cell.topNeighbor = neighbor
            neighbor.bottomNeighbor = cell
        }
    }
}

func getNeighborsInMaze(cell *Cell, maze [][]*Cell) []*Cell {
    var neighbors []*Cell
    if cell.Row > 0 {
        neighbors = append(neighbors, maze[cell.Row-1][cell.Col])
    }
    if cell.Row < len(maze)-1 {
        neighbors = append(neighbors, maze[cell.Row+1][cell.Col])
    }
    if cell.Col > 0 {
        neighbors = append(neighbors, maze[cell.Row][cell.Col-1])
    }
    if cell.Col < len(maze[0])-1 {
        neighbors = append(neighbors, maze[cell.Row][cell.Col+1])
    }
    return neighbors
}

func main() {
    maze(10)
}
