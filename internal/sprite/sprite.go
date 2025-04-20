package sprite

import (
	imageManager "alibabarobotgame/internal/image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func New(spriteImages []SpriteImage, w, h int, options SpriteOptions) Sprite {
	if len(spriteImages) == 0 {
		panic("Sprite require at least one image")
	}

	sprite := &sprite{
		options: options,
		imgs:    make([][]*ebiten.Image, len(spriteImages)),
	}

	for i, spriteImage := range spriteImages {
		for _, imageName := range spriteImage.Paths {
			img, width, height, _ := imageManager.RescaleImageToFit(imageName, w, h)
			if spriteImage.FlipVertically {
				img = imageManager.FlipImageVertically(img)
			}

			if spriteImage.FlipHorizontally {
				img = imageManager.FlipImageHorizontally(img)
			}

			sprite.imgs[i] = append(sprite.imgs[i], img)
			sprite.width = width
			sprite.height = height
		}
	}

	sprite.init()
	return sprite
}

type SpriteImage struct {
	FlipHorizontally bool
	FlipVertically   bool
	Paths            []string
}

type Sprite interface {
	Render(screen *ebiten.Image)
	SetId(id int)
	SetCollection(id int)
	Id() int
	GetImg() *ebiten.Image
	GetX() float64
	GetY() float64
	GetWidth() float64
	GetHeight() float64
	MoveX(x float64) bool
	MoveY(y float64) bool
	SetX(x float64) bool
	SetY(y float64) bool
	Animate(animate bool)
	Soft(soft bool)
	IsMoving() bool
	RunAfterAnimation()
	Close()
}

type sprite struct {
	options                     SpriteOptions
	collectionIdx               int
	imgIndex                    int
	imgs                        [][]*ebiten.Image
	width                       int
	height                      int
	x                           float64
	y                           float64
	currX                       float64
	currY                       float64
	closed                      bool
	animationCount              int
	inAfterAnimation            bool
	afterAnimationImageId       int
	afterAnimationImageDelayCnt int
}

type SpriteOptions struct {
	Id                           int
	Soft                         bool
	ScreenW                      int
	ScreenH                      int
	X                            float64
	Y                            float64
	SoftX                        float64
	SoftY                        float64
	SoftSpeedUp                  bool
	Animate                      bool
	AnimationSpeed               int
	AnimateOnMove                bool
	CollisionSprites             []Sprite
	CollisionCallback            func(Sprite, []Sprite)
	AfterAnimationAnimationDelay int
	AfterAnimationImages         []*ebiten.Image
	AfterAnimationCallback       func(Sprite)
}

func (s *sprite) init() {
	if s.options.ScreenW == 0 {
		s.options.ScreenW = 640
	}

	if s.options.ScreenH == 0 {
		s.options.ScreenH = 480
	}

	if s.options.AnimationSpeed == 0 {
		s.options.AnimationSpeed = 50
	}

	s.animationCount = int(s.options.AnimationSpeed)

	if s.options.SoftX == 0 {
		s.options.SoftX = 20
	}

	if s.options.SoftY == 0 {
		s.options.SoftY = 20
	}

	s.SetX(s.options.X)
	s.SetY(s.options.Y)
}

func (s *sprite) Render(screen *ebiten.Image) {
	if s.closed {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.currX, s.currY)

	if s.inAfterAnimation {
		if s.afterAnimationImageId == len(s.options.AfterAnimationImages) {
			s.inAfterAnimation = false
			if s.options.AfterAnimationCallback != nil {
				s.options.AfterAnimationCallback(s)
			}

			return
		}

		screen.DrawImage(s.options.AfterAnimationImages[s.afterAnimationImageId], op)
		if s.afterAnimationImageDelayCnt == s.options.AfterAnimationAnimationDelay {
			s.afterAnimationImageId++
			s.afterAnimationImageDelayCnt = 0
		} else {
			s.afterAnimationImageDelayCnt++
		}

		return
	}

	screen.DrawImage(s.imgs[s.collectionIdx][s.imgIndex], op)
	s.collisionDetection()

	s.correctSoftPos()
	s.stepToNextImage()
}

