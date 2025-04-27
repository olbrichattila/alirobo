// Package facility is the underground facility
package facility

import (
	"alibabarobotgame/internal/defaultconfig"
	"alibabarobotgame/internal/gametext"
	imageManager "alibabarobotgame/internal/image"
	"alibabarobotgame/internal/resourceloader"
	"alibabarobotgame/internal/sprite"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type FacilityResult struct {
	FloorPos          float64
	NearLeft          bool
	NearRight         bool
	NearLeftSafeDoor  bool
	NearRightSafeDoor bool
	CanGoDown         bool
	CanGoUp           bool
	IsSwitchingLevel  bool
	HintText          []string
	RoomType          defaultconfig.AlibabaServiceType
}

type Facility interface {
	Update()
	Draw(screen *ebiten.Image, robotX, robotY float64, facilityLevel int) FacilityResult
	Reset()
	GetX() float64
	SetX(x float64)
	RemoveLeftSafeDoor(roomType defaultconfig.AlibabaServiceType)
	RemoveRightSafeDoor(roomType defaultconfig.AlibabaServiceType)
}

type spriteList struct {
	officeMan sprite.Sprite
}

type images struct {
	bricksImage     *ebiten.Image
	vBricksImage    *ebiten.Image
	ladderImage     *ebiten.Image
	safeDoorImage   *ebiten.Image
	bossImage       *ebiten.Image
	wallImage       *ebiten.Image
	officeImage     *ebiten.Image
	office2Image    *ebiten.Image
	office3Image    *ebiten.Image
	serverRoomImage *ebiten.Image
}

type fac struct {
	screen              *ebiten.Image
	images              images
	resourceLoader      resourceloader.ResourceLoader
	sprites             spriteList
	eventCallback       func(defaultconfig.AlibabaServiceType)
	undergroundFacility facility
	isActiveLevel       bool
	inBossRoom          bool
	result              FacilityResult
	hoverOnBadgePickup  defaultconfig.AlibabaServiceType
	spaceKeyIsDown      bool

	// rom rendering properties
	robotX             float64
	robotY             float64
	level              level
	room               room
	left               float64
	roomX              float64
	roomY              float64
	roomXRight         float64
	roomHeight         float64
	roomBottom         float64
	roomBottomRight    float64
	robotXinGrid       float64
	officeWalkerOffset float64
}

func New(eventCallback func(eventCallback defaultconfig.AlibabaServiceType), images resourceloader.ResourceLoader) Facility {
	f := &fac{
		resourceLoader: images,
	}
	f.initImages()
	f.init()
	f.initGameMap()
	f.initSprites()
	f.calculatePassageHeights()
	f.eventCallback = eventCallback
	return f
}

func (f *fac) Reset() {
	f.undergroundFacility.x = 0
	f.undergroundFacility.y = 0
}

func (f *fac) GetX() float64 {
	return f.undergroundFacility.x
}

func (f *fac) SetX(x float64) {
	f.undergroundFacility.x = x
}

func (f *fac) initImages() {
	f.images = images{
		bricksImage:     f.resourceLoader.GetImageResource("BricksImage"),
		vBricksImage:    f.resourceLoader.GetImageResource("VBricksImage"),
		ladderImage:     f.resourceLoader.GetImageResource("LadderImage"),
		safeDoorImage:   f.resourceLoader.GetImageResource("SafeDoorImage"),
		bossImage:       f.resourceLoader.GetImageResource("BossImage"),
		wallImage:       f.resourceLoader.GetImageResource("WallImage"),
		officeImage:     f.resourceLoader.GetImageResource("OfficeImage"),
		office2Image:    f.resourceLoader.GetImageResource("Office2Image"),
		office3Image:    f.resourceLoader.GetImageResource("Office3Image"),
		serverRoomImage: f.resourceLoader.GetImageResource("ServerRoomImage"),
	}
}

func (f *fac) init() {
	f.officeWalkerOffset = 1
}

func (f *fac) Update() {
	keyJustPressed := false
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if !f.spaceKeyIsDown {
			keyJustPressed = true
		}
		f.spaceKeyIsDown = true
	} else {
		f.spaceKeyIsDown = false
	}

	if f.hoverOnBadgePickup > 0 && f.eventCallback != nil && keyJustPressed {
		f.eventCallback(f.hoverOnBadgePickup)
	}

	if f.inBossRoom {
		f.eventCallback(defaultconfig.BossRom)
	}
}

