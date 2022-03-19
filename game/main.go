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
	ui         *widget.Ui
	size       util.Vector2[int]
	background *widget.RepeatImage
	logo       *widget.Image
	title      *widget.Label
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

	imgPattern, err := assets.Image("pattern")
	if err != nil {
		log.Fatal(err)
	}

	// Background pattern
	bg := ui.AddRepeatImage(imgPattern, 0, 0, 0, 0)

	// Logo
	logo := ui.AddImage(imgLogo, 0, 0, 64, 64)

	// Title
	title := ui.AddTitle(64, 64, "Ubongo")

	// Buttons
	ui.AddStretchedButton(widget.ScreenPadding*4+widget.ButtonHeight, enum.AlignBottom, "Start Game")
	ui.AddStretchedButton(widget.ScreenPadding*3, enum.AlignBottom, "Settings")

	return &Game{
		ui:         ui,
		logo:       logo,
		title:      title,
		background: bg,
	}
}

func (g *Game) Update() error {
	// Update pattern size
	g.background.SetSize(g.size)

	// Update logo size and position
	g.logo.SetWidth(g.size.X / 2)
	logoX := g.size.X/2 - (g.logo.GetWidth() / 2)
	logoY := int(float64(g.size.Y) * 0.1)
	g.logo.SetPosition(logoX, logoY)

	// Update title position
	g.title.SetPosition(logoX-32, logoY+g.logo.GetHeight()-8)

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
