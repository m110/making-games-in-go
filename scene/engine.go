package scene

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func toRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180.0)
}

type Vector struct {
	X float64
	Y float64
}

type Rect struct {
	Position Vector
	Size     Vector
}

func (r Rect) Intersects(other Rect) bool {
	return r.Position.X < other.Position.X+other.Size.X &&
		r.Position.X+r.Size.X > other.Position.X &&
		r.Position.Y < other.Position.Y+other.Size.Y &&
		r.Position.Y+r.Size.Y > other.Position.Y
}

type Timer struct {
	targetTicks  int
	currentTicks int
}

func NewTimer(duration time.Duration) *Timer {
	return &Timer{
		targetTicks: int(float64(duration) / (float64(time.Second) / float64(ebiten.TPS()))),
	}
}

func (t *Timer) Update() {
	t.currentTicks++
}

func (t *Timer) IsDone() bool {
	return t.currentTicks >= t.targetTicks
}

func (t *Timer) Reset() {
	t.currentTicks = 0
}

func (t *Timer) Finish() {
	t.currentTicks = t.targetTicks
}

func (t *Timer) PercentDone() float64 {
	pct := float64(t.currentTicks) / float64(t.targetTicks)
	if pct > 1.0 {
		return 1.0
	}

	return pct
}
