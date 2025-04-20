package main

import (
	"alibabarobotgame/internal/defaultconfig"
	"alibabarobotgame/internal/gametext"
	"alibabarobotgame/internal/messages"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *game) updateIntro() {
	g.introButtons.Update()
}

func (g *game) drawIntro(screen *ebiten.Image) {
	if g.introLine > -900 {
		g.introLine = g.introLine - 0.2
	}

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(0, g.introLine)
	screen.DrawImage(g.introImage, op)
	g.introButtons.Render(screen)
}

func (g *game) showOpenScreen() {
	g.gameStatus.gameMode = modeOpenScreen
}

func (*game) buildGameIntroImage() *ebiten.Image {
	intro := messages.GameIntro()
	introLen := len(intro)

	img := ebiten.NewImage(defaultconfig.ScreenW-100, 500+introLen*25)
	img.Fill(color.RGBA{33, 33, 33, 255})

	for i, line := range intro {
		gametext.Draw(img, line, 25, float64(450+i*25))
	}

	return img
}
