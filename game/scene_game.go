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
	totalDuration time.Duration
	prevTime      time.Time
	ui            *widget.Ui
	currentTime   *widget.Label
	pieces        [entities.PieceCount]*entities.Piece
	piece         *entities.Piece
	pieceOffset   image.Point
	pieceStart    image.Point
	board         *entities.Board
	panel         *widget.NinePatch
	level         int64
	levelText     *widget.Label
	game          *Game
	winDialog     *entities.WinDialog
	pauseDialog   *entities.PauseDialog
	debugTile     *ebiten.Image
	debugPoint    *ebiten.Image
}

func NewSceneGame(game *Game, level int64) (*SceneGame, error) {
	ui, err := widget.NewUi()
	if err != nil {
		return nil, err
	}

	imgPieces, err := res.Image("pieces")
	if err != nil {
		return nil, err
	}

	// Seed for current level
	rand.Seed(game.seed + level)

	panelPos := getPanelPos(game)
	panel := ui.AddNinePatch(res.BackgroundPanel, 0, 0, 0, 0)
	panel.SetTargetRect(panelPos)

	boardSize := game.size.X - widget.ScreenPadding*2
	pieces := getPieces(imgPieces, panelPos, entities.TileSize(boardSize))

	boardX := widget.ScreenPadding
	boardY := panelPos.Min.Y - widget.ScreenPadding - boardSize
	board := entities.NewBoard(pieces, boardX, boardY, boardSize, boardSize)

	const pauseSize = 50
	pauseX := game.size.X - pauseSize - widget.ScreenPadding*2
	pauseY := boardY - pauseSize
	pause := ui.AddImageButton(res.Pause, pauseX, pauseY, pauseSize, pauseSize)

	levelLabel := ui.AddLabel(0, 0, fmt.Sprintf("Level %d", level))
	levelSize := levelLabel.Size()
	levelX := pauseX - levelSize.X - widget.ScreenPadding
	levelY := pauseY + levelSize.Y + ((pauseSize - levelSize.Y) / 2)
	levelLabel.SetPosition(levelX, levelY)

	currentTime := ui.AddLabel(widget.ScreenPadding*2, levelY, "0.0")

	scene := &SceneGame{
		prevTime:    time.Now(),
		ui:          ui,
		currentTime: currentTime,
		pieces:      pieces,
		board:       board,
		panel:       panel,
		level:       level,
		game:        game,
	}

	pause.SetOnPressed(func(*widget.Button) {
		if dialog, err := scene.getPauseDialog(); err == nil {
			scene.pauseDialog = dialog
		}
	})

	if game.opt.DebugMode {
		scene.debugTile = debugTile(board.TileSize())
		scene.debugPoint = debugPoint()
	}

	return scene, nil
}

func (s *SceneGame) Update(game *Game) error {
	s.ui.Update(game.size)

	if s.pauseDialog != nil {
		s.pauseDialog.Update(game.size)
		return nil
	}

	if s.winDialog == nil {
		s.currentTime.SetText(s.elapsedTime())
	}

	if s.winDialog != nil {
		s.winDialog.Update(game.size)
	}

	pos := widget.TouchPositions()
	s.updatePiece(pos)

	for _, piece := range s.pieces {
		piece.Update()
	}

	now := time.Now()
	s.totalDuration += now.Sub(s.prevTime)
	s.prevTime = now

	return nil
}

func (s *SceneGame) Draw(screen *ebiten.Image) {
	s.ui.Draw(screen)
	s.board.Draw(screen)

	for _, piece := range s.pieces {
		piece.Draw(screen)
	}

	if s.pauseDialog != nil {
		s.pauseDialog.Draw(screen)
	}

	if s.winDialog != nil {
		s.winDialog.Draw(screen)
	}

	if s.debugTile != nil {
		s.drawDebugTile(screen)
	}
}

func (s *SceneGame) updatePiece(pos []image.Point) {
	// Not moving, deselect
	if len(pos) <= 0 {
		if s.piece != nil {
			if s.board.AllTilesFilled(s.pieces) {
				var err error
				s.winDialog, err = s.getWinDialog(s.totalDuration)
				if err != nil {
					fmt.Println("Failed to instance win dialog:", err)
				}
			} else if s.shouldRotatePiece() {
				s.piece.Rotate(90)
			}
			s.piece = nil
		}
		return
	}

	// Update position of current piece
	if s.piece != nil {
		x := pos[0].X - s.pieceOffset.X
		y := pos[0].Y - s.pieceOffset.Y

		if s.piece.Rect().In(s.board.Rect()) {
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
		if pos[0].In(s.pieces[i].Rect()) {
			// Set current piece
			s.piece = s.pieces[i]
			// Get offset from where we clicked it
			s.pieceOffset = pos[0].Sub(s.piece.GetPosition())
			// Where we started dragging from
			s.pieceStart = s.piece.GetPosition()
			// Move to back to draw last
			s.pieces[i] = s.pieces[len(s.pieces)-1]
			s.pieces[len(s.pieces)-1] = s.piece
			break
		}
	}
}

func (s *SceneGame) elapsedTime() string {
	return fmt.Sprintf("%.1f", s.totalDuration.Seconds())
}

func (s *SceneGame) getWinDialog(total time.Duration) (*entities.WinDialog, error) {
	dialog, err := entities.NewWinDialog(total, s.game.size)
	if err != nil {
		return nil, err
	}

	dialog.SetOnBack(func() {
		s.game.GoBack()
	})

	dialog.SetOnNext(func() {
		var scene *SceneGame
		if scene, err = NewSceneGame(s.game, s.level+1); err == nil {
			s.game.GoBack()
			s.game.GoTo(scene)
		}
	})

	return dialog, nil
}

func (s *SceneGame) getPauseDialog() (*entities.PauseDialog, error) {
	dialog, err := entities.NewPauseDialog(s.game.size)
	if err != nil {
		return nil, err
	}

	dialog.SetOnBack(func() {
		s.game.GoBack()
	})

	dialog.SetOnResume(func() {
		s.pauseDialog = nil
		s.prevTime = time.Now()
	})

	return dialog, nil
}

func (s *SceneGame) shouldRotatePiece() bool {
	threshold := s.board.TileSize() / 8
	dist := s.piece.GetPosition().Sub(s.pieceStart)
	return dist.X < threshold && dist.X > -threshold && dist.Y < threshold && dist.Y > -threshold
}

func nextPieceIndex() int {
	return rand.Intn(len(res.PieceImageRects))
}

func getPieces(image *ebiten.Image, container image.Rectangle, tileSize int) [entities.PieceCount]*entities.Piece {
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
