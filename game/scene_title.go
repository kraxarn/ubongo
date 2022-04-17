package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/game/app"
	"github.com/kraxarn/ubongo/game/settings"
	"github.com/kraxarn/ubongo/res"
	"github.com/kraxarn/ubongo/widget"
)

type SceneTitle struct {
	ui          *widget.Ui
	logo        *widget.Image
	title       *widget.Label
	version     *widget.Label
	seedName    *widget.Label
	musicToggle *widget.ImageButton
	music       *MusicManager
}

func NewSceneTitle(game *Game) (*SceneTitle, error) {
	ui, err := widget.NewUi()
	if err != nil {
		return nil, err
	}

	imgLogo, err := res.Image("logo")
	if err != nil {
		return nil, err
	}

	// Buttons
	start := ui.AddStretchedButton(widget.ScreenPadding*4+widget.ButtonHeight, widget.AlignBottom, "Start Game")
	start.SetOnPressed(func(*widget.Button) {
		sceneGame, err := NewSceneGame(game)
		if err == nil {
			game.GoTo(sceneGame)
		}
	})

	// Settings
	setting := settings.Load()

	// Music toggle
	var musicImage res.UiImageType
	if setting.IsMusicEnabled() {
		musicImage = res.MusicOn
	} else {
		musicImage = res.MusicOff
	}
	musicToggle := ui.AddImageButton(musicImage, 16, 16, 50, 50)

	// Seed name
	seedName := ui.AddLabel(32, 32, res.RandomWord(game.seed))

	newSeed := ui.AddStretchedButton(widget.ScreenPadding*3, widget.AlignBottom, "Generate Seed")
	newSeed.SetOnPressed(func(*widget.Button) {
		game.seed = generateSeed()
		seedName.SetText(res.RandomWord(game.seed))
	})

	// Music
	music, err := NewMusicManager()
	if err != nil {
		return nil, err
	}

	if setting.IsMusicEnabled() {
		music.volume = setting.MusicVolume
		if err = music.Play(); err != nil {
			fmt.Println("failed to play music:", err)
			setting.ToggleMusic(false)
		}
	}

	// Music toggle
	musicToggle.SetOnPressed(func(b *widget.Button) {
		musicEnabled := setting.IsMusicEnabled()
		setting.ToggleMusic(!musicEnabled)
		var imageType res.UiImageType
		var err error
		if musicEnabled {
			imageType = res.MusicOff
			err = music.Stop()
		} else {
			imageType = res.MusicOn
			music.volume = setting.MusicVolume
			err = music.Play()
		}
		if err != nil {
			fmt.Println("failed to toggle music:", err)
		} else {
			musicToggle.SetImageType(imageType)
			setting.Save()
		}
	})

	return &SceneTitle{
		ui:          ui,
		logo:        ui.AddImage(imgLogo, 0, 0, 64, 64),
		title:       ui.AddTitle(64, 64, app.Name),
		version:     ui.AddDebugLabel(64, 64, app.Version()),
		musicToggle: musicToggle,
		seedName:    seedName,
		music:       music,
	}, nil
}

func (t *SceneTitle) Update(game *Game) error {
	// Update logo size and position
	t.logo.SetWidth(game.size.X / 2)
	logoX := game.size.X/2 - (t.logo.GetWidth() / 2)
	logoY := int(float64(game.size.Y) * 0.15)
	t.logo.SetPosition(logoX, logoY)

	// Update title position
	titleX := logoX - widget.ScreenPadding
	titleY := logoY + t.logo.GetHeight() - 8
	titleSize := t.title.Size()
	t.title.SetPosition(titleX, titleY)

	// Update version position
	verX := titleX + titleSize.X - t.version.Size().X
	verY := titleY + titleSize.Y - 12
	t.version.SetPosition(verX, verY)

	// Update seed name position
	seedNameSize := t.seedName.Size()
	t.seedName.SetPosition(game.size.X/2-seedNameSize.X/2,
		titleY+titleSize.Y+seedNameSize.Y+40)

	// Music toggle
	musicWidth := t.musicToggle.Size().X
	t.musicToggle.SetPosition(game.size.X-widget.ScreenPadding-musicWidth, widget.ScreenPadding)

	t.ui.Update(game.size)
	return nil
}

func (t *SceneTitle) Draw(screen *ebiten.Image) {
	t.ui.Draw(screen)
}
