// Package resourceloader loads images, sounds form HTTP server async
package resourceloader

import (
	"alibabarobotgame/internal/defaultconfig"
	"alibabarobotgame/internal/gametext"
	imageManager "alibabarobotgame/internal/image"
	"alibabarobotgame/internal/sound"
	"fmt"
	"image"
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	badgeSize     = 25
	resourceCount = 23
)

type audioDataResources map[string][]byte
type audioResources map[string]*audio.Player
type imageResources map[string]*ebiten.Image

var imageResourcesFiles = map[string]string{
	"BadgeEcs":          defaultconfig.CdnUrl + "static/badges/ecs.png",
	"BadgeFn":           defaultconfig.CdnUrl + "static/badges/fc-function-calculation.png",
	"BadgeOss":          defaultconfig.CdnUrl + "static/badges/object-storage-oss.png",
	"BadgeServerless":   defaultconfig.CdnUrl + "static/badges/sae-serverless-application.png",
	"BadgeBlockStorage": defaultconfig.CdnUrl + "static/badges/block-storage.png",
	"BadgeCloudBackup":  defaultconfig.CdnUrl + "static/badges/cbs-database-backup.png",
	"BadgeCdn":          defaultconfig.CdnUrl + "static/badges/cdn.png",
	"BadgeApsaraDB":     defaultconfig.CdnUrl + "static/badges/mysql-cloud-database-mysql-version.png",
	"WallImage":         defaultconfig.CdnUrl + "static/bg/1.png",
	"ServerRoomImage":   defaultconfig.CdnUrl + "static/servers/server-room.png",
	"OfficeImage":       defaultconfig.CdnUrl + "static/office.png",
	"Office2Image":      defaultconfig.CdnUrl + "static/office-2.png",
	"Office3Image":      defaultconfig.CdnUrl + "static/office-3.png",
	"BricksImage":       defaultconfig.CdnUrl + "static/bg/bricks.png",
	"VBricksImage":      defaultconfig.CdnUrl + "static/bg/vbricks.png",
	"SafeDoorImage":     defaultconfig.CdnUrl + "static/door.png",
	"LadderImage":       defaultconfig.CdnUrl + "static/ladder.png",
	"BossImage":         defaultconfig.CdnUrl + "static/boss.png",
	"OpenScreenImage":   defaultconfig.CdnUrl + "static/alirobo2.png",
}

var audioResourceFiles = map[string]string{
	"ExplosionSnd": defaultconfig.CdnUrl + "static/sound/rlauncher.mp3",
	"BgMusic":      defaultconfig.CdnUrl + "static/sound/music.mp3",
}

var audioDataResourceFiles = map[string]string{
	"LaunchSndData": defaultconfig.CdnUrl + "static/sound/fire.mp3",
}

// ResourceLoader encapsulates load logic
type ResourceLoader interface {
	Draw(screen *ebiten.Image)
	Update()
	GetAudioDataResource(key string) []byte
	GetAudioResource(key string) *audio.Player
	GetImageResource(key string) *ebiten.Image
}

type loader struct {
	isLoaded           bool
	isLoading          bool
	bgBar              *ebiten.Image
	gBar               *ebiten.Image
	audioDataResources map[string][]byte
	audioResources     map[string]*audio.Player
	imageResources     map[string]*ebiten.Image
	readyCallback      func(resourceLoader ResourceLoader)
	mu                 sync.RWMutex
}

// New resource loader
func New(readyCallback func(resourceLoader ResourceLoader)) ResourceLoader {
	bgBar := ebiten.NewImage(defaultconfig.ScreenW-100, 60)
	bgBar.Fill(color.RGBA{255, 255, 255, 255})

	gBar := ebiten.NewImage(defaultconfig.ScreenW-110, 40)
	gBar.Fill(color.RGBA{0, 0, 255, 255})

	return &loader{
		readyCallback:      readyCallback,
		bgBar:              bgBar,
		gBar:               gBar,
		audioDataResources: map[string][]byte{},
		audioResources:     map[string]*audio.Player{},
		imageResources:     map[string]*ebiten.Image{},
	}
}

func (l *loader) Update() {
	if l.isLoaded && l.readyCallback != nil {
		l.readyCallback(l)
		return
	}

	if !l.isLoading {
		l.isLoading = true
		go l.loadAll()
	}
}

// Draw draws the loading status to the screen
func (l *loader) Draw(screen *ebiten.Image) {
	loadedCount := len(l.imageResources) + len(l.audioResources) + len(l.audioDataResources)
	top := float64(300)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(50, top)
	screen.DrawImage(l.bgBar, op)

	barWidth := int(float64(loadedCount) / float64(21) * float64(defaultconfig.ScreenW-120))
	if barWidth > 0 {
		barImg := l.gBar.SubImage(image.Rect(0, 0, barWidth, 40)).(*ebiten.Image)
		op.GeoM.Reset()
		op.GeoM.Translate(55, top+10)
		screen.DrawImage(barImg, op)
	}

	gametext.Draw(screen, "Loading your game", defaultconfig.ScreenW/2-100, defaultconfig.ScreenH/2-25)
	gametext.Draw(screen, fmt.Sprintf("%d Resources loaded", loadedCount), defaultconfig.ScreenW/2-100, defaultconfig.ScreenH/2)
	gametext.Draw(screen, "Please wait...", defaultconfig.ScreenW/2-100, defaultconfig.ScreenH/2+25)
}

func (l *loader) loadAll() {
	for key, url := range imageResourcesFiles {
		if key == "LadderImage" {
			l.loadImage(key, url, 40)
			continue
		}

		l.loadImage(key, url, 0)
	}

	for key, url := range audioResourceFiles {
		l.loadAudioContext(key, url)
	}

	for key, url := range audioDataResourceFiles {
		l.loadAudioData(key, url)
	}

	l.mu.Lock()
	l.isLoaded = true
	l.mu.Unlock()
}

func (l *loader) loadImage(key, url string, replicate int) {
	if replicate == 0 {
		l.mu.Lock()
		l.imageResources[key] = imageManager.Load(url)
		l.mu.Unlock()
		return
	}

	l.mu.Lock()
	l.imageResources[key] = imageManager.ReplicateImageVertically(
		imageManager.LoadWithSize(defaultconfig.CdnUrl+"static/ladder.png", 40, 40),
		replicate,
	)
	l.mu.Unlock()
}

func (l *loader) loadAudioContext(key, url string) {
	context, _ := sound.LoadMp3Sound(url)
	// TODO error handling
	l.mu.Lock()
	l.audioResources[key] = context
	l.mu.Unlock()
}

func (l *loader) loadAudioData(key, url string) {
	data, _ := sound.LoadMp3SoundData(url)
	l.mu.Lock()
	l.audioDataResources[key] = data
	l.mu.Unlock()
}

func (l *loader) GetAudioDataResource(key string) []byte {
	if r, ok := l.audioDataResources[key]; ok {
		return r
	}

	return nil
}

func (l *loader) GetAudioResource(key string) *audio.Player {
	if r, ok := l.audioResources[key]; ok {
		return r
	}

	return nil
}

func (l *loader) GetImageResource(key string) *ebiten.Image {
	if r, ok := l.imageResources[key]; ok {
		return r
	}

	return nil
}
