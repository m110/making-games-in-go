package scene

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/m110/making-games-in-go/assets"
)

type FirstMoveThenScale struct {
	timer *Timer
	moved bool
}

func NewFirstMoveThenScale() Scene {
	return &FirstMoveThenScale{
		timer: NewTimer(time.Second),
		moved: false,
	}
}

func (f *FirstMoveThenScale) Update() error {
	f.timer.Update()
	if f.timer.IsDone() {
		if !f.moved {
			f.moved = true
			f.timer.Reset()
		}
	}

	return nil
}

func (f *FirstMoveThenScale) Draw(screen *ebiten.Image) {
	pct := f.timer.PercentDone()

	op := &ebiten.DrawImageOptions{}

	if f.moved {
		op.GeoM.Translate(350, 0)
		op.GeoM.Scale(1+pct, 1+pct)
	} else {
		op.GeoM.Translate(350*pct, 0)
	}

	if Debug {
		screen.DrawImage(assets.GopherDebug, op)
	}

	screen.DrawImage(assets.GopherSprite, op)
}

type FirstScaleThenMove struct {
	timer  *Timer
	scaled bool
}

func NewFirstScaleThenMove() Scene {
	return &FirstScaleThenMove{
		timer:  NewTimer(time.Second),
		scaled: false,
	}
}

func (f *FirstScaleThenMove) Update() error {
	f.timer.Update()
	if f.timer.IsDone() {
		if !f.scaled {
			f.scaled = true
			f.timer.Reset()
		}
	}

	return nil
}

func (f *FirstScaleThenMove) Draw(screen *ebiten.Image) {
	pct := f.timer.PercentDone()

	op := &ebiten.DrawImageOptions{}

	if f.scaled {
		op.GeoM.Scale(2, 2)
		op.GeoM.Translate(350*pct, 0)
	} else {
		op.GeoM.Scale(1+pct, 1+pct)
	}

	if Debug {
		screen.DrawImage(assets.GopherDebug, op)
	}

	screen.DrawImage(assets.GopherSprite, op)
}
