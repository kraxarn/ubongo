package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/app"
	"github.com/kraxarn/ubongo/enum"
	"github.com/kraxarn/ubongo/resources"
	"github.com/kraxarn/ubongo/settings"
	"github.com/kraxarn/ubongo/widget"
)

type Title struct {
	ui    *widget.Ui
	logo  *widget.Image
	title *widget.Label
}

func NewTitle() (*Title, error) {
	ui, err := widget.NewUi()
	if err != nil {
		return nil, err
	}

	imgLogo, err := resources.Image("logo")
	if err != nil {
		return nil, err
	}

	// Buttons
	ui.AddStretchedButton(widget.ScreenPadding*4+widget.ButtonHeight, enum.AlignBottom, "Start Game")
	ui.AddStretchedButton(widget.ScreenPadding*3, enum.AlignBottom, "Settings")

	// TODO: For testing
	opt := settings.Load()
	opt.Save()

	return &Title{
		ui:    ui,
		logo:  ui.AddImage(imgLogo, 0, 0, 64, 64),
		title: ui.AddTitle(64, 64, app.Name),
	}, nil
}

func (t *Title) Update(game *Game) error {
	// Update logo size and position
	t.logo.SetWidth(game.size.X / 2)
	logoX := game.size.X/2 - (t.logo.GetWidth() / 2)
	logoY := int(float64(game.size.Y) * 0.1)
	t.logo.SetPosition(logoX, logoY)

	// Update title position
	t.title.SetPosition(logoX-32, logoY+t.logo.GetHeight()-8)

	t.ui.Update(game.size)
	return nil
}

func (t *Title) Draw(screen *ebiten.Image) {
	t.ui.Draw(screen)
}
