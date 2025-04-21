package main

import (
	"alibabarobotgame/internal/api"
	"alibabarobotgame/internal/button"
	"alibabarobotgame/internal/defaultconfig"
	"alibabarobotgame/internal/facility"
	"alibabarobotgame/internal/gametext"
	"alibabarobotgame/internal/inputbox"
	"alibabarobotgame/internal/resourceloader"
	"alibabarobotgame/internal/sound"
	"alibabarobotgame/internal/sprite"
	"alibabarobotgame/internal/timerwidget"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type gameMode = int

// Generic defaults
const (
	levelCount         = 6
	maxDoorOpenRetries = 3
	serverRoomCount    = 8
)

// Badge defaults
const (
	badgeSize             = 25
	buttonHeight          = 25
	badgeDistance float64 = 35
	badgeTop      float64 = 10
)

// Game modes
const (
	modeLoading gameMode = iota
	modeOpenScreen
	modeIntro
	modeTop10
	modeGame
	modeWon
	modeLost
)

type Game interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
}

type badge struct {
	kind        defaultconfig.AlibabaServiceType
	img         *ebiten.Image
	x           float64
	isCollected bool
}

type gameStatus struct {
	gameMode          gameMode
	facilityLevel     int
	robotShapeSifted  bool
	lives             int
	score             int
	openedServerCount int
	openRetries       int
}

type keyboardPressTrack struct {
	arrowUp   bool
	arrowDown bool
}

type sprites struct {
	robot sprite.Sprite
}

type game struct {
	apiClient           api.APIClient
	resourceLoading     bool
	openScreenImage     *ebiten.Image
	resourceLoader      resourceloader.ResourceLoader
	userScores          api.UserScores
	userScoresLoadError string
	underGroundFacility facility.Facility
	facilityResult      facility.FacilityResult
	timerWidget         timerwidget.TimerWidget
	inputBox            inputbox.InputBox
	winButtons          button.Button
	openScreenButtons   button.Button
	introButtons        button.Button
	top10Buttons        button.Button
	sprites             sprites
	badges              []badge
	gameStatus          gameStatus
	keyboard            keyboardPressTrack
	popupLines          []string
	badgeFlashCnt       int
	mouseDown           bool
	introLine           float64
	introImage          *ebiten.Image
}

func main() {
	ebiten.SetWindowSize(defaultconfig.ScreenW, defaultconfig.ScreenH)
	ebiten.SetWindowTitle("AliRobo Guardian of the Cloud")

	if err := ebiten.RunGame(New()); err != nil {
		log.Fatal(err)
	}
}

func New() Game {
	g := &game{}
	return g
}

func (g *game) preInit() {
	g.loadSprites()
	g.preInitButtons()
	g.preInitAPI()
	g.preInitTimer()
	g.preInitIntroImage()
}

func (g *game) preInitIntroImage() {
	g.introImage = g.buildGameIntroImage()
}

func (g *game) preInitTimer() {
	g.timerWidget = timerwidget.New(func() {
		g.gameStatus.lives = 0
		g.gameStatus.gameMode = modeLost
	})
}

func (g *game) preInitAPI() {
	g.apiClient = api.New()
}

func (g *game) preInitResources() {
	g.resourceLoader = resourceloader.New(func(resourceLoader resourceloader.ResourceLoader) {
		g.underGroundFacility = facility.New(g.gameEventHandler, resourceLoader)
		g.openScreenImage = resourceLoader.GetImageResource("OpenScreenImage")
		g.preInit()
		g.init()
		g.loadBadges()
		g.gameStatus.gameMode = modeOpenScreen
	})
}

func (g *game) preInitButtons() {
	g.inputBox = inputbox.New()

	g.winButtons = button.New()
	g.winButtons.New("Cancel", 100, 400, 65, buttonHeight, g.cancelSave)
	g.winButtons.New("Save", 600, 400, 55, buttonHeight, g.saveToWinnerBoard)

	g.openScreenButtons = button.New()
	g.openScreenButtons.New("Play the game", 100, 425, 130, buttonHeight, g.startGame)
	g.openScreenButtons.New("Intro", 400, 425, 50, buttonHeight, g.showIntro)
	g.openScreenButtons.New("Display score board", 600, 425, 175, buttonHeight, g.showScoreBoard)

	g.introButtons = button.New()
	g.introButtons.New("Back", defaultconfig.ScreenW-80, defaultconfig.ScreenH-40, 50, 28, g.showOpenScreen)

	g.top10Buttons = button.New()
	g.top10Buttons.New("Back", defaultconfig.ScreenW/2-50, defaultconfig.ScreenH-40, 50, 28, g.backFromTop10)
}

