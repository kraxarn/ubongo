package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/res"
	"github.com/kraxarn/ubongo/widget"
	"math/rand"
	"time"
)

type SceneGame struct {
	startTime   time.Time
	currentTime *widget.Label
	tempTile    *widget.Image
}

func NewSceneGame(game *Game) (*SceneGame, error) {
	ui, err := widget.NewUi()
	if err != nil {
		return nil, err
	}

	imgTiles, err := res.Image("pieces")
	if err != nil {
		return nil, err
	}

	tempTile := ui.AddImage(imgTiles, 0, 0, 0, 0)

	rand.Seed(game.seed)
	tempTile.SetSourceRect(res.PieceImageRects[rand.Int()%len(res.PieceImageRects)])

	currentTime := ui.AddLabel(0, 0, "0.0")

	// We only set position first to avoid jumping text
	timeSize := currentTime.Size()
	currentTime.SetPosition(game.size.X/2-timeSize.X/2, widget.ScreenPadding+timeSize.Y)

	return &SceneGame{
		startTime:   time.Now(),
		currentTime: currentTime,
		tempTile:    tempTile,
	}, nil
}

func (s *SceneGame) Update(game *Game) error {
	s.currentTime.SetText(s.elapsedTime())

	tileWidth := s.tempTile.GetWidth()
	tileHeight := s.tempTile.GetHeight()

	s.tempTile.SetPosition(game.size.X/2-tileWidth/2, game.size.Y-widget.ScreenPadding-tileHeight)

	return nil
}

func (s *SceneGame) Draw(screen *ebiten.Image) {
	s.currentTime.Draw(screen)
	s.tempTile.Draw(screen)
}

func (s *SceneGame) elapsedTime() string {
	duration := time.Now().Sub(s.startTime)
	return fmt.Sprintf("%.1f", duration.Seconds())
}
