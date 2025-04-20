package main

import (
	"alibabarobotgame/internal/gametext"
	"alibabarobotgame/internal/messages"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *game) updateWon() {
	g.inputBox.Update()
	g.winButtons.Update()

}

func (g *game) drawWon(screen *ebiten.Image) {
	winnerText := messages.WinnerText()
	score := g.getScore()
	winnerText = append(winnerText, "Your score is "+strconv.Itoa(score))
	for i, line := range winnerText {
		gametext.Draw(screen, line, 60, float64(60+i*25))
	}

	g.inputBox.Draw(screen, 300, float64(len(winnerText)*25+85))
	g.winButtons.Render(screen)
}

func (g *game) cancelSave() {
	g.init()
}

func (g *game) getScore() int {
	return g.timerWidget.TimeLeft() * g.gameStatus.lives
}

func (g *game) saveToWinnerBoard() {
	name := g.inputBox.Text()
	if len(name) == 0 {
		return
	}

	score := g.getScore()
	g.apiClient.AddScore(name, score)
	g.init()
}
