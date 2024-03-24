package bubbletea

import (
    "fmt"
    "os"
    "maze/Maze"
    "maze/Player"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    maze [][]*maze.Cell
    player *player.Player
    mazeString [][]string
}

func initialModel() model {
    maze := maze.Maze(10)
    return model{
        maze: maze,
        player: &player.Player{Position: &player.Point{X: 0, Y: 0}},
        mazeString: make([][]string, 10),
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q":
            return m, tea.Quit
        case "j", "down":
            if (m.player.Position.Y < len(m.maze) - 1 && m.maze[m.player.Position.X][m.player.Position.Y].BottomNeighbor != nil) {
                m.player.MoveDown()
            }  
        case "k", "up":
            if (m.player.Position.Y < 0 && m.maze[m.player.Position.X][m.player.Position.Y].TopNeighbor != nil) {
                m.player.MoveUp()
            }
        case "h", "left":
            if (m.player.Position.X > 0 && m.maze[m.player.Position.X][m.player.Position.X].LeftNeighbor != nil) {
                m.player.MoveLeft()
            }  
        case "l", "right":
            if (m.player.Position.X < len(m.maze) - 1 && m.maze[m.player.Position.X][m.player.Position.Y].RightNeighbor != nil) {
                m.player.MoveRight()
            }
        case "enter":
            m.maze = maze.Maze(10);
            m.player.Position.X = 0
            m.player.Position.Y = 0
        }
    }

    return m, nil
}

func (m model) View() string {
    s := "Press 'q' to quit\n Press Enter to generate a new maze\n\n"

    for i, row := range m.maze {
        for j := range row {
            if m.player.Position.X == j && m.player.Position.Y == i {
                s += "P"
            } else {
                s += " "
            }
        }
        s += "\n"
    }

    return s
}

func Bubbletea() {
    p := tea.NewProgram(initialModel())
    if err := p.Start(); err != nil {
        fmt.Fprintf(os.Stderr, "Error starting program: %v", err)
        os.Exit(1)
    }
}
