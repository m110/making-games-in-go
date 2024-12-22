package scene

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/m110/making-games-in-go/assets"
)

type Scale struct{}

func NewScale() Scene {
	return &Scale{}
}

func (s *Scale) Update() error {
	return nil
}

func (s *Scale) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2.5, 2.5)
	screen.DrawImage(assets.GopherSprite, op)
}

type Flip struct{}

func NewFlip() Scene {
	return &Flip{}
}

func (f *Flip) Update() error {
	return nil
}

func (f *Flip) Draw(screen *ebiten.Image) {
	bounds := assets.GopherSprite.Bounds()
	halfW := bounds.Dx() / 2
	halfH := bounds.Dy() / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(halfW), -float64(halfH))
	op.GeoM.Scale(1, -1)
	op.GeoM.Translate(float64(halfW), float64(halfH))

	screen.DrawImage(assets.GopherSprite, op)
}
