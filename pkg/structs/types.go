package structs

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Left goes left
func (c Coordinate) Left() Coordinate {
	result := Coordinate{
		X: c.X - 1,
		Y: c.Y,
	}
	return result
}

// Left goes left
func (c Coordinate) Right() Coordinate {
	result := Coordinate{
		X: c.X + 1,
		Y: c.Y,
	}
	return result
}

// Left goes left
func (c Coordinate) Up() Coordinate {
	result := Coordinate{
		X: c.X,
		Y: c.Y - 1,
	}
	return result
}

// Left goes left
func (c Coordinate) Down() Coordinate {
	result := Coordinate{
		X: c.X,
		Y: c.Y + 1,
	}
	return result
}

type MoveResponse struct {
	Move  string
	Shout string
}

type EndRequest struct {
	Game  string
	Turn  string
	Board string
	You   string
}

type StartRequest struct {
	Game  string `json:"game,omitempty"`
	Turn  string `json:"turn,omitempty"`
	Board string `json:"board,omitempty"`
	You   string `json:"you,omitempty"`
}

type StartResponse struct {
	Color    string
	Headtype string
	Tailtype string
}

type Game struct {
	ID string `json:"id"`
}

type Board struct {
	Height int          `json:"height"`
	Width  int          `json:"width"`
	Food   []Coordinate `json:"food"`
	Snakes []Snake      `json:"snakes"`
}

type Snake struct {
	ID     string       `json:"id"`
	Name   string       `json:"name"`
	Health int          `json:"health"`
	Body   []Coordinate `json:"body"`
	Shout  string       `json:"shout"`
}

// This is the request structure from the gameserver -
// We can pass the whole thing to any functions which are
// computing moves.
type MoveRequest struct {
	Game  `json:"game"`
	Turn  int `json:"turn"`
	Board `json:"board"`
	You   Snake `json:"you"`
}
