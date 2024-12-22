package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"

	"github.com/m110/making-games-in-go/assets"
)

type RGB struct{}

func NewRed() Scene {
	return &RGB{}
}

func (r *RGB) Update() error {
	return nil
}

func (r *RGB) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleWithColor(colornames.Red)
	screen.DrawImage(assets.GopherSprite, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(320, 0)
	op.ColorScale.ScaleWithColor(colornames.Green)
	screen.DrawImage(assets.GopherSprite, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(640, 0)
	op.ColorScale.ScaleWithColor(colornames.Blue)
	screen.DrawImage(assets.GopherSprite, op)
}

type Alpha struct{}

func NewAlpha() Scene {
	return &Alpha{}
}

func (a *Alpha) Update() error {
	return nil
}

func (a *Alpha) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(0.75)
	screen.DrawImage(assets.GopherSprite, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(320, 0)
	op.ColorScale.ScaleAlpha(0.5)
	screen.DrawImage(assets.GopherSprite, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(640, 0)
	op.ColorScale.ScaleAlpha(0.25)
	screen.DrawImage(assets.GopherSprite, op)
}
