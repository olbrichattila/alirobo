package main

import "github.com/hajimehoshi/ebiten/v2"

func (g *game) updateOpenScreen() {
	g.openScreenButtons.Update()
}

func (g *game) drawOpenScreen(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(g.resources.OpenScreenImage, op)
	g.openScreenButtons.Render(screen)
}

func (g *game) startGame() {
	g.init()
	g.gameStatus.gameMode = modeGame
}

func (g *game) showIntro() {
	g.introLine = 0
	g.gameStatus.gameMode = modeIntro
}

func (g *game) showScoreBoard() {
	g.userScoresLoadError = ""
	g.userScores = nil
	scores, err := g.apiClient.Top10()
	if err != nil {
		g.userScoresLoadError = err.Error()
	}
	g.userScores = scores
	g.gameStatus.gameMode = modeTop10
}
