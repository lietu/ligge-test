package liggetest

import (
	"github.com/lietu/ligge"
)

var Game *GameState = nil

type GameState struct {
	Char *Character
	Map *Map
}

type Character struct {
	X int32
	Y int32
	Sound *ligge.Sound
	Image *ligge.Image
}

func (c *Character) SayHi() {
	c.Sound.Play()
}

func (c *Character) Render() {
	xPos := c.X * 128
	yPos := c.Y * 128

	c.Image.Render(xPos, yPos)
}

func NewCharacter() *Character {
	var sound *ligge.Sound = nil
	var image *ligge.Image = nil

	if Gui != nil {
		sound = ligge.NewSound("sfx", "hi")
		image = ligge.GetImage("butt")
	}

	c := Character{
		0,
		0,
		sound,
		image,
	}

	return &c
}

func NewGameState() *GameState {
	g := GameState{
		NewCharacter(),
		NewMap(),
	}

	return &g
}

func StartNewGame() {
	Game = NewGameState()
	Gui.SetScene("main")
}