func (g *game) init() {
	g.gameStatus.gameMode = modeLoading
	if g.underGroundFacility != nil {
		g.underGroundFacility.Reset()
	}

	g.initBadges()
	g.userScores = nil
	g.userScoresLoadError = ""
	g.gameStatus.facilityLevel = 0
	g.gameStatus.openedServerCount = 0
	g.gameStatus.score = 0
	g.gameStatus.lives = 5
	g.gameStatus.robotShapeSifted = false
	g.gameStatus.openRetries = 0
	g.sprites.robot.SetX(100)
	g.sprites.robot.SetY(200)
	g.timerWidget.Reset()

	// For testing stages without playing the game
	//g.gameStatus.gameMode = modeGame
	// g.gameStatus.gameMode = modeWon
	//g.gameStatus.gameMode = modeLost
}

func (g *game) Update() error {
	switch g.gameStatus.gameMode {
	case modeLoading:
		g.updateLoading()
	case modeOpenScreen:
		g.updateOpenScreen()
	case modeIntro:
		g.updateIntro()
	case modeTop10:
		g.updateTop10()
	case modeWon:
		g.updateWon()
	case modeLost:
		g.updateLost()
	case modeGame:
		g.updateGame()
	default:
		panic("Game status not implemented")
	}

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	switch g.gameStatus.gameMode {
	case modeLoading:
		g.drawLoading(screen)
	case modeOpenScreen:
		g.drawOpenScreen(screen)
	case modeIntro:
		g.drawIntro(screen)
	case modeTop10:
		g.drawTop10(screen)
	case modeWon:
		g.drawWon(screen)
	case modeLost:
		g.drawLost(screen)
	case modeGame:
		g.drawGame(screen)
	default:
		panic("Game status not implemented")
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return defaultconfig.ScreenW, defaultconfig.ScreenH
}

func (g *game) handleShowHint() {
	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		if len(g.facilityResult.HintText) > 0 {
			g.popupLines = append(g.facilityResult.HintText, "", "Press enter to continue")
		}
	}
}

func (*game) handleFullScreen() {
	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
		if !ebiten.IsFullscreen() {
			ebiten.SetWindowSize(defaultconfig.ScreenW, defaultconfig.ScreenH)
		}
	}
}

func (g *game) handleNavigation() {
	if g.facilityResult.IsSwitchingLevel {
		g.sprites.robot.Soft(true)
		return
	}

	g.sprites.robot.Soft(false)
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.sprites.robot.SetCollection(0)
		g.gameStatus.robotShapeSifted = false
		g.sprites.robot.Animate(true)
		if g.sprites.robot.GetX() < 400 {
			g.sprites.robot.MoveX(2)
		} else {
			g.underGroundFacility.SetX(g.underGroundFacility.GetX() - 5)
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.sprites.robot.SetCollection(1)
		g.gameStatus.robotShapeSifted = false
		g.sprites.robot.Animate(true)
		if g.sprites.robot.GetX() < 50 {
			g.underGroundFacility.SetX(g.underGroundFacility.GetX() + 5)
		} else {
			g.sprites.robot.MoveX(-2)

		}
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if !g.keyboard.arrowUp {
			g.keyboard.arrowUp = true
			g.sprites.robot.Animate(true)
			if g.facilityResult.CanGoUp && g.gameStatus.facilityLevel > 0 {
				g.gameStatus.facilityLevel--
			}
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		if !g.keyboard.arrowDown {
			g.keyboard.arrowDown = true
			g.sprites.robot.Animate(true)
			if g.facilityResult.CanGoDown && g.gameStatus.facilityLevel < levelCount {
				g.gameStatus.facilityLevel++
			}
		}
	} else {
		g.keyboard.arrowDown = false
		g.keyboard.arrowUp = false
		g.sprites.robot.Animate(false)
	}
}

func (g *game) handleShapeShift() {
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.gameStatus.robotShapeSifted = !g.gameStatus.robotShapeSifted
		if g.gameStatus.robotShapeSifted {
			g.sprites.robot.SetCollection(2)
		} else {
			g.sprites.robot.SetCollection(0)
		}
	}
}

