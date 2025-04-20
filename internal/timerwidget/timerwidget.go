package timerwidget

import (
	"alibabarobotgame/internal/defaultconfig"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	totalTime = 72000 // 20 minutes
)

type TimerWidget interface {
	Render(screen *ebiten.Image)
	Update()
	Reset()
	TimeLeft() int
}

type widget struct {
	endCallback       func()
	tick              int
	endCallbackCalled bool
	isEnded           bool

	bgBar   *ebiten.Image
	fgBar   *ebiten.Image
	barRect image.Rectangle
}

func New(endCallback func()) TimerWidget {
	w := &widget{
		endCallback: endCallback,
		barRect:     image.Rect(0, 0, defaultconfig.ScreenW-10, 15),
	}

	w.bgBar = ebiten.NewImage(defaultconfig.ScreenW, 15)
	w.bgBar.Fill(color.RGBA{33, 33, 33, 255})

	w.fgBar = ebiten.NewImage(defaultconfig.ScreenW-10, 15)
	w.fgBar.Fill(color.RGBA{0, 33, 255, 255})

	return w
}

func (w *widget) Reset() {
	w.tick = 0
}

func (w *widget) TimeLeft() int {
	return totalTime - w.tick
}

func (w *widget) Update() {
	if w.isEnded && !w.endCallbackCalled {
		w.endCallbackCalled = true
		if w.endCallback != nil {
			w.endCallback()
		}
	}
}

func (w *widget) Render(screen *ebiten.Image) {
	w.tick++
	if w.tick >= totalTime {
		w.tick = totalTime
		w.isEnded = true
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(defaultconfig.ScreenH-15))
	screen.DrawImage(w.bgBar, op)

	barWidth := int(float64(totalTime-w.tick) / float64(totalTime) * float64(defaultconfig.ScreenW-10))
	if barWidth > 0 {
		barImg := w.fgBar.SubImage(image.Rect(0, 0, barWidth, 15)).(*ebiten.Image)
		screen.DrawImage(barImg, op)
	}
}
