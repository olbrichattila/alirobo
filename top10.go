package main

import (
	"alibabarobotgame/internal/gametext"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *game) updateTop10() {
	g.top10Buttons.Update()
}

func (g *game) drawTop10(screen *ebiten.Image) {
	if g.userScores == nil {
		gametext.Draw(screen, "Loading user scores", 350, 200)
		return
	}

	if g.userScoresLoadError != "" {
		gametext.Draw(screen, g.userScoresLoadError, 350, 250)
		return
	}

	gametext.Draw(screen, "NAME", 60, 50)
	gametext.Draw(screen, "SCORE", 350, 50)
	gametext.Draw(screen, "DATE", 450, 50)

	for i, score := range g.userScores {
		dateStr := score.CreatedAt
		dt1, err := time.Parse(time.RFC3339, score.CreatedAt)
		if err == nil {
			dateStr = dt1.Format("06-01-02 15:04")
		}
		gametext.Draw(screen, score.Name, 60, float64(100+i*25))
		gametext.Draw(screen, strconv.Itoa(score.Score), 350, float64(100+i*25))
		gametext.Draw(screen, dateStr, 450, float64(100+i*25))
	}
	g.top10Buttons.Render(screen)
}

func (g *game) backFromTop10() {
	g.gameStatus.gameMode = modeOpenScreen
}