func (s *sprite) RunAfterAnimation() {
	s.afterAnimationImageId = 0
	s.inAfterAnimation = true
}

func (s *sprite) SetId(id int) {
	if s.closed {
		return
	}
	s.options.Id = id
}

func (s *sprite) SetCollection(id int) {
	if s.collectionIdx != id && id < len(s.imgs) {
		s.collectionIdx = id
		s.imgIndex = 0
	}
}

func (s *sprite) Id() int {
	return s.options.Id
}

func (s *sprite) GetImg() *ebiten.Image {
	return s.imgs[s.collectionIdx][s.imgIndex]
}

func (s *sprite) GetX() float64 {
	return s.currX
}

func (s *sprite) GetY() float64 {
	return s.currY
}

func (s *sprite) GetWidth() float64 {
	return float64(s.width)
}

func (s *sprite) GetHeight() float64 {
	return float64(s.height)
}

func (s *sprite) MoveX(x float64) bool {
	return s.SetX(s.x + x)
}

func (s *sprite) MoveY(y float64) bool {
	return s.SetY(s.y + y)
}

func (s *sprite) SetX(x float64) bool {
	if s.closed {
		return false
	}

	if x < 0 || x > float64(s.options.ScreenW-s.width) {
		return false
	}
	s.x = x
	if !s.options.Soft {
		s.currX = x
	}

	return true
}

func (s *sprite) SetY(y float64) bool {
	if s.closed {
		return false
	}

	if y < 0 || y > float64(s.options.ScreenH-s.height) {
		return false
	}
	s.y = y
	if !s.options.Soft {
		s.currY = y
	}

	return true
}

func (s *sprite) Animate(animate bool) {
	s.options.Animate = animate
}

func (s *sprite) Soft(soft bool) {
	s.options.Soft = soft
}

func (s *sprite) IsMoving() bool {
	return !(s.currY == s.y && s.currX == s.x)
}

func (s *sprite) Close() {
	s.closed = true
}

func (s *sprite) collisionDetection() {
	if s.options.CollisionCallback == nil {
		return
	}

	result := []Sprite{}
	for _, otherSprite := range s.options.CollisionSprites {
		if s.isCollided(otherSprite) {
			result = append(result, otherSprite)
		}
	}

	if len(result) > 0 {
		s.options.CollisionCallback(s, result)
	}
}

func (s *sprite) isCollided(otherSprite Sprite) bool {
	if otherSprite == nil {
		return false
	}

	otherSpriteX := otherSprite.GetX()
	otherSpriteY := otherSprite.GetY()

	if s.currX > otherSpriteX+otherSprite.GetWidth() || s.currX+float64(s.width) < otherSpriteX {
		return false
	}

	if s.currY > otherSpriteY+otherSprite.GetHeight() || s.currY+float64(s.height) < otherSpriteY {
		return false
	}

	return true
}

func (s *sprite) correctSoftPos() {
	if !s.options.Soft {
		return
	}

	if s.options.SoftSpeedUp {
		if s.currX == s.x && s.currY == s.y {
			return
		}

		s.currX += (s.x - s.currX) / s.options.SoftX
		// s.currY += (s.y - s.currY) / s.options.SoftY
		s.currY += s.options.SoftY * 5 / (s.y - s.currY)

		if math.Abs(s.x-s.currX) < 0 {
			s.currX = s.x
		}

		if math.Abs(s.y-s.currY) < 0 {
			s.currY = s.y
		}

		return
	}

	s.currX += (s.x - s.currX) / s.options.SoftX
	s.currY += (s.y - s.currY) / s.options.SoftY
	if math.Abs(s.x-s.currX) < 1 {
		s.currX = s.x
	}

	if math.Abs(s.y-s.currY) < 1 {
		s.currY = s.y
	}
}

func (s *sprite) stepToNextImage() {

	if !s.options.Animate || s.closed {
		return
	}

	if s.options.AnimateOnMove && !s.IsMoving() {
		return
	}

	if s.animationCount == 0 {
		s.animationCount = s.options.AnimationSpeed
		if len(s.imgs[s.collectionIdx])-1 == s.imgIndex {
			s.imgIndex = 0
			return
		}

		s.imgIndex++
		return
	}

	s.animationCount--
}
