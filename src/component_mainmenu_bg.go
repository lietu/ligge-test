package liggetest

import (
	"github.com/lietu/ligge"
	"github.com/veandco/go-sdl2/sdl"
)

type MainMenuBG struct {
	ligge.BaseSceneComponent
	Background *ligge.Image
}

func (m *MainMenuBG) Init() {
	m.Background = ligge.GetImage("test")
}

func (m *MainMenuBG) Render() {
	renderer := Gui.Renderer
	width, height := Gui.GetSize()

	renderer.SetDrawColor(GRAY.R, GRAY.G, GRAY.B, GRAY.A)

	rect := sdl.Rect{0, 0, width, height}
	renderer.FillRect(&rect)

	m.Background.Render(0, 0)
}

func NewMainmenuBGComponent() *MainMenuBG {
	m := MainMenuBG{
		ligge.NewBaseSceneComponent(),
		nil,
	}

	return &m
}