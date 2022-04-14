package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/widget"
	"time"
)

type SceneGame struct {
	startTime   time.Time
	currentTime *widget.Label
}

func NewSceneGame(game *Game) (*SceneGame, error) {
	ui, err := widget.NewUi()
	if err != nil {
		return nil, err
	}

	currentTime := ui.AddLabel(0, 0, "0.0")

	// We only set position first to avoid jumping text
	timeSize := currentTime.Size()
	currentTime.SetPosition(game.size.X/2-timeSize.X/2, widget.ScreenPadding+timeSize.Y)

	return &SceneGame{
		startTime:   time.Now(),
		currentTime: currentTime,
	}, nil
}

func (s *SceneGame) Update(game *Game) error {
	s.currentTime.SetText(s.elapsedTime())
	return nil
}

func (s *SceneGame) Draw(screen *ebiten.Image) {
	s.currentTime.Draw(screen)
}

func (s *SceneGame) elapsedTime() string {
	duration := time.Now().Sub(s.startTime)
	return fmt.Sprintf("%.1f", duration.Seconds())
}
