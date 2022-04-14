package game

import (
	"fmt"
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
	ui.AddStretchedButton(widget.ScreenPadding*6, widget.AlignBottom, "Start Game")

	// Music toggle
	musicToggle := ui.AddImageButton(imgMusic, 16, 16, 50, 50)
	musicToggle.SetSourceRect(0, 0, 50, 50)

	// Seed name
	seedName := ui.AddLabel(32, 32, res.RandomWord(time.Now().UnixNano()))

	// Music
	music, err := NewMusicManager()
	if err != nil {
		return nil, err
	}

	// Music toggle
	musicToggle.SetOnPressed(func(b *widget.Button) {
		playing := music.IsPlaying()
		var x int
		var err error
		if playing {
			x = 0
			err = music.Stop()
		} else {
			x = 50
			err = music.Play()
		}
		if err == nil {
			musicToggle.SetSourceRect(x, 0, 50, 50)
		} else {
			fmt.Println("failed to toggle music:", err)
		}
	})

	// TODO: For testing
	opt := settings.Load()
	opt.Save()

	return &Title{
		ui:          ui,
		logo:        ui.AddImage(imgLogo, 0, 0, 64, 64),
		title:       ui.AddTitle(64, 64, app.Name),
		musicToggle: musicToggle,
		seedName:    seedName,
		music:       music,
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