package maze

import (
	//"fmt"
	"math/rand"
	"slices"
)

// Cell structure
type Cell struct {
    Row int
    Col int
    LeftNeighbor *Cell
    RightNeighbor *Cell
    TopNeighbor *Cell
    BottomNeighbor *Cell
    Visited bool
}

func (cell *Cell) hasConnection() bool {
    return cell.LeftNeighbor != nil || cell.RightNeighbor != nil || cell.TopNeighbor != nil || cell.BottomNeighbor != nil
}

func Maze(dim int) [][]*Cell {
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
                LeftNeighbor: nil,
                RightNeighbor: nil,
                TopNeighbor: nil,
                BottomNeighbor: nil,
            }
            maze[i][j] = nil
            cells[i][j] = c
        }
    }
    primsGeneration(maze, cells, frontier)
    printMaze(maze)
    return maze
}

func primsGeneration(maze [][]*Cell, cells [][]*Cell, frontier []*Cell) {
    // initialize frontier with first two cells
    frontier = append(frontier, maze[1][0])
    frontier = append(frontier, maze[0][1])

    rand := rand.New(rand.NewSource(rand.Int63()))
    for len(frontier) > 0 {
        // pick a random cell from frontier
        idx := rand.Intn(len(frontier))
        cell := frontier[idx]
        cell.Visited = true
        frontier = slices.Delete(frontier, idx, idx+1)
        // get neighbors of cell
        neighbors := getNeighbors(cell, cells)
        for _, neighbor := range neighbors {
            if neighbor.Visited {
                continue
            }
            if neighbor.hasConnection() {
                continue
            }
            frontier = append(frontier, neighbor)
            setNeighbors(cell, neighbor)
        }
    }
}

func getNeighbors(cell *Cell, cells [][]*Cell) []*Cell {
    var neighbors []*Cell
    if cell.Row != 0 {
        neighbors = append(neighbors, cells[cell.Row-1][cell.Col])
    }
    if cell.Row != len(cells)-1 {
        neighbors = append(neighbors, cells[cell.Row+1][cell.Col])
    }
    if cell.Col != 0 {
        neighbors = append(neighbors, cells[cell.Row][cell.Col-1])
    }
    if cell.Col != len(cells[0])-1 {
        neighbors = append(neighbors, cells[cell.Row][cell.Col+1])
    }
    return neighbors
}

func (c *Cell) String() string {
    s := ""
    if c.LeftNeighbor != nil {
        s += "L"
    }
    if c.RightNeighbor != nil {
        s += "R"
    }
    if c.TopNeighbor != nil {
        s += "T"
    }
    if c.BottomNeighbor != nil {
        s += "B"
    }
    s += ", "
    return s
}

func printMaze(maze [][]*Cell) {
    for i := 0; i < len(maze); i++ {
        for j := 0; j < len(maze[0]); j++ {
            cell := maze[i][j]
            print(i, j, cell.String())
        }
        println()
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
            cell.RightNeighbor = neighbor
            neighbor.LeftNeighbor = cell
        } else {
            cell.LeftNeighbor = neighbor
            neighbor.RightNeighbor = cell
        }
    } else {
        if cell.Row < neighbor.Row {
            cell.BottomNeighbor = neighbor
            neighbor.TopNeighbor = cell
        } else {
            cell.TopNeighbor = neighbor
            neighbor.BottomNeighbor = cell
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

