package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/color"
	"github.com/kraxarn/ubongo/enum"
	"github.com/kraxarn/ubongo/util"
	"github.com/kraxarn/ubongo/widget"
	"log"
)

type Game struct {
	ui   *widget.Ui
	size util.Vector2[int]
}

func NewGame() *Game {
	ui, err := widget.NewUi()
	if err != nil {
		log.Fatal(err)
	}

	ui.AddStretchedButton(widget.ScreenPadding*4+widget.ButtonHeight, enum.AlignBottom, "Start Game")
	ui.AddStretchedButton(widget.ScreenPadding*3, enum.AlignBottom, "Settings")

	return &Game{
		ui: ui,
	}
}

func (g *Game) Update() error {
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
