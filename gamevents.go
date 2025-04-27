package main

import (
	"alibabarobotgame/internal/defaultconfig"
	"alibabarobotgame/internal/messages"
	"fmt"
)

func (g *game) gameEventHandler(event defaultconfig.AlibabaServiceType) {
	if event != defaultconfig.BossRom {
		g.playLaunchSound()
	}

	switch event {
	case defaultconfig.BossRom:
		g.renderBossRoom()
	case defaultconfig.Ecs:
		g.collectPasswordForEcs()
	case defaultconfig.FunctionCompute:
		g.collectPasswordForFunctionCompute()
	case defaultconfig.ServerlessComputing:
		g.collectPasswordForServerlessComputing()
	case defaultconfig.ObjectStorageService:
		g.collectPasswordForObjectStorageService()
	case defaultconfig.BlockStorage:
		g.collectPasswordForBlockStorage()
	case defaultconfig.CloudBackup:
		g.collectPasswordForCloudBackup()
	case defaultconfig.Cdn:
		g.collectPasswordCdn()
	case defaultconfig.ApsaraDB:
		g.collectPasswordForDBStorage()
	default:
		panic(fmt.Sprintf("Missing game event, %v", event))
	}
}

func (g *game) renderBossRoom() {
	if g.gameStatus.openedServerCount == serverRoomCount {
		g.gameStatus.gameMode = modeWon
	}
}

func (g *game) collectPasswordForBadgeType(kind defaultconfig.AlibabaServiceType) {
	for i, badge := range g.badges {
		if badge.kind == kind {
			g.badges[i].isCollected = true
			return
		}
	}
}

func (g *game) collectPasswordForEcs() {
	if g.gameStatus.robotShapeSifted {
		g.collectPasswordForBadgeType(defaultconfig.Ecs)
		g.popupLines = []string{"You have collected a password", "for ECS server room", "Press enter to continue"}
		return
	}

	g.removeLife(messages.NotShapeShifted())
}

func (g *game) collectPasswordForFunctionCompute() {
	if g.gameStatus.robotShapeSifted {
		g.collectPasswordForBadgeType(defaultconfig.FunctionCompute)
		g.popupLines = []string{"You have collected a password for", "Function Compute server room", "Press enter to continue"}
		return
	}

	g.removeLife(messages.NotShapeShifted())
}

func (g *game) collectPasswordForServerlessComputing() {
	g.collectPasswordForBadgeType(defaultconfig.ServerlessComputing)
	g.popupLines = []string{"You have collected a password for", "Serverless compute room", "Press enter to continue"}
	return
}

func (g *game) collectPasswordForObjectStorageService() {
	if g.gameStatus.robotShapeSifted {
		g.collectPasswordForBadgeType(defaultconfig.ObjectStorageService)
		g.popupLines = []string{"You have collected a password for", "Function Object Storage Server room", "Press enter to continue"}
		return
	}

	g.removeLife(messages.NotShapeShifted())
}

func (g *game) collectPasswordForBlockStorage() {
	if g.gameStatus.robotShapeSifted {
		g.collectPasswordForBadgeType(defaultconfig.BlockStorage)
		g.popupLines = []string{"You have collected a password for", "Block Storage Server room", "Press enter to continue"}
		return
	}

	g.removeLife(messages.NotShapeShifted())
}

func (g *game) collectPasswordForCloudBackup() {
	if g.gameStatus.robotShapeSifted {
		g.collectPasswordForBadgeType(defaultconfig.CloudBackup)
		g.popupLines = []string{"You have collected a password for", "Cloud Backup Server room", "Press enter to continue"}
		return
	}

	g.removeLife(messages.NotShapeShifted())
}

func (g *game) collectPasswordCdn() {
	if g.gameStatus.robotShapeSifted {
		g.collectPasswordForBadgeType(defaultconfig.Cdn)
		g.popupLines = []string{"You have collected a password for", "CDN Server room", "Press enter to continue"}
		return
	}

	g.removeLife(messages.NotShapeShifted())
}

func (g *game) collectPasswordForDBStorage() {
	g.collectPasswordForBadgeType(defaultconfig.ApsaraDB)
	g.popupLines = []string{"You have collected a password for", "ApsaraDB for RDS room", "Press enter to continue"}
	return
}

func (g *game) removeLife(messages []string) {
	g.gameStatus.lives--
	if g.gameStatus.lives == 0 {
		g.gameStatus.gameMode = modeLost
		return
	}

	if len(messages) > 0 {
		g.popupLines = messages
	}
}
