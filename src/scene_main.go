package liggetest

import (
	"log"
	"github.com/lietu/ligge"
	"github.com/veandco/go-sdl2/sdl"
)

type MainScene struct {
	*ligge.BaseScene
}

func (m *MainScene) Activate() {
	if !m.Initialized {
		m.Initialize()
	}

	ligge.PlayMusic("music")
}

func (m *MainScene) Initialize() {
	log.Printf("Initializing MainScene")
	m.InitComponents()
	m.Initialized = true
}

func (m *MainScene) Render() {
	width, height := Gui.GetSize()
	dst := sdl.Rect{0, 0, width, height}
	renderer := Gui.Renderer
	renderer.Clear()

	renderer.SetDrawColor(50, 50, 50, 255)
	renderer.FillRect(&dst)

	var xPos, yPos int32

	var xMin, xMax, yMin, yMax int32

	xMin = 0
	yMin = 0
	xMax = 16
	yMax = 16

	xPos = 0
	for x := xMin; x < xMax; x++ {
		yPos = 0
		for y := yMin; y < yMax; y++ {
			tile := Game.Map.Tiles[x][y]

			tile.Render(xPos, yPos)

			yPos += 128
		}
		xPos += 128
	}

	Game.Char.Render()
}

func (m *MainScene) KeyPress(keycode sdl.Keycode) {
	if keycode == sdl.K_ESCAPE {
		Gui.SetScene("mainmenu")
	}

	hi := false
	if keycode == sdl.K_w {
		Game.Char.Y -= 1
		hi = true
	}

	if keycode == sdl.K_s {
		Game.Char.Y += 1
		hi = true
	}

	if keycode == sdl.K_a {
		Game.Char.X -= 1
		hi = true
	}

	if keycode == sdl.K_d {
		Game.Char.X += 1
		hi = true
	}

	if hi {
		Game.Char.SayHi()
	}
}


func init() {
	m := MainScene{
		ligge.NewBaseScene(),
	}

	ligge.RegisterScene("main", &m)
}

