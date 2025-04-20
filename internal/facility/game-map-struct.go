package facility

import (
	"alibabarobotgame/internal/defaultconfig"

	"github.com/hajimehoshi/ebiten/v2"
)

type facility struct {
	x      float64
	y      float64
	levels []level
}

type level struct {
	deepness float64
	rooms    []room
}

type room struct {
	bg            *ebiten.Image
	rightWall     bool
	leftWall      bool
	leftSafeDoor  bool
	rightSafeDoor bool
	floor         float64
	ceil          float64
	width         float64
	hintText      []string
	roomType      defaultconfig.AlibabaServiceType
	roomLabel     string
	// Add gap left, right to separate rooms
	passages      []passage
	upperPassages []passage
	render        func()
}

type passage struct {
	pos          float64
	ladderHeight int
}
