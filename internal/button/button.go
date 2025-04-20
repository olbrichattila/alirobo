// Package button manages clickable buttons
package button

import (
	"alibabarobotgame/internal/gametext"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Button interface {
	New(caption string, x, y, w, h int, onClick func())
	Remove(caption string)
	Update()
	Render(screen *ebiten.Image)
}

type btnEvent struct {
	background *ebiten.Image
	x, y, w, h int
	onClick    func()
}

type btn struct {
	mouseDown bool
	buttons   map[string]btnEvent
}

func New() Button {
	return &btn{
		buttons: map[string]btnEvent{},
	}
}

func (b *btn) New(caption string, x, y, w, h int, onClick func()) {
	btnBg := ebiten.NewImage(w, h)
	btnBg.Fill(color.RGBA{240, 33, 33, 255})
	gametext.Draw(btnBg, caption, 5, 20)
	b.buttons[caption] = btnEvent{
		background: btnBg,
		x:          x,
		y:          y,
		w:          w,
		h:          h,
		onClick:    onClick,
	}
}

func (b *btn) Remove(caption string) {
	delete(b.buttons, caption)
}

func (b *btn) Update() {
	x, y := ebiten.CursorPosition()

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if b.mouseDown {
			return
		}

		b.mouseDown = true
		for _, btn := range b.buttons {
			if x >= btn.x && x <= btn.x+btn.w && y >= btn.y && y <= btn.y+btn.h {
				if btn.onClick != nil {
					btn.onClick()
				}
			}
		}
	} else {
		b.mouseDown = false
	}
}

func (b *btn) Render(screen *ebiten.Image) {
	for caption, btn := range b.buttons {
		b.renderButton(screen, caption, btn)
	}
}

func (b *btn) renderButton(screen *ebiten.Image, caption string, btn btnEvent) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(btn.x), float64(btn.y))
	screen.DrawImage(btn.background, op)
}
