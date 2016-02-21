package liggetest

import (
	"fmt"
	"github.com/lietu/ligge"
	"github.com/veandco/go-sdl2/sdl"
)

type MainMenuScene struct {
	*ligge.BaseScene
}

func (m *MainMenuScene) Activate() {
	if !m.Initialized {
		m.Initialize()
	}
}

func (m *MainMenuScene) Initialize() {
	m.AddComponent(NewMainmenuBGComponent())
	m.AddComponent(NewMainMenu())
	m.InitComponents()
	m.Initialized = true
}

func (m *MainMenuScene) Render() {
	renderer := Gui.Renderer
	renderer.Clear()
	m.RenderComponents()
}

func (m *MainMenuScene) KeyPress(keycode sdl.Keycode) {
	if keycode == sdl.K_ESCAPE {
		fmt.Printf("Quitting\n")
		Gui.Stop()
	}

	if keycode == sdl.K_RETURN {
		StartNewGame()
	}

	m.BaseScene.KeyPress(keycode)
}

func init() {
	m := MainMenuScene{
		ligge.NewBaseScene(),
	}

	ligge.RegisterScene("mainmenu", &m)
}
