package scene

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/m110/making-games-in-go/assets"
)

type TinyGopher struct {
	Object
	Rotation float64
	Color    color.RGBA
}

type Party struct {
	objects []*TinyGopher
}

func NewParty() Scene {
	var objects []*TinyGopher

	for y := 0; y < 4; y++ {
		for x := 0; x < 20; x++ {
			pos := Vector{
				X: float64(x) * (75 + float64(rand.Intn(10))),
				Y: 90 + float64(y)*(70+float64(rand.Intn(10))),
			}
			objects = append(objects, newTinyGopher(pos))
		}
	}

	return &Party{
		objects: objects,
	}
}

func newTinyGopher(pos Vector) *TinyGopher {
	return &TinyGopher{
		Object: Object{
			Position: &Vector{
				X: pos.X,
				Y: pos.Y,
			},
			Scale: &Vector{
				X: 0.5,
				Y: 0.5,
			},
			Sprite: assets.SmallGopherSprite,
		},
		Rotation: -10 + float64(rand.Intn(20)),
		Color: color.RGBA{
			R: uint8(55 + rand.Intn(200)),
			G: uint8(55 + rand.Intn(200)),
			B: uint8(55 + rand.Intn(200)),
			A: 255,
		},
	}
}

func (s *Party) Update() error {
	return nil
}

func (s *Party) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.25, 0.25)
	screen.DrawImage(assets.BgSprite, op)

	for _, obj := range s.objects {
		op := &ebiten.DrawImageOptions{}
		op.Filter = ebiten.FilterLinear

		bounds := obj.Sprite.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2

		op.GeoM.Translate(-halfW, -halfH)
		op.GeoM.Rotate(toRadians(obj.Rotation))
		op.GeoM.Translate(halfW, halfH)

		op.GeoM.Scale(obj.Scale.X, obj.Scale.Y)
		op.GeoM.Translate(obj.Position.X, obj.Position.Y)

		op.ColorScale.ScaleWithColor(obj.Color)

		screen.DrawImage(obj.Sprite, op)
	}
}
