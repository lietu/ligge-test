package liggetest

import (
	"github.com/lietu/ligge"
)

type Tile struct {
	Src string
	Image *ligge.Image
}

func (t *Tile) Render(x int32, y int32) {
	t.Image.Render(x, y)
}

func NewTile(src string) *Tile {
	var image *ligge.Image = nil

	if Gui != nil {
		image = ligge.GetImage(src)
	}

	t := Tile{
		src,
		image,
	}

	return &t
}