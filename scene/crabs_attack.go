package scene

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"

	"github.com/m110/making-games-in-go/assets"
)

const (
	floorY    = 170
	jumpForce = -17
	gravity   = 0.8
)

var spawnPosition = Vector{X: 960, Y: 300}

var crabCollider = Rect{
	Position: Vector{X: 30, Y: 10},
	Size:     Vector{X: 130, Y: 100},
}

type Crab struct {
	Object
	IsStopped bool
}

type CrabsAttack struct {
	playerPosition *Vector
	isJumping      bool

	crabs       []*Crab
	spawnTimer  *Timer
	velocityY   float64
	shadowScale float64
	gameOver    bool
}

func NewCrabsAttack() Scene {
	timer := NewTimer(5 * time.Second)
	timer.Finish()

	return &CrabsAttack{
		playerPosition: &Vector{X: 150, Y: floorY},
		spawnTimer:     timer,
		shadowScale:    1,
	}
}

func (c *CrabsAttack) Update() error {
	if c.gameOver {
		return nil
	}

	if c.playerPosition.Y >= floorY && ebiten.IsKeyPressed(ebiten.KeySpace) {
		c.velocityY = jumpForce
		c.isJumping = true
	}

	c.velocityY += gravity
	c.playerPosition.Y += c.velocityY

	if c.playerPosition.Y > floorY {
		c.playerPosition.Y = floorY
		c.velocityY = 0
		c.isJumping = false
	}

	heightFromGround := floorY - c.playerPosition.Y

	// Shadow gets smaller as player goes higher
	// 1.0 is original size, minimum size is 0.3
	c.shadowScale = math.Max(1.0-(heightFromGround/200), 0.3)

	c.spawnTimer.Update()
	if c.spawnTimer.IsDone() {
		c.crabs = append(c.crabs, newAttackingCrab(spawnPosition))
		c.spawnTimer.Reset()
	}

	for _, crab := range c.crabs {
		if crab.IsStopped {
			if crab.Scale.Y > 0.1 {
				crab.Scale.X += 0.02
				crab.Scale.Y -= 0.05
			}
		} else {
			crab.Position.X -= 4
		}
	}

	bounds := assets.SmallGopherSprite.Bounds()
	playerRect := Rect{
		Position: *c.playerPosition,
		Size:     Vector{X: float64(bounds.Dx()), Y: float64(bounds.Dy())},
	}

	for _, crab := range c.crabs {
		if crab.IsStopped {
			continue
		}

		crabRect := Rect{
			Position: Vector{
				X: crab.Position.X + crabCollider.Position.X,
				Y: crab.Position.Y + crabCollider.Position.Y,
			},
			Size: Vector{X: crabCollider.Size.X, Y: crabCollider.Size.Y},
		}

		if playerRect.Intersects(crabRect) {
			if playerRect.Position.Y+playerRect.Size.Y < crabRect.Position.Y+crabCollider.Position.Y {
				crab.IsStopped = true
			} else {
				c.gameOver = true
			}
		}
	}

	return nil
}

func (c *CrabsAttack) Draw(screen *ebiten.Image) {
	if c.gameOver {
		op := &text.DrawOptions{}
		op.GeoM.Translate(250, 180)
		text.Draw(screen, "Game Over", assets.NormalFont, op)
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.25, 0.25)
	screen.DrawImage(assets.BgSprite, op)

	crabBounds := assets.CrabSprite.Bounds()

	for _, crab := range c.crabs {
		if !crab.IsStopped {
			op = &ebiten.DrawImageOptions{}
			op.GeoM.Scale(0.7, 0.5)
			op.GeoM.Translate(crab.Position.X+5, crab.Position.Y+75)
			op.ColorScale.ScaleAlpha(0.6)
			screen.DrawImage(assets.ShadowSprite, op)
		}

		op = &ebiten.DrawImageOptions{}
		op.GeoM.Scale(crab.Scale.X, crab.Scale.Y)
		op.GeoM.Translate(crab.Position.X, crab.Position.Y)
		if crab.IsStopped {
			op.GeoM.Translate(
				float64(crabBounds.Dx())*(1-crab.Scale.X),
				float64(crabBounds.Dy())*(1-crab.Scale.Y),
			)
		}

		if Debug {
			screen.DrawImage(assets.CrabDebug, op)
		}
		screen.DrawImage(crab.Sprite, op)

		if !crab.IsStopped && Debug {
			vector.StrokeRect(
				screen,
				float32(crab.Position.X+crabCollider.Position.X),
				float32(crab.Position.Y+crabCollider.Position.Y),
				float32(crabCollider.Size.X),
				float32(crabCollider.Size.Y),
				5,
				colornames.Limegreen,
				false,
			)
		}
	}

	bounds := assets.ShadowSprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Scale(0.5*c.shadowScale, 0.5*c.shadowScale)
	op.GeoM.Translate(halfW/2, halfH/2)
	op.GeoM.Translate(c.playerPosition.X, 365)
	op.ColorScale.ScaleAlpha(0.3 + 0.3*float32(c.shadowScale))
	screen.DrawImage(assets.ShadowSprite, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.playerPosition.X, c.playerPosition.Y)

	if Debug {
		screen.DrawImage(assets.SmallGopherDebug, op)
	}
	screen.DrawImage(assets.SmallGopherSprite, op)
	if Debug {
		bounds := assets.SmallGopherSprite.Bounds()
		vector.StrokeRect(
			screen,
			float32(c.playerPosition.X),
			float32(c.playerPosition.Y),
			float32(bounds.Dx()),
			float32(bounds.Dy()),
			5,
			colornames.Limegreen,
			false,
		)
	}
}

func newAttackingCrab(pos Vector) *Crab {
	return &Crab{
		Object: Object{
			Position: &Vector{
				X: pos.X,
				Y: pos.Y,
			},
			Scale: &Vector{
				X: 1,
				Y: 1,
			},
			Sprite: assets.CrabSprite,
		},
	}
}
