package scene

import (
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"

	"github.com/m110/making-games-in-go/assets"
)

type HierarchyObject struct {
	Position *Vector
	Sprite   *ebiten.Image
	Children []*HierarchyObject
}

type Hierarchy struct {
	objects []*HierarchyObject
}

func newHealtyGopher(pos *Vector) *HierarchyObject {
	bar := &HierarchyObject{
		Position: &Vector{X: 0, Y: -50},
		Sprite:   newHealthBar(150, 30, 0.50+rand.Float64()*0.25),
	}

	return &HierarchyObject{
		Position: pos,
		Sprite:   assets.SmallGopherSprite,
		Children: []*HierarchyObject{bar},
	}
}

func NewHierarchy() Scene {
	return &Hierarchy{
		objects: []*HierarchyObject{
			newHealtyGopher(&Vector{X: 0, Y: 150}),
			newHealtyGopher(&Vector{X: 250, Y: 250}),
			newHealtyGopher(&Vector{X: 550, Y: 190}),
		},
	}
}

func (h *Hierarchy) Update() error {
	for _, obj := range h.objects {
		obj.Position.X += 1
	}
	return nil
}

func (h *Hierarchy) Draw(screen *ebiten.Image) {
	for _, obj := range h.objects {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(obj.Position.X, obj.Position.Y)
		screen.DrawImage(obj.Sprite, op)

		for _, child := range obj.Children {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(
				obj.Position.X+child.Position.X,
				obj.Position.Y+child.Position.Y,
			)
			screen.DrawImage(child.Sprite, op)
		}
	}
}

func newHealthBar(width int, height int, percent float64) *ebiten.Image {
	bar := ebiten.NewImage(width, height)
	bar.Fill(colornames.Darkred)

	barWidth := int(float64(width) * percent)

	vector.DrawFilledRect(bar, 0, 0, float32(barWidth), float32(height), colornames.Red, true)

	op := &text.DrawOptions{}
	op.GeoM.Translate(10, -5)
	text.Draw(bar, fmt.Sprintf("%v%%", int(percent*100)), assets.SmallFont, op)

	return bar
}
