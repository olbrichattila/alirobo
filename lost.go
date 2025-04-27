package main

import (
	"alibabarobotgame/internal/messages"
	"alibabarobotgame/internal/popup"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *game) updateLost() {
	g.popupLines = messages.LooserText()
	if ebiten.IsKeyPressed(ebiten.KeyEnter) || ebiten.IsKeyPressed(ebiten.KeyKPEnter) {
		g.popupLines = nil
		g.init()
	}
}

func (g *game) drawLost(screen *ebiten.Image) {
	if len(g.popupLines) > 0 {
		popup.Show(screen, g.popupLines)
		return
	}
}
