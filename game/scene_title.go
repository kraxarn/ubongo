package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/game/app"
	"github.com/kraxarn/ubongo/game/settings"
	"github.com/kraxarn/ubongo/res"
	"github.com/kraxarn/ubongo/widget"
	"time"
)

type Title struct {
	ui          *widget.Ui
	logo        *widget.Image
	title       *widget.Label
	seedName    *widget.Label
	musicToggle *widget.ImageButton
	music       *MusicManager
}

func NewTitle() (*Title, error) {
	ui, err := widget.NewUi()
	if err != nil {
		return nil, err
	}

	imgLogo, err := res.Image("logo")
	if err != nil {
		return nil, err
	}

	imgMusic, err := res.Image("music")
	if err != nil {
		return nil, err
	}

	// Buttons
	ui.AddStretchedButton(widget.ScreenPadding*4+widget.ButtonHeight, widget.AlignBottom, "Start Game")
	ui.AddStretchedButton(widget.ScreenPadding*3, widget.AlignBottom, "Settings")

	// Seed name
	seedName := ui.AddLabel(32, 32, res.RandomWord(time.Now().UnixNano()))

	// Music
	music, err := NewMusicManager()
	if err != nil {
		return nil, err
	}

	err = music.Play()
	if err != nil {
		return nil, err
	}

	// TODO: For testing
	opt := settings.Load()
	opt.Save()

	return &Title{
		ui:          ui,
		logo:        ui.AddImage(imgLogo, 0, 0, 64, 64),
		title:       ui.AddTitle(64, 64, app.Name),
		musicToggle: ui.AddImageButton(imgMusic, 16, 16, 50, 50),
		seedName:    seedName,
	}, nil
}

func (t *Title) Update(game *Game) error {
	// Update logo size and position
	t.logo.SetWidth(game.size.X / 2)
	logoX := game.size.X/2 - (t.logo.GetWidth() / 2)
	logoY := int(float64(game.size.Y) * 0.15)
	t.logo.SetPosition(logoX, logoY)

	// Update title position
	titleY := logoY + t.logo.GetHeight() - 8
	titleSize := t.title.Size()
	t.title.SetPosition(logoX-32, titleY)

	// Update seed name position
	seedNameSize := t.seedName.Size()
	t.seedName.SetPosition(game.size.X/2-seedNameSize.X/2,
		titleY+titleSize.Y+seedNameSize.Y+32)

	// Music toggle
	musicWidth := t.musicToggle.Size().X
	t.musicToggle.SetPosition(game.size.X-widget.ScreenPadding-musicWidth, widget.ScreenPadding)

	t.ui.Update(game.size)
	return nil
}

func (t *Title) Draw(screen *ebiten.Image) {
	t.ui.Draw(screen)
}
