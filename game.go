package main

import (
	"alibabarobotgame/internal/gametext"
	"alibabarobotgame/internal/popup"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *game) updateGame() {
	if len(g.popupLines) > 0 {
		g.popupLineOk.Update()
		if ebiten.IsKeyPressed(ebiten.KeyEnter) || ebiten.IsKeyPressed(ebiten.KeyKPEnter) {
			g.popupLines = nil
		}
	}

	g.underGroundFacility.Update()
	g.playBgMusic()
	g.handleFullScreen()
	g.handleShowHint()
	g.handleNavigation()
	g.handleShapeShift()
	g.handleBadgeClick()
	g.timerWidget.Update()
}

func (g *game) drawGame(screen *ebiten.Image) {
	if len(g.popupLines) > 0 {
		popup.Show(screen, g.popupLines)
		g.popupLineOk.Render(screen)
		return
	}

	g.facilityResult = g.underGroundFacility.Draw(
		screen,
		g.sprites.robot.GetX(),
		g.sprites.robot.GetY(),
		g.gameStatus.facilityLevel,
	)

	if g.facilityResult.FloorPos > 0 {
		g.sprites.robot.SetY(g.facilityResult.FloorPos)
		if g.facilityResult.NearLeft || g.facilityResult.NearLeftSafeDoor {
			g.sprites.robot.MoveX(2)
		}

		if g.facilityResult.NearRight || g.facilityResult.NearRightSafeDoor {
			g.sprites.robot.MoveX(-2)
		}
	}
	g.sprites.robot.Render(screen)
	g.drawCollectedPasswordBadges(screen)
	g.timerWidget.Render(screen)

	gametext.Draw(screen, fmt.Sprintf("Lives: %d", g.gameStatus.lives), 5, 25)
}
