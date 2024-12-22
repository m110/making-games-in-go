package scene

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/m110/making-games-in-go/assets"
)

const (
	zombieWalkSpeed = 5
)

type StaticAnimation struct {
	currentFrameIndex int
	animationTimer    *Timer
}

func NewStaticAnimation() Scene {
	return &StaticAnimation{
		currentFrameIndex: 0,
		animationTimer:    NewTimer(time.Millisecond * 100),
	}
}

func (s *StaticAnimation) Update() error {
	s.animationTimer.Update()
	if s.animationTimer.IsDone() {
		s.animationTimer.Reset()

		s.currentFrameIndex++
		if s.currentFrameIndex >= len(assets.ZombieWalkFrames) {
			s.currentFrameIndex = 0
		}
	}

	return nil
}

func (s *StaticAnimation) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear
	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(250, 0)

	img := assets.ZombieWalkFrames[s.currentFrameIndex]
	screen.DrawImage(img, op)
}

type Animation struct {
	position          *Vector
	currentFrameIndex int
	isWalking         bool
	flip              bool
	animationTimer    *Timer
}

func NewAnimation() Scene {
	return &Animation{
		position: &Vector{
			X: 50,
			Y: 20,
		},
		currentFrameIndex: 0,
		isWalking:         false,
		animationTimer:    NewTimer(time.Millisecond * 100),
	}
}

func (a *Animation) Update() error {
	if a.isWalking {
		a.animationTimer.Update()
		if a.animationTimer.IsDone() {
			a.animationTimer.Reset()

			a.currentFrameIndex++
			if a.currentFrameIndex >= len(assets.ZombieWalkFrames) {
				a.currentFrameIndex = 0
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyL) {
		a.isWalking = true
		a.flip = false
		a.position.X += zombieWalkSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyH) {
		a.isWalking = true
		a.flip = true
		a.position.X -= zombieWalkSpeed
	} else {
		a.currentFrameIndex = 0
		a.isWalking = false
	}

	return nil
}

func (a *Animation) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.25, 0.25)
	screen.DrawImage(assets.BgSprite, op)

	img := assets.ZombieSprite
	if a.isWalking {
		img = assets.ZombieWalkFrames[a.currentFrameIndex]
	}

	op = &ebiten.DrawImageOptions{}
	if a.flip {
		bounds := img.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2

		op.GeoM.Translate(-halfW, -halfH)
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(halfW, halfH)
	}
	op.GeoM.Scale(1.5, 1.5)
	op.GeoM.Translate(a.position.X, a.position.Y)

	screen.DrawImage(img, op)
}
