// Package inputbox creates a user input textbox
package inputbox

import (
	"alibabarobotgame/internal/gametext"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputBox interface {
	Update()
	Draw(screen *ebiten.Image, x, y float64)
	Reset()
	Text() string
}

type ib struct {
	input      string
	background *ebiten.Image
}

func New() InputBox {
	boxImg := ebiten.NewImage(250, 30)

	return &ib{
		background: boxImg,
	}
}

func (i *ib) Update() {
	i.background.Fill(color.RGBA{33, 33, 33, 255})
	var runes []rune
	runes = ebiten.AppendInputChars(runes)
	for _, r := range runes {
		if r >= 32 && r <= 126 {
			i.input += string(r)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) && len(i.input) > 0 {
		i.input = i.input[:len(i.input)-1]
	}
}

func (i *ib) Draw(screen *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	gametext.Draw(i.background, i.input, 5, 20)

	op.GeoM.Reset()
	op.GeoM.Translate(x, y)
	screen.DrawImage(i.background, op)
}

func (i *ib) Reset() {
	i.input = ""
}

func (i *ib) Text() string {
	return i.input
}
