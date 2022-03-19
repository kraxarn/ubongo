package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/assets"
	"github.com/kraxarn/ubongo/color"
	"github.com/kraxarn/ubongo/enum"
	"github.com/kraxarn/ubongo/util"
	"github.com/kraxarn/ubongo/widget"
	"log"
)

type Game struct {
	ui   *widget.Ui
	size util.Vector2[int]
	logo *widget.Image
}

func NewGame() *Game {
	ui, err := widget.NewUi()
	if err != nil {
		log.Fatal(err)
	}

	imgLogo, err := assets.Image("logo")
	if err != nil {
		log.Fatal(err)
	}

	// Logo
	logo := ui.AddImage(imgLogo, 0, 0, 64, 64)

	// Buttons
	ui.AddStretchedButton(widget.ScreenPadding*4+widget.ButtonHeight, enum.AlignBottom, "Start Game")
	ui.AddStretchedButton(widget.ScreenPadding*3, enum.AlignBottom, "Settings")

	return &Game{
		ui:   ui,
		logo: logo,
	}
}

func (g *Game) Update() error {
	// Update logo size and position
	g.logo.SetWidth(g.size.X / 2)
	g.logo.SetPosition(g.size.X/2-(g.logo.GetWidth()/2), int(float64(g.size.Y)*0.1))

	g.ui.Update(g.size)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Background)
	g.ui.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	g.size = util.Vec2(outsideWidth, outsideHeight)
	return outsideWidth, outsideHeight
}
