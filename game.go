package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/widget"
	"image/color"
)

type Game struct {
	ui widget.Ui
}

func NewGame() *Game {
	ui := widget.NewUi()
	return &Game{
		ui: ui,
	}
}

func (g *Game) Update() error {
	g.ui.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.ui.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
