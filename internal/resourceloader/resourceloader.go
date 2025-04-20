// Package resourceloader loads images, sounds form HTTP server async
package resourceloader

import (
	"alibabarobotgame/internal/defaultconfig"
	"alibabarobotgame/internal/gametext"
	imageManager "alibabarobotgame/internal/image"
	"alibabarobotgame/internal/sound"
	"fmt"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	badgeSize     = 25
	resourceCount = 23
)

// Resources loaded
type Resources struct {
	mu                sync.RWMutex
	loadedCount       int
	BadgeEcs          *ebiten.Image
	BadgeFn           *ebiten.Image
	BadgeOss          *ebiten.Image
	BadgeServerless   *ebiten.Image
	BadgeBlockStorage *ebiten.Image
	BadgeCloudBackup  *ebiten.Image
	BadgeCdn          *ebiten.Image
	BadgeApsaraDB     *ebiten.Image
	WallImage         *ebiten.Image
	ServerRoomImage   *ebiten.Image
	OfficeImage       *ebiten.Image
	Office2Image      *ebiten.Image
	Office3Image      *ebiten.Image
	BricksImage       *ebiten.Image
	VBricksImage      *ebiten.Image
	SafeDoorImage     *ebiten.Image
	LadderImage       *ebiten.Image
	BossImage         *ebiten.Image
	OpenScreenImage   *ebiten.Image
	LaunchSndData     []byte
	ExplosionSnd      *audio.Player
	BgMusic           *audio.Player
}

// ResourceLoader encapsulates load logic
type ResourceLoader interface {
	Draw(screen *ebiten.Image)
	Update()
}

type loader struct {
	isLoading     bool
	isLoaded      bool
	resources     *Resources
	readyCallback func(*Resources)
}

// New resource loader
func New(readyCallback func(*Resources)) ResourceLoader {
	return &loader{
		readyCallback: readyCallback,
	}
}

func (l *loader) Update() {
	if l.isLoaded && l.readyCallback != nil {
		l.readyCallback(l.resources)
	}

	if !l.isLoading {
		l.isLoading = true
		l.loadAll()
		return
	}

	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	if l.resources.loadedCount == resourceCount-1 {
		l.isLoaded = true
	}
}

// Draw draws the loading status to the screen
func (l *loader) Draw(screen *ebiten.Image) {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	gametext.Draw(screen, "Loading your game", defaultconfig.ScreenW/2-100, defaultconfig.ScreenH/2-25)
	gametext.Draw(screen, fmt.Sprintf("%d Resources loaded", l.resources.loadedCount), defaultconfig.ScreenW/2-100, defaultconfig.ScreenH/2)
	gametext.Draw(screen, "Please wait...", defaultconfig.ScreenW/2-100, defaultconfig.ScreenH/2+25)
}

func (l *loader) loadAll() {
	l.resources = &Resources{}
	go l.loadBadgeEcs()
	go l.loadBadgeFn()
	go l.loadBadgeOsss()
	go l.loadBadgeServerlesss()
	go l.loadBadgeBlockStorages()
	go l.loadBadgeCloudBackups()
	go l.loadBadgeCdns()
	go l.loadBadgeApsaraDB()
	go l.loadWallImage()
	go l.loadServerRoomImage()
	go l.loadOfficeImage()
	go l.loadOffice2Image()
	go l.loadOffice3Image()
	go l.loadBricksImage()
	go l.loadVBricksImage()
	go l.loadSafeDoorImage()
	go l.loadLadderImage()
	go l.loadBossImage()
	go l.loadOpenScreenImage()
	go l.loadBadgeFireSound()
	go l.loadLaunchSound()
	go l.loadBadgeBgMusic()
}

func (l *loader) loadBadgeEcs() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.BadgeEcs = imageManager.LoadWithSize(defaultconfig.CdnUrl+"static/badges/ecs.png", badgeSize, badgeSize)
	l.resources.loadedCount++
}

func (l *loader) loadBadgeFn() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.BadgeFn = imageManager.LoadWithSize(defaultconfig.CdnUrl+"static/badges/fc-function-calculation.png", badgeSize, badgeSize)
	l.resources.loadedCount++
}

func (l *loader) loadBadgeOsss() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.BadgeOss = imageManager.LoadWithSize(defaultconfig.CdnUrl+"static/badges/object-storage-oss.png", badgeSize, badgeSize)
	l.resources.loadedCount++
}

