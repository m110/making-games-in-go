package scene

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/m110/making-games-in-go/assets"
)

type TicksCounting struct {
	position *Vector

	ticks      int
	movingLeft bool
}

func NewTicksCounting() Scene {
	return &TicksCounting{
		position: &Vector{X: 150, Y: 250},
	}
}

func (f *TicksCounting) Update() error {
	f.ticks++

	if f.ticks == 120 {
		f.movingLeft = !f.movingLeft
		f.ticks = 0
	}

	if f.movingLeft {
		f.position.X -= 2
	} else {
		f.position.X += 2
	}

	return nil
}

func (f *TicksCounting) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(f.position.X, f.position.Y)

	if Debug {
		screen.DrawImage(assets.GopherDebug, op)
	}

	screen.DrawImage(assets.GopherSprite, op)

	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(25, -25)

	text.Draw(
		screen,
		fmt.Sprintf("Ticks: %d", f.ticks),
		assets.NormalFont,
		textOp,
	)

	textOp.GeoM.Translate(0, 128)

	direction := "Right"
	if f.movingLeft {
		direction = "Left"
	}

	text.Draw(
		screen,
		fmt.Sprintf("Direction: %v", direction),
		assets.NormalFont,
		textOp,
	)
}

type WithTimer struct {
	position *Vector

	timer      *Timer
	movingLeft bool
}

func NewWithTimer() Scene {
	return &WithTimer{
		position: &Vector{X: 150, Y: 250},
		timer:    NewTimer(2 * time.Second),
	}
}

func (w *WithTimer) Update() error {
	w.timer.Update()

	if w.timer.IsDone() {
		w.movingLeft = !w.movingLeft
		w.timer.Reset()
	}

	if w.movingLeft {
		w.position.X -= 2
	} else {
		w.position.X += 2
	}

	return nil
}

func (w *WithTimer) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(float32(w.timer.PercentDone()))
	op.GeoM.Translate(w.position.X, w.position.Y)

	if Debug {
		screen.DrawImage(assets.GopherDebug, op)
	}

	screen.DrawImage(assets.GopherSprite, op)

	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(25, -25)

	text.Draw(
		screen,
		fmt.Sprintf("Timer: %v%%", int(w.timer.PercentDone()*100)),
		assets.NormalFont,
		textOp,
	)

	textOp.GeoM.Translate(0, 128)

	direction := "Right"
	if w.movingLeft {
		direction = "Left"
	}

	text.Draw(
		screen,
		fmt.Sprintf("Direction: %v", direction),
		assets.NormalFont,
		textOp,
	)
}
