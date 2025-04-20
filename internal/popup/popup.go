// Package pupup shows a popup window this text
package popup

import (
	"alibabarobotgame/internal/defaultconfig"
	"alibabarobotgame/internal/gametext"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func Show(screen *ebiten.Image, lines []string) {
	screen.Fill(color.RGBA{R: 255, G: 165, B: 0, A: 255})

	vector.DrawFilledRect(
		screen,
		50,
		50,
		defaultconfig.ScreenW-100,
		defaultconfig.ScreenH-100,
		color.RGBA{R: 33, G: 33, B: 33, A: 255},
		false,
	)

	for i, line := range lines {
		gametext.Draw(screen, line, 80, float64(80+i*30))
	}
}