// DrawFacility draws the tunnel system visible part on screen and returns properties related to the player robot
func (f *fac) Draw(screen *ebiten.Image, robotX, robotY float64, facilityLevel int) FacilityResult {
	f.screen = screen
	f.robotX = robotX
	f.robotY = robotY
	f.hoverOnBadgePickup = 0
	f.inBossRoom = false

	f.resetFacilityProps()
	targetY := -f.undergroundFacility.levels[facilityLevel].deepness + 50
	distance := f.undergroundFacility.y - targetY
	if distance > -1 && distance < 1 {
		f.undergroundFacility.y = targetY
	} else {
		f.undergroundFacility.y -= distance / 20
	}

	for i, level := range f.undergroundFacility.levels {
		f.level = level
		f.isActiveLevel = i == facilityLevel
		f.drawLevel()
	}

	f.result.IsSwitchingLevel = f.undergroundFacility.y != targetY
	return f.result
}

// RemoveLeftSafeDoor remove the safe door by room name
func (f *fac) RemoveLeftSafeDoor(roomType defaultconfig.AlibabaServiceType) {
	for il, level := range f.undergroundFacility.levels {
		for ir, room := range level.rooms {
			if room.roomType == roomType {
				f.undergroundFacility.levels[il].rooms[ir].leftSafeDoor = false
			}
		}
	}
}

// RemoveLeftSafeDoor remove the safe door by room name
func (f *fac) RemoveRightSafeDoor(roomType defaultconfig.AlibabaServiceType) {
	for il, level := range f.undergroundFacility.levels {
		for ir, room := range level.rooms {
			if room.roomType == roomType {
				f.undergroundFacility.levels[il].rooms[ir].rightSafeDoor = false
			}
		}
	}
}

func (f *fac) drawLevel() {
	f.left = 0
	f.robotXinGrid = f.robotX - f.undergroundFacility.x
	for _, room := range f.level.rooms {
		f.room = room
		f.drawRoom()
	}
}

func (f *fac) drawRoom() {
	f.initRoomProps()
	if !f.isRoomInScreen() {
		f.left += f.room.width
		return
	}

	f.drawRoomBackground()
	f.drawRightWallIfNecessary()
	f.drawLeftWallIfNecessary()
	f.drawLeftSafeDoorIfNecessary()
	f.drawRightSafeDoorIfNecessary()
	f.drawPassageDown()
	f.drawPassageUp()
	f.drawRoomLabel()
	f.drawNextRoomTypeLabel()
	f.processActiveRoomEvents()

	f.left += f.room.width
}

func (f *fac) initRoomProps() {
	f.roomX = f.left + f.undergroundFacility.x
	f.roomY = f.undergroundFacility.y + f.level.deepness + f.room.ceil
	f.roomXRight = f.roomX + f.room.width
	f.roomHeight = f.room.floor - f.room.ceil
	f.roomBottom = f.roomY + f.roomHeight
	f.roomBottomRight = f.roomBottom + f.room.width
}

func (f *fac) resetFacilityProps() {
	f.result.FloorPos = 0
	f.result.NearLeft = false
	f.result.NearRight = false
	f.result.NearLeftSafeDoor = false
	f.result.NearRightSafeDoor = false
	f.result.CanGoDown = false
	f.result.CanGoUp = false
	f.result.IsSwitchingLevel = false
}

func (f *fac) isRoomInScreen() bool {
	// Blank spacer room
	if f.room.bg == nil {
		return false
	}

	// room outside of the screen left
	if f.roomXRight < 0 && f.roomBottom < 0 {
		return false
	}

	// room outside of the screen right
	if f.roomX > defaultconfig.ScreenW && f.roomY > defaultconfig.ScreenH {
		return false
	}

	return true
}

func (f *fac) drawRoomLabel() {
	if f.room.roomLabel != "" {
		gametext.Draw(f.screen, f.room.roomLabel, f.roomX+25, f.roomY+45)
	}
}

