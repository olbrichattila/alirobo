package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *game) updateLoading() {
	if !g.resourceLoading {
		g.resourceLoading = true
		g.preInitResources()
	}
	g.resourceLoader.Update()
}

func (g *game) drawLoading(screen *ebiten.Image) {
	g.resourceLoader.Draw(screen)
}
