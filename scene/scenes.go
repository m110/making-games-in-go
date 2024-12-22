package scene

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/m110/making-games-in-go/assets"
)

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type Static struct{}

func NewStatic() Scene {
	return &Static{}
}

func (s *Static) Update() error {
	return nil
}

func (s *Static) Draw(screen *ebiten.Image) {
	if Debug {
		screen.DrawImage(assets.GopherDebug, nil)
	}

	screen.DrawImage(assets.GopherSprite, nil)
}

type Moved struct{}

func NewMoved() Scene {
	return &Moved{}
}

func (m *Moved) Update() error {
	return nil
}

func (m *Moved) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(300, 00)

	if Debug {
		screen.DrawImage(assets.GopherDebug, op)
	}

	screen.DrawImage(assets.GopherSprite, op)
}

type ConstantVelocity struct {
	playerPosition *Vector
}

func NewConstantVelocity() Scene {
	return &ConstantVelocity{
		playerPosition: &Vector{},
	}
}

func (c *ConstantVelocity) Update() error {
	c.playerPosition.X += 1
	return nil
}

func (c *ConstantVelocity) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.playerPosition.X, c.playerPosition.Y)

	if Debug {
		screen.DrawImage(assets.GopherDebug, op)
	}

	screen.DrawImage(assets.GopherSprite, op)
}