func (f *fac) drawNextRoomTypeLabel() {
	if f.room.roomType != 0 {
		if label, ok := defaultconfig.ServiceDescriptionMap[f.room.roomType]; ok {
			gametext.Draw(f.screen, label+" room >>", f.roomX+25, f.roomY+45)
		}
	}
}

func (f *fac) drawRoomBackground() {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(f.roomX, f.roomY)

	// Background
	f.screen.DrawImage(
		imageManager.CropVertical(f.room.bg, f.room.width, f.roomHeight),
		op,
	)

	croppedBricksImage := imageManager.CropVertical(f.images.bricksImage, f.room.width, 25)

	// Bricks on top
	f.screen.DrawImage(croppedBricksImage, op)

	// Bricks bottom
	op.GeoM.Reset()
	op.GeoM.Translate(f.roomX, f.roomBottom)
	f.screen.DrawImage(croppedBricksImage, op)
}

func (f *fac) drawLeftWallIfNecessary() {
	if !f.room.leftWall {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(f.roomX, f.roomY)
	f.screen.DrawImage(
		imageManager.CropVertical(f.images.vBricksImage, 25, f.roomHeight),
		op,
	)
}

func (f *fac) drawRightWallIfNecessary() {
	if !f.room.rightWall {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(f.roomX+f.room.width-20, f.roomY)
	f.screen.DrawImage(
		imageManager.CropVertical(f.images.vBricksImage, 25, f.roomHeight),
		op,
	)
}

func (f *fac) drawLeftSafeDoorIfNecessary() {
	if !f.room.leftSafeDoor {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(f.roomX, f.roomY)
	f.screen.DrawImage(
		imageManager.CropVertical(f.images.safeDoorImage, 25, f.roomHeight),
		op,
	)
}

func (f *fac) drawRightSafeDoorIfNecessary() {
	if !f.room.rightSafeDoor {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(f.roomX+f.room.width-20, f.roomY)
	f.screen.DrawImage(
		imageManager.CropVertical(f.images.safeDoorImage, 25, f.roomHeight),
		op,
	)
}

func (f *fac) drawPassageDown() {
	op := &ebiten.DrawImageOptions{}
	for _, passage := range f.room.passages {
		absolutePassagePos := f.left + passage.pos
		op.GeoM.Reset()
		op.GeoM.Translate(f.roomX+passage.pos, f.roomBottom)

		img := f.images.ladderImage.SubImage(image.Rect(0, 0, 40, passage.ladderHeight))
		f.screen.DrawImage(img.(*ebiten.Image), op)

		if f.robotXinGrid > absolutePassagePos-30 && f.robotXinGrid < absolutePassagePos+30 && f.isActiveLevel {
			f.result.CanGoDown = true
		}
	}
}

func (f *fac) drawPassageUp() {
	op := &ebiten.DrawImageOptions{}
	for _, passage := range f.room.upperPassages {
		absolutePassagePos := f.left + passage.pos

		op.GeoM.Reset()
		op.GeoM.Translate(f.roomX+passage.pos, f.roomBottom-float64(passage.ladderHeight))

		img := f.images.ladderImage.SubImage(image.Rect(0, 0, 40, passage.ladderHeight))
		f.screen.DrawImage(img.(*ebiten.Image), op)

		if f.robotXinGrid > absolutePassagePos-30 && f.robotXinGrid < absolutePassagePos+30 && f.isActiveLevel {
			f.result.CanGoUp = true
		}
	}
}

func (f *fac) processActiveRoomEvents() {
	if f.robotXinGrid > f.left && f.robotXinGrid < f.left+f.room.width && f.isActiveLevel {
		f.result.RoomType = f.room.roomType
		f.result.HintText = f.room.hintText

		if f.room.leftWall && f.robotXinGrid < f.left+10 {
			f.result.NearLeft = true
		}

		if f.room.rightWall && f.robotXinGrid > f.left+f.room.width-70 {
			f.result.NearRight = true
		}

		if f.room.leftSafeDoor && f.robotXinGrid < f.left+10 {
			f.result.NearLeftSafeDoor = true
		}

		if f.room.rightSafeDoor && f.robotXinGrid > f.left+f.room.width-70 {
			f.result.NearRightSafeDoor = true
		}

		f.result.FloorPos = f.roomY + f.roomHeight - 80

		if f.room.render != nil {
			f.room.render()
		}
	}
}