func (l *loader) loadBadgeServerlesss() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.BadgeServerless = imageManager.LoadWithSize(defaultconfig.CdnUrl+"static/badges/sae-serverless-application.png", badgeSize, badgeSize)
	l.resources.loadedCount++
}

func (l *loader) loadBadgeBlockStorages() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.BadgeBlockStorage = imageManager.LoadWithSize(defaultconfig.CdnUrl+"static/badges/block-storage.png", badgeSize, badgeSize)
	l.resources.loadedCount++
}

func (l *loader) loadBadgeCloudBackups() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.BadgeCloudBackup = imageManager.LoadWithSize(defaultconfig.CdnUrl+"static/badges/cbs-database-backup.png", badgeSize, badgeSize)
	l.resources.loadedCount++
}

func (l *loader) loadBadgeCdns() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.BadgeCdn = imageManager.LoadWithSize(defaultconfig.CdnUrl+"static/badges/cdn.png", badgeSize, badgeSize)
	l.resources.loadedCount++
}

func (l *loader) loadBadgeApsaraDB() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.BadgeApsaraDB = imageManager.LoadWithSize(defaultconfig.CdnUrl+"static/badges/mysql-cloud-database-mysql-version.png", badgeSize, badgeSize)
	l.resources.loadedCount++
}

func (l *loader) loadWallImage() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.WallImage = imageManager.Load(defaultconfig.CdnUrl + "static/bg/1.png")
	l.resources.loadedCount++
}

func (l *loader) loadServerRoomImage() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.ServerRoomImage = imageManager.LoadWithSize(defaultconfig.CdnUrl+"static/servers/server-room.png", 400, 265)
	l.resources.loadedCount++
}

func (l *loader) loadOfficeImage() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.OfficeImage = imageManager.LoadWithSize(defaultconfig.CdnUrl+"static/office.png", 640, 480)
	l.resources.loadedCount++
}

func (l *loader) loadOffice2Image() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.Office2Image = imageManager.LoadWithSize(defaultconfig.CdnUrl+"static/office-2.png", 320, 200)
	l.resources.loadedCount++
}

func (l *loader) loadOffice3Image() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.Office3Image = imageManager.LoadWithSize(defaultconfig.CdnUrl+"static/office-3.png", 799, 399)
	l.resources.loadedCount++
}

func (l *loader) loadBricksImage() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.BricksImage = imageManager.Load(defaultconfig.CdnUrl + "static/bg/bricks.png")
	l.resources.loadedCount++
}

func (l *loader) loadVBricksImage() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.VBricksImage = imageManager.Load(defaultconfig.CdnUrl + "static/bg/vbricks.png")
	l.resources.loadedCount++
}

func (l *loader) loadSafeDoorImage() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.SafeDoorImage = imageManager.ReplicateImageVertically(imageManager.Load(defaultconfig.CdnUrl+"static/door.png"), 10)
	l.resources.loadedCount++
}

func (l *loader) loadLadderImage() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.LadderImage = imageManager.ReplicateImageVertically(
		imageManager.LoadWithSize(defaultconfig.CdnUrl+"static/ladder.png", 40, 40),
		40,
	)
	l.resources.loadedCount++
}

func (l *loader) loadBossImage() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.BossImage = imageManager.LoadWithSize(defaultconfig.CdnUrl+"static/boss.png", 213, 327)
	l.resources.loadedCount++
}

// OpenScreenImage: imageManager.LoadWithSize(defaultconfig.CdnUrl + "static/alirobo.png", defaultconfig.ScreenW, defaultconfig.ScreenH),
func (l *loader) loadOpenScreenImage() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.OpenScreenImage = imageManager.LoadWithSize(defaultconfig.CdnUrl+"static/alirobo.png", defaultconfig.ScreenW, defaultconfig.ScreenH)
	l.resources.loadedCount++
}

func (l *loader) loadBadgeFireSound() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.ExplosionSnd, _ = sound.LoadMp3Sound(defaultconfig.CdnUrl + "static/sound/fire.mp3")
	l.resources.loadedCount++
}

func (l *loader) loadLaunchSound() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.LaunchSndData, _ = sound.LoadMp3SoundData(defaultconfig.CdnUrl + "static/sound/rlauncher.mp3")
	l.resources.loadedCount++
}

func (l *loader) loadBadgeBgMusic() {
	l.resources.mu.Lock()
	defer l.resources.mu.Unlock()
	l.resources.BgMusic, _ = sound.LoadMp3Sound(defaultconfig.CdnUrl + "static/sound/music.mp3")
	l.resources.loadedCount++
}
