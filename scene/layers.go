package scene

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/m110/making-games-in-go/assets"
)

const (
	LayerBackground = iota
	LayerObjects
	LayerForeground
)

type Layers struct {
	objects [][]Object
}

func NewLayers() Scene {
	return &Layers{
		objects: [][]Object{
			{
				Object{
					Position: &Vector{X: 0, Y: 0},
					Scale:    &Vector{X: 0.25, Y: 0.25},
					Sprite:   assets.BgSprite,
				},
			},
			{
				Object{
					Position: &Vector{X: 150, Y: 150},
					Sprite:   assets.GopherSprite,
				},
			},
			{
				Object{
					Position: &Vector{X: 100, Y: 150},
					Scale:    &Vector{X: 2, Y: 2},
					Sprite:   assets.CrabSprite,
				},
			},
		},
	}
}

func (l *Layers) Update() error {
	for layer := range l.objects {
		for _, obj := range l.objects[layer] {
			obj.Position.X += float64(layer)
		}
	}

	return nil
}

func (l *Layers) Draw(screen *ebiten.Image) {
	for layer := range l.objects {
		for _, obj := range l.objects[layer] {
			op := &ebiten.DrawImageOptions{}
			if obj.Scale != nil {
				op.GeoM.Scale(obj.Scale.X, obj.Scale.Y)
			}

			op.GeoM.Translate(obj.Position.X, obj.Position.Y)

			screen.DrawImage(obj.Sprite, op)
		}
	}
}
