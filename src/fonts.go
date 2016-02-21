package liggetest

import (
	"github.com/lietu/ligge"
)

var DEJAVU_SANS = "../assets/fonts/dejavu-fonts-ttf-2.35/ttf/DejavuSans.ttf"

func init() {
	ligge.RegisterFont("dejavu14", 14, DEJAVU_SANS)
	ligge.RegisterFont("dejavu32", 32, DEJAVU_SANS)
}