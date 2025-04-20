package facility

import (
	"alibabarobotgame/internal/defaultconfig"
	"alibabarobotgame/internal/sprite"
)

func (f *fac) initSprites() {
	f.sprites = spriteList{
		officeMan: f.getOfficeManSprite(),
	}

}

func (f *fac) getOfficeManSprite() sprite.Sprite {
	paths := []string{
		defaultconfig.CdnUrl + "static/person/officeman/1.png",
		defaultconfig.CdnUrl + "static/person/officeman/2.png",
		defaultconfig.CdnUrl + "static/person/officeman/3.png",
		defaultconfig.CdnUrl + "static/person/officeman/4.png",
		defaultconfig.CdnUrl + "static/person/officeman/5.png",
		defaultconfig.CdnUrl + "static/person/officeman/6.png",
	}

	return sprite.New(
		[]sprite.SpriteImage{
			{Paths: paths},
			{FlipHorizontally: true, Paths: paths},
		},
		80, 80,
		sprite.SpriteOptions{
			ScreenW:        defaultconfig.ScreenW,
			ScreenH:        defaultconfig.ScreenH,
			Soft:           false,
			SoftY:          60,
			X:              0,
			Y:              200,
			AnimateOnMove:  false,
			Animate:        true,
			AnimationSpeed: 10,
		},
	)
}
