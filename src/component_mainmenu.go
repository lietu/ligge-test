package liggetest

import (
	"log"
	"github.com/lietu/ligge"
	"github.com/veandco/go-sdl2/sdl"
)

type MainMenu struct {
	ligge.BaseSceneComponent
	Background *ligge.Image
	Start *ligge.Text
	Quit *ligge.Text
	LastKey string
}

func (m *MainMenu) Init() {
	m.Background = ligge.GetImage("test")
	m.Start = ligge.CreateText("dejavu32", "Start", ORANGE)
	m.Quit = ligge.CreateText("dejavu32", "Quit", ORANGE)
}

func (m *MainMenu) Render() {
	m.Start.Render(100, 100)
	m.Quit.Render(100, 200)

	key := ligge.CreateText("dejavu32", m.LastKey, RED)
	key.Render(50, 50)
}

func (m *MainMenu) KeyPress(code sdl.Keycode) {
	m.LastKey = sdl.GetKeyName(code)
	log.Printf("%s", m.LastKey)
}

func NewMainMenu() *MainMenu {
	m := MainMenu{
		ligge.NewBaseSceneComponent(),
		nil,
		nil,
		nil,
		" ",
	}

	return &m
}