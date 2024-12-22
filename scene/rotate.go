package scene

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/m110/making-games-in-go/assets"
)

type NaiveRotate struct {
	startTimer *Timer
	timer      *Timer
}

func NewNaiveRotate() Scene {
	return &NaiveRotate{
		startTimer: NewTimer(1 * time.Second),
		timer:      NewTimer(2 * time.Second),
	}
}

func (r *NaiveRotate) Update() error {
	r.startTimer.Update()
	if r.startTimer.IsDone() {
		r.timer.Update()
	}
	return nil
}

func (r *NaiveRotate) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(toRadians(90 * r.timer.PercentDone()))

	if Debug {
		screen.DrawImage(assets.GopherDebug, op)
	}

	screen.DrawImage(assets.GopherSprite, op)
}

type Rotate struct {
	startTimer *Timer
	timer      *Timer
}

func NewRotate() Scene {
	return &Rotate{
		startTimer: NewTimer(time.Second),
		timer:      NewTimer(time.Second),
	}
}

func (r *Rotate) Update() error {
	r.startTimer.Update()
	if r.startTimer.IsDone() {
		r.timer.Update()
	}
	return nil
}

func (r *Rotate) Draw(screen *ebiten.Image) {
	bounds := assets.GopherSprite.Bounds()
	halfWidth := float64(bounds.Dx()) / 2
	halfHeight := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfWidth, -halfHeight)
	op.GeoM.Rotate(toRadians(90 * r.timer.PercentDone()))
	op.GeoM.Translate(halfWidth, halfHeight)

	if Debug {
		screen.DrawImage(assets.GopherDebug, op)
	}

	screen.DrawImage(assets.GopherSprite, op)
}

type RotateStepByStep struct {
	startTimer *Timer
	timer      *Timer
	stage      int
}

func NewRotateStepByStep() Scene {
	return &RotateStepByStep{
		startTimer: NewTimer(time.Second),
		timer:      NewTimer(time.Second),
		stage:      0,
	}
}

func (r *RotateStepByStep) Update() error {
	r.startTimer.Update()
	if r.startTimer.IsDone() {
		r.timer.Update()
		if r.timer.IsDone() {
			r.stage++
			r.timer.Reset()
		}
	}

	return nil
}

func (r *RotateStepByStep) Draw(screen *ebiten.Image) {
	bounds := assets.GopherSprite.Bounds()
	halfWidth := float64(bounds.Dx()) / 2
	halfHeight := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}

	pct := r.timer.PercentDone()
	switch r.stage {
	case 0:
		op.GeoM.Translate(-halfWidth*pct, -halfHeight*pct)
	case 1:
		op.GeoM.Translate(-halfWidth, -halfHeight)
		op.GeoM.Rotate(toRadians(90 * pct))
	case 2:
		op.GeoM.Translate(-halfWidth, -halfHeight)
		op.GeoM.Rotate(toRadians(90))
		op.GeoM.Translate(halfWidth*pct, halfHeight*pct)
	default:
		op.GeoM.Translate(-halfWidth, -halfHeight)
		op.GeoM.Rotate(toRadians(90))
		op.GeoM.Translate(halfWidth, halfHeight)
	}

	if Debug {
		screen.DrawImage(assets.GopherDebug, op)
	}

	screen.DrawImage(assets.GopherSprite, op)
}
