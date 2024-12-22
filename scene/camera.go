package scene

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/m110/making-games-in-go/assets"
)

type Camera struct {
	cameraPosition *Vector
	offscreen      *ebiten.Image
}

func NewCamera() Scene {
	return &Camera{
		cameraPosition: &Vector{},
		offscreen:      ebiten.NewImage(2500, 540),
	}
}

func (c *Camera) Update() error {
	c.cameraPosition.X += 1

	return nil
}

func (c *Camera) Draw(screen *ebiten.Image) {
	c.offscreen.Clear()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	c.offscreen.DrawImage(assets.BgSprite, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-c.cameraPosition.X, 0)
	screen.DrawImage(c.offscreen, op)
}
