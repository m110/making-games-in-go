package scene

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/m110/making-games-in-go/assets"
)

type Object struct {
	Position *Vector
	Scale    *Vector
	Sprite   *ebiten.Image
}

func newCrab(pos Vector) *Object {
	return &Object{
		Position: &Vector{
			X: pos.X,
			Y: pos.Y,
		},
		Sprite: assets.CrabSprite,
	}
}

type Objects struct {
	objects []*Object
}

func NewObjects() Scene {
	objects := []*Object{}

	for range 500 {
		objects = append(objects, newCrab(Vector{
			X: float64(rand.Intn(1000)),
			Y: float64(-50 + rand.Intn(600)),
		}))
	}

	return &Objects{
		objects: objects,
	}
}

func (c *Objects) Update() error {
	for _, obj := range c.objects {
		obj.Position.X -= 2
	}
	return nil
}

func (c *Objects) Draw(screen *ebiten.Image) {
	for _, obj := range c.objects {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(obj.Position.X, obj.Position.Y)
		screen.DrawImage(obj.Sprite, op)
	}
}
