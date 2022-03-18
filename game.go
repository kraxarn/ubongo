package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/assets"
	"github.com/kraxarn/ubongo/color"
	"github.com/kraxarn/ubongo/widget"
	"log"
)

type Game struct {
	ui widget.Ui
}

func NewGame() *Game {
	uiImage, err := assets.ImageUi()
	if err != nil {
		log.Fatal(err)
	}

	uiFont, err := assets.Font(assets.FontSubmenu, 14)
	if err != nil {
		log.Fatal(err)
	}

	ui := widget.NewUi()

	button := widget.NewButton(uiImage, uiFont, 64, 16, 200, 40, "Button1")
	button.SetOnPressed(func(b *widget.Button) {
		fmt.Println("Button pressed")
	})

	ui.AddWidgets(button)

	return &Game{
		ui: ui,
	}
}

func (g *Game) Update() error {
	g.ui.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Background)
	g.ui.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
