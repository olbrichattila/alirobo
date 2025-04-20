// Package imageManager is responsible for image manipulation
package imageManager

import (
	"image"
	_ "image/png"

	"math"
	"net/http"

	"github.com/hajimehoshi/ebiten/v2"
)

func Load(fileName string) *ebiten.Image {
	// img, _, err := ebitenutil.NewImageFromFile(fileName)
	img, err := loadImageFromURL(fileName)
	if err != nil {
		panic(err)
	}

	return img
}

func LoadWithSize(fileName string, targetWidth, targetHeight int) *ebiten.Image {
	img, _, _, err := RescaleImageToFit(fileName, targetWidth, targetHeight)
	if err != nil {
		panic(err)
	}

	return img
}

func RescaleImageToFitFloat(imageName string, targetWidth, targetHeight float64) (*ebiten.Image, int, int, error) {
	return RescaleImageToFit(imageName, int(targetWidth), int(targetHeight))
}

func RescaleImageToFit(fileName string, targetWidth, targetHeight int) (*ebiten.Image, int, int, error) {
	img, err := loadImageFromURL(fileName)
	// img, _, err := ebitenutil.NewImageFromFile(imageName)
	if err != nil {
		return nil, 0, 0, err
	}

	bounds := img.Bounds()
	origWidth := bounds.Dx()
	origHeight := bounds.Dy()

	widthRatio := float64(targetWidth) / float64(origWidth)
	heightRatio := float64(targetHeight) / float64(origHeight)

	scale := math.Min(widthRatio, heightRatio)
	newWidth := int(float64(origWidth) * scale)
	newHeight := int(float64(origHeight) * scale)

	rescaled := ebiten.NewImage(newWidth, newHeight)

	geom := ebiten.GeoM{}
	geom.Scale(scale, scale)

	rescaled.DrawImage(img, &ebiten.DrawImageOptions{
		GeoM: geom,
	})

	return rescaled, newWidth, newHeight, nil
}

func ReplicateImageVertically(src *ebiten.Image, times int) *ebiten.Image {
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	dst := ebiten.NewImage(w, h*times)

	for i := 0; i < times; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, float64(i*h))
		dst.DrawImage(src, op)
	}

	return dst
}

func CropVertical(src *ebiten.Image, w, h float64) *ebiten.Image {
	subImg := src.SubImage(image.Rect(0, 0, int(w), int(h)))

	return subImg.(*ebiten.Image)
}

func FlipImageVertically(src *ebiten.Image) *ebiten.Image {
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	dst := ebiten.NewImage(w, h)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, -1) // Flip vertically
	op.GeoM.Translate(0, float64(h))

	dst.DrawImage(src, op)

	return dst
}

func FlipImageHorizontally(src *ebiten.Image) *ebiten.Image {
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	dst := ebiten.NewImage(w, h)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(float64(w), 0)

	dst.DrawImage(src, op)

	return dst
}

func loadImageFromURL(url string) (*ebiten.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// bodyBytes, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// bodyString := string(bodyBytes)
	// fmt.Println(bodyString)

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err

	}

	return ebiten.NewImageFromImage(img), nil
}
