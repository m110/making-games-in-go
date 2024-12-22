package assets

import (
	"bytes"
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/colornames"
	"golang.org/x/text/language"
)

//go:embed fonts/Not-Jam-UI-Condensed-19.ttf
var fontBytes []byte

//go:embed zombie.png
var zombieBytes []byte

//go:embed gopher.png
var gopherBytes []byte

//go:embed crab.png
var crabBytes []byte

//go:embed bg.png
var bgBytes []byte

//go:embed shadow.png
var shadowBytes []byte

//go:embed walk/*.png
var walkBytes embed.FS

var ZombieSprite *ebiten.Image
var ZombieWalkFrames []*ebiten.Image

var GopherSprite, SmallGopherSprite, CrabSprite, BgSprite, ShadowSprite *ebiten.Image
var GopherDebug, SmallGopherDebug, CrabDebug *ebiten.Image

var NormalFont, SmallFont, TinyFont *text.GoTextFace

func LoadAssets() error {
	gopherImage, _, err := image.Decode(bytes.NewReader(gopherBytes))
	if err != nil {
		return err
	}
	GopherSprite = ebiten.NewImageFromImage(gopherImage)

	GopherDebug = ebiten.NewImageFromImage(GopherSprite)
	GopherDebug.Fill(colornames.Magenta)

	SmallGopherSprite = ebiten.NewImage(GopherSprite.Bounds().Dx()/2, GopherSprite.Bounds().Dy()/2)
	SmallGopherDebug = ebiten.NewImage(SmallGopherSprite.Bounds().Dx(), SmallGopherSprite.Bounds().Dy())
	SmallGopherDebug.Fill(colornames.Magenta)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	op.Filter = ebiten.FilterLinear
	SmallGopherSprite.DrawImage(GopherSprite, op)

	crabImage, _, err := image.Decode(bytes.NewReader(crabBytes))
	if err != nil {
		return err
	}

	CrabSprite = ebiten.NewImageFromImage(crabImage)
	CrabDebug = ebiten.NewImageFromImage(CrabSprite)
	CrabDebug.Fill(colornames.Magenta)

	zombieImage, _, err := image.Decode(bytes.NewReader(zombieBytes))
	if err != nil {
		return err
	}
	ZombieSprite = ebiten.NewImageFromImage(zombieImage)

	bgImage, _, err := image.Decode(bytes.NewReader(bgBytes))
	if err != nil {
		return err
	}
	BgSprite = ebiten.NewImageFromImage(bgImage)

	shadowImage, _, err := image.Decode(bytes.NewReader(shadowBytes))
	if err != nil {
		return err
	}
	ShadowSprite = ebiten.NewImageFromImage(shadowImage)

	walkFiles, err := walkBytes.ReadDir("walk")
	if err != nil {
		return err
	}

	for _, file := range walkFiles {
		f, err := walkBytes.ReadFile("walk/" + file.Name())
		if err != nil {
			return err
		}

		img, _, err := image.Decode(bytes.NewReader(f))
		if err != nil {
			return err
		}

		ZombieWalkFrames = append(ZombieWalkFrames, ebiten.NewImageFromImage(img))
	}

	NormalFont, err = loadFont(fontBytes, 128)
	if err != nil {
		return err
	}

	SmallFont, err = loadFont(fontBytes, 36)
	if err != nil {
		return err
	}

	TinyFont, err = loadFont(fontBytes, 24)
	if err != nil {
		return err
	}

	return nil
}

func loadFont(data []byte, size int) (*text.GoTextFace, error) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	return &text.GoTextFace{
		Source:    s,
		Direction: text.DirectionLeftToRight,
		Size:      float64(size),
		Language:  language.English,
	}, nil
}
