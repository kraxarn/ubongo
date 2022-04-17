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

type SceneGame struct {
	startTime   time.Time
	ui          *widget.Ui
	currentTime *widget.Label
	pieces      [entities.PieceCount]*entities.Piece
	piece       *entities.Piece
	pieceOffset image.Point
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

	currentTime := ui.AddLabel(0, 0, "000.0")

	// We only set an initial position to avoid jumping text
	timeSize := currentTime.Size()
	timeX := game.size.X - int(float64(timeSize.X)*1.25) - widget.ScreenPadding
	timeY := widget.ScreenPadding + timeSize.Y
	currentTime.SetPosition(timeX, timeY)

	panelPos := getPanelPos(game)
	panel := ui.AddNinePatch(res.PanelBackground, 0, 0, 0, 0)
	panel.SetTargetRect(panelPos)

	boardSize := game.size.X - widget.ScreenPadding*2
	pieces := getPieces(game.seed, imgPieces, panelPos, entities.TileSize(boardSize))

	boardX := widget.ScreenPadding
	boardY := panelPos.Min.Y - widget.ScreenPadding - boardSize
	board := entities.NewBoard(pieces, boardX, boardY, boardSize, boardSize)

	return &SceneGame{
		startTime:   time.Now(),
		ui:          ui,
		currentTime: currentTime,
		pieces:      pieces,
		board:       board,
		panel:       panel,
	}, nil
}

func (s *SceneGame) Update(game *Game) error {
	s.ui.Update(game.size)
	s.currentTime.SetText(s.elapsedTime())

	pos := widget.TouchPositions()
	s.updatePiece(pos)

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

func (s *SceneGame) updatePiece(pos []image.Point) {
	// Not moving, deselect
	if len(pos) <= 0 {
		s.piece = nil
		return
	}

	// Update position of current piece
	if s.piece != nil {
		x := pos[0].X - s.pieceOffset.X
		y := pos[0].Y - s.pieceOffset.Y

		if s.piece.GetRect().In(s.board.Rect()) {
			// Snap to grid
			boardPos := s.board.Position()
			tileSize := s.board.TileSize()
			x = ((x - boardPos.X) / tileSize * tileSize) + boardPos.X
			y = ((y - boardPos.Y) / tileSize * tileSize) + boardPos.Y
		}

		s.piece.SetPosition(x, y)
		return
	}

	// Check if we're pressing any piece
	for i := len(s.pieces) - 1; i >= 0; i-- {
		if pos[0].In(s.pieces[i].GetRect()) {
			// Set current piece
			s.piece = s.pieces[i]
			// Get offset from where we clicked it
			s.pieceOffset = pos[0].Sub(s.piece.GetPosition())
			// Move to back to draw last
			s.pieces[i] = s.pieces[len(s.pieces)-1]
			s.pieces[len(s.pieces)-1] = s.piece
			break
		}
	}
}

func (s *SceneGame) elapsedTime() string {
	duration := time.Now().Sub(s.startTime)
	return fmt.Sprintf("%.1f", duration.Seconds())
}

func nextPieceIndex() int {
	return rand.Intn(len(res.PieceImageRects))
}

func getPieces(seed int64, image *ebiten.Image, container image.Rectangle, tileSize int) [entities.PieceCount]*entities.Piece {
	rand.Seed(seed)
	var pieces [entities.PieceCount]*entities.Piece
	var indexes [res.PieceCount]bool

	for i := 0; i < entities.PieceCount; i++ {
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
