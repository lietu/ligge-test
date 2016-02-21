package liggetest

import (
	"github.com/lietu/ligge"
)

var Gui *ligge.GUI;

func RunClient() {
	ligge.InitializeClientComponents()
	ligge.RunClient(SetGUI)
	ligge.UninitializeClientComponents()
}

func SetGUI() {
	Gui = ligge.Gui
}