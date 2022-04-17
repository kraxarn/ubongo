package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/game/entities"
	"github.com/kraxarn/ubongo/res"
	"github.com/kraxarn/ubongo/widget"
	"image"
	"math/rand"
	"time"
)

// pieceCount is the total amount of pieces
const pieceCount = 5

type SceneGame struct {
	startTime   time.Time
	ui          *widget.Ui
	currentTime *widget.Label
	pieces      [pieceCount]*entities.Piece
	piece       *entities.Piece
	board       *entities.Board
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

	pieceContainer := getPanelPos(game)
	boardSize := game.size.X - widget.ScreenPadding*2
	board := entities.NewBoard(widget.ScreenPadding, pieceContainer.Min.Y-widget.ScreenPadding-boardSize,
		boardSize, boardSize)

	return &SceneGame{
		startTime:   time.Now(),
		ui:          ui,
		currentTime: currentTime,
		pieces:      getPieces(game, imgPieces, pieceContainer, board.TileSize()),
		board:       board,
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
			// TODO: We probably want an offset from where we started, not center
			size := s.piece.Size()
			s.piece.SetPosition(pos[0].X-size.X/2, pos[0].Y-size.Y/2)
		} else {
			for i := len(s.pieces) - 1; i >= 0; i-- {
				if pos[0].In(s.pieces[i].GetRect()) {
					s.piece = s.pieces[i]
					s.pieces[i] = s.pieces[len(s.pieces)-1]
					s.pieces[len(s.pieces)-1] = s.piece
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
	s.board.Draw(screen)

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

func getPieces(game *Game, image *ebiten.Image, container image.Rectangle, tileSize int) [pieceCount]*entities.Piece {
	rand.Seed(game.seed)
	var pieces [pieceCount]*entities.Piece
	var indexes [res.PieceCount]bool

	for i := 0; i < pieceCount; i++ {
		index := nextPieceIndex()
		for indexes[index] {
			index = nextPieceIndex()
		}
		indexes[index] = true

		piece := entities.NewPiece(image, tileSize, index, 0, 0)
		pos := getPiecePos(piece, container)
		piece.SetPosition(pos.X, pos.Y)
		pieces[i] = piece
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

func getPiecePos(piece *entities.Piece, container image.Rectangle) image.Point {
	const padding = widget.ScreenPadding / 4
	minX := container.Min.X + padding
	minY := container.Min.Y + padding

	size := piece.Size()
	maxX := container.Max.X - size.X - padding
	maxY := container.Max.Y - size.Y - padding

	x := rand.Intn(maxX-minX) + minX
	y := rand.Intn(maxY-minY) + minY
	return image.Pt(x, y)
}