func (g *game) handleBadgeClick() {
	x, y := ebiten.CursorPosition()

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if g.mouseDown {
			return
		}

		g.mouseDown = true
		for i, badge := range g.badges {
			if !badge.isCollected {
				continue
			}

			if x >= int(badge.x) && x <= int(badge.x+badgeSize) && y >= int(badgeTop) && y <= int(badgeTop+badgeSize) {
				if g.facilityResult.RoomType != 0 {
					if badge.kind != g.facilityResult.RoomType {
						g.gameStatus.openRetries++
						if g.gameStatus.openRetries == maxDoorOpenRetries {
							g.gameStatus.openRetries = 0
							g.removeLife(
								[]string{
									"You have chosen an incorrect badge",
									"You can try " + strconv.Itoa(maxDoorOpenRetries-g.gameStatus.openRetries) + " times",
									"",
									"Press enter to continue",
								},
							)
						}

						return
					}

					g.underGroundFacility.RemoveRightSafeDoor(g.facilityResult.RoomType)
					g.badges[i].isCollected = false
					g.gameStatus.openedServerCount++
					return
				}
			}
		}
	} else {
		g.mouseDown = false
	}
}

func (g *game) drawCollectedPasswordBadges(screen *ebiten.Image) {
	if g.facilityResult.NearRightSafeDoor {
		gametext.Draw(screen, "Select >>", defaultconfig.ScreenW-(badgeDistance*8)-100, 30)
	}

	if g.allPasswordsCollected() {
		if g.badgeFlashCnt == 100 {
			g.badgeFlashCnt = 0
		} else {
			g.badgeFlashCnt++
			if g.badgeFlashCnt < 50 {
				return
			}
		}
	}
	op := &ebiten.DrawImageOptions{}

	for _, badge := range g.badges {
		if badge.isCollected {
			op.GeoM.Reset()
			op.GeoM.Translate(badge.x, badgeTop)
			screen.DrawImage(badge.img, op)
		}
	}
}

func (g *game) allPasswordsCollected() bool {
	for _, badge := range g.badges {
		if !badge.isCollected {
			return false
		}
	}

	return true
}

func (g *game) loadBadges() {
	var x float64 = defaultconfig.ScreenW - 40
	g.badges = []badge{
		{
			kind: defaultconfig.Ecs,
			img:  g.resourceLoader.GetImageResource("BadgeEcs"),
			x:    x,
		},
		{
			kind: defaultconfig.FunctionCompute,
			img:  g.resourceLoader.GetImageResource("BadgeFn"),
			x:    x - badgeDistance,
		},
		{

			kind: defaultconfig.ObjectStorageService,
			img:  g.resourceLoader.GetImageResource("BadgeOss"),
			x:    x - (2 * badgeDistance),
		},
		{
			kind: defaultconfig.ServerlessComputing,
			img:  g.resourceLoader.GetImageResource("BadgeServerless"),
			x:    x - (3 * badgeDistance),
		},
		{
			kind: defaultconfig.BlockStorage,
			img:  g.resourceLoader.GetImageResource("BadgeBlockStorage"),
			x:    x - (4 * badgeDistance),
		},
		{
			kind: defaultconfig.CloudBackup,
			img:  g.resourceLoader.GetImageResource("BadgeCloudBackup"),
			x:    x - (5 * badgeDistance),
		},
		{
			kind: defaultconfig.Cdn,
			img:  g.resourceLoader.GetImageResource("BadgeCdn"),
			x:    x - (6 * badgeDistance),
		},
		{
			kind: defaultconfig.ApsaraDB,
			img:  g.resourceLoader.GetImageResource("BadgeApsaraDB"),
			x:    x - (7 * badgeDistance),
		},
	}
}

func (g *game) initBadges() {
	for i, _ := range g.badges {
		g.badges[i].isCollected = false
	}
}

func (g *game) loadSprites() {
	g.sprites = sprites{
		robot: g.initRobotSprite(),
	}
}

func (g *game) playExplosionSound() {
	sound.Play(g.resourceLoader.GetAudioResource("ExplosionSnd"))
}

func (g *game) playLaunchSound() {
	sound.PlayNewFromData(g.resourceLoader.GetAudioDataResource("LaunchSndData"))
}

func (g *game) playBgMusic() {
	sound.Play(g.resourceLoader.GetAudioResource("BgMusic"))
}
