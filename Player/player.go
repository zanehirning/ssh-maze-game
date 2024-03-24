package player

type Point struct {
    X int
    Y int
}

type Player struct {
    Position *Point
}

func (p *Player) MoveUp() {
    p.Position.Y--
}

func (p *Player) MoveDown() {
    p.Position.Y++
}

func (p *Player) MoveLeft() {
    p.Position.X--
}

func (p *Player) MoveRight() {
    p.Position.X++
}

