package liggetest

import (
	"log"
	"time"
	"math/rand"
)

var SIZE_X int32 = 16
var SIZE_Y int32 = 16

var RND = rand.NewSource(int64(time.Now().Nanosecond()))

type Map struct {
	Tiles map[int32]map[int32]*Tile
}

func randbetween(min int64, max int64) int64 {
	diff := max - min
	val := RND.Int63() % (diff + 1)

	return min + val
}

func pick() string {
	which := randbetween(1, 4)

	if which == 1 {
		return "tile"
	} else if which == 2 {
		return "tile2"
	} else if which == 3 {
		return "tile3"
	} else {
		return "tile4"
	}
}

func NewMap() *Map {
	var x int32
	var y int32

	log.Printf("Generating a new map")

	tiles := map[int32]map[int32]*Tile{};

	for x = 0; x < SIZE_X; x++ {
		tiles[x] = map[int32]*Tile{}
		for y = 0; y < SIZE_Y; y++ {
			tiles[x][y] = NewTile(pick())
		}
	}

	m := Map{
		tiles,
	}

	return &m
}