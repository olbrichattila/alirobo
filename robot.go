package main

import (
	"alibabarobotgame/internal/defaultconfig"
	"alibabarobotgame/internal/sprite"
)

func (g *game) initRobotSprite() sprite.Sprite {
	paths := []string{
		defaultconfig.CdnUrl + "static/robot/Walk_000.png",
		defaultconfig.CdnUrl + "static/robot/Walk_001.png",
		defaultconfig.CdnUrl + "static/robot/Walk_002.png",
		defaultconfig.CdnUrl + "static/robot/Walk_003.png",
		defaultconfig.CdnUrl + "static/robot/Walk_004.png",
		defaultconfig.CdnUrl + "static/robot/Walk_005.png",
		defaultconfig.CdnUrl + "static/robot/Walk_006.png",
		defaultconfig.CdnUrl + "static/robot/Walk_007.png",
		defaultconfig.CdnUrl + "static/robot/Walk_008.png",
		defaultconfig.CdnUrl + "static/robot/Walk_009.png",
		defaultconfig.CdnUrl + "static/robot/Walk_010.png",
		defaultconfig.CdnUrl + "static/robot/Walk_011.png",
	}

	return sprite.New(
		[]sprite.SpriteImage{
			{Paths: paths},
			{FlipHorizontally: true, Paths: paths},
			{Paths: []string{defaultconfig.CdnUrl + "static/shapesifter.png"}},
		},
		80, 80,
		sprite.SpriteOptions{
			ScreenW:        defaultconfig.ScreenW,
			ScreenH:        defaultconfig.ScreenH,
			Soft:           false,
			SoftY:          60,
			X:              5,
			Y:              200,
			AnimateOnMove:  false,
			AnimationSpeed: 5,
		},
	)
}
