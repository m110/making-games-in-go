package main

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/colornames"

	"github.com/m110/making-games-in-go/assets"
	"github.com/m110/making-games-in-go/scene"
)

const (
	scenesPerColumn = 15

	windowWidth  = 960
	windowHeight = 540
)

func main() {
	err := assets.LoadAssets()
	if err != nil {
		panic(err)
	}

	scenes := []func() scene.Scene{
		// Drawing
		scene.NewStatic,
		scene.NewMoved,
		scene.NewScale,
		scene.NewFirstScaleThenMove,
		scene.NewFirstMoveThenScale,
		scene.NewNaiveRotate,
		scene.NewRotate,
		scene.NewRotateStepByStep,
		scene.NewFlip,
		scene.NewRed,
		scene.NewAlpha,
		scene.NewParty,

		// Logic updates
		scene.NewConstantVelocity,
		scene.NewTicksCounting,
		scene.NewWithTimer,
		scene.NewObjects,

		// Beyond essentials
		scene.NewLayers,
		scene.NewCamera,
		scene.NewStaticAnimation,
		scene.NewAnimation,
		scene.NewHierarchy,
		scene.NewCrabsAttack,
	}

	game := &Game{
		scenes:            scenes,
		currentSceneIndex: 0,
		currentScene:      scenes[0](),
		showMenu:          true,
		menuBackground:    newMenuBackground(),
	}

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(windowWidth, windowHeight)

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

type Game struct {
	scenes            []func() scene.Scene
	currentScene      scene.Scene
	currentSceneIndex int

	showMenu       bool
	menuBackground *ebiten.Image
}

func newMenuBackground() *ebiten.Image {
	bg := ebiten.NewImage(windowWidth, windowHeight)
	bgColor := color.RGBA{
		R: 30,
		G: 30,
		B: 30,
		A: 220,
	}
	bg.Fill(bgColor)

	linesMargin := 30.0

	op := &text.DrawOptions{}
	op.GeoM.Translate(550, 50)
	text.Draw(
		bg,
		"Controls:",
		assets.SmallFont,
		op,
	)

	op.GeoM.Translate(10, 40)
	text.Draw(
		bg,
		"- Arrow keys or WASD to change the level",
		assets.TinyFont,
		op,
	)

	op.GeoM.Translate(0, linesMargin)
	text.Draw(
		bg,
		"- Esc to show/hide this screen",
		assets.TinyFont,
		op,
	)

	op.GeoM.Translate(0, linesMargin)
	text.Draw(
		bg,
		"- Slash (/) to toggle debug mode",
		assets.TinyFont,
		op,
	)

	op.GeoM.Translate(0, linesMargin)
	text.Draw(
		bg,
		"- R to restart current scene",
		assets.TinyFont,
		op,
	)

	return bg
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) ||
		inpututil.IsKeyJustPressed(ebiten.KeyUp) ||
		inpututil.IsKeyJustPressed(ebiten.KeyW) ||
		inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if g.currentSceneIndex > 0 {
			g.currentSceneIndex--
		} else {
			g.currentSceneIndex = len(g.scenes) - 1
		}
		g.currentScene = g.scenes[g.currentSceneIndex]()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) ||
		inpututil.IsKeyJustPressed(ebiten.KeyDown) ||
		inpututil.IsKeyJustPressed(ebiten.KeyS) ||
		inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if g.currentSceneIndex < len(g.scenes)-1 {
			g.currentSceneIndex++
		} else {
			g.currentSceneIndex = 0
		}
		g.currentScene = g.scenes[g.currentSceneIndex]()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.currentScene = g.scenes[g.currentSceneIndex]()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.showMenu = !g.showMenu
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySlash) {
		scene.Debug = !scene.Debug
	}

	err := g.currentScene.Update()
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentScene.Draw(screen)

	if g.showMenu {
		screen.DrawImage(g.menuBackground, nil)

		op := &text.DrawOptions{}
		op.GeoM.Translate(50, 50)
		text.Draw(
			screen,
			"Examples:",
			assets.SmallFont,
			op,
		)

		for i, s := range g.scenes {
			name := fmt.Sprintf("%T", s())
			name = strings.TrimPrefix(strings.Split(name, ".")[1], "*")
			txt := fmt.Sprintf("%v. %v", i+1, name)

			op := &text.DrawOptions{}
			if i == g.currentSceneIndex {
				op.ColorScale.ScaleWithColor(colornames.Lime)
			}

			x := 50 + float64(i/scenesPerColumn)*200
			y := 100 + float64(i%scenesPerColumn)*20

			op.GeoM.Translate(x, y)
			text.Draw(
				screen,
				txt,
				assets.TinyFont,
				op,
			)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}
