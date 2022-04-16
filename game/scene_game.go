package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/game/entities"
	"github.com/kraxarn/ubongo/res"
	"github.com/kraxarn/ubongo/util/vec2"
	"github.com/kraxarn/ubongo/widget"
	"image"
	"math/rand"
	"time"
)

type SceneGame struct {
	startTime   time.Time
	ui          *widget.Ui
	currentTime *widget.Label
	pieces      []*entities.Piece
	piece       *entities.Piece
	panel       *widget.NinePatch
}

func NewSceneGame(game *Game) (*SceneGame, error) {
	ui, err := widget.NewUi()
	if err != nil {
		return nil, err
	}

	imgPieces, err := res.Image("pieces")
	if err != nil {
		return nil, err
	}

	currentTime := ui.AddLabel(0, 0, "0.0")

	// We only set an initial position to avoid jumping text
	timeSize := currentTime.Size()
	currentTime.SetPosition(game.size.X/2-timeSize.X/2, widget.ScreenPadding+timeSize.Y)

	return &SceneGame{
		startTime:   time.Now(),
		ui:          ui,
		currentTime: currentTime,
		pieces:      getPieces(game, imgPieces),
		panel:       ui.AddNinePatch(res.PanelBackground, 0, 0, 0, 0),
	}, nil
}

func (s *SceneGame) Update(game *Game) error {
	s.ui.Update(game.size)
	s.currentTime.SetText(s.elapsedTime())

	// Panel for pieces
	s.panel.SetTargetRect(getPanelPos(game))

	pos := widget.TouchPositions()
	if len(pos) > 0 {
		if s.piece != nil {
			s.piece.SetPosition(pos[0].X, pos[0].Y)
		} else {
			for _, piece := range s.pieces {
				if vec2.ToPoint(pos[0]).In(piece.GetRect()) {
					s.piece = piece
					break
				}
			}
		}
	} else {
		s.piece = nil
	}

	for _, piece := range s.pieces {
		piece.Update()
	}

	return nil
}

func (s *SceneGame) Draw(screen *ebiten.Image) {
	s.ui.Draw(screen)

	for _, piece := range s.pieces {
		piece.Draw(screen)
	}
}

func (s *SceneGame) elapsedTime() string {
	duration := time.Now().Sub(s.startTime)
	return fmt.Sprintf("%.1f", duration.Seconds())
}

func nextPieceIndex() int {
	return rand.Intn(len(res.PieceImageRects))
}

func getPieces(game *Game, image *ebiten.Image) []*entities.Piece {
	rand.Seed(game.seed)
	container := getPanelPos(game)
	var pieces []*entities.Piece

	for i := 0; i < 5; i++ {
		// TODO: Avoid duplicates
		index := nextPieceIndex()
		piece := entities.NewPiece(image, index, 0, 0)

		pos := getPiecePos(piece, container)
		piece.SetPosition(pos.X, pos.Y)
		pieces = append(pieces, piece)
	}

	return pieces
}

func getPanelPos(game *Game) image.Rectangle {
	panelX0 := widget.ScreenPadding
	panelY0 := int(float64(game.size.Y) * 0.6)
	panelX1 := game.size.X - widget.ScreenPadding
	panelY1 := game.size.Y - widget.ScreenPadding
	return image.Rect(panelX0, panelY0, panelX1, panelY1)
}

func getPiecePos(piece *entities.Piece, container image.Rectangle) vec2.Vector2[int] {
	const padding = widget.ScreenPadding / 4
	minX := container.Min.X + padding
	minY := container.Min.Y + padding

	size := piece.Size()
	maxX := container.Max.X - size.X - padding
	maxY := container.Max.Y - size.Y - padding

	x := rand.Intn(maxX-minX) + minX
	y := rand.Intn(maxY-minY) + minY
	return vec2.New(x, y)
}
