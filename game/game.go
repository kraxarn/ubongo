package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/game/colors"
	"github.com/kraxarn/ubongo/res"
	"github.com/kraxarn/ubongo/widget"
	"image"
	"math"
	"time"
)

type Game struct {
	ui         *widget.Ui
	size       image.Point
	background *widget.RepeatImage
	scenes     *SceneManager
	seed       int64
}

func NewGame(opt *Options) *Game {
	ui, err := widget.NewUi()
	if err != nil {
		panic(err)
	}

	imgPattern, err := res.Image("pattern")
	if err != nil {
		panic(err)
	}

	// Background pattern
	bg := ui.AddRepeatImage(imgPattern, 0, 0, 0, 0)

	// Debug overlay
	if opt.DebugOverlay {
		ui.AddDebugOverlay()
	}

	sceneManager := NewSceneManager()

	game := &Game{
		ui:         ui,
		background: bg,
		scenes:     sceneManager,
		seed:       generateSeed(),
	}

	// Load title scene
	titleScene, err := NewSceneTitle(game)
	if err != nil {
		panic(err)
	}
	sceneManager.Push(titleScene)

	return game
}

func (g *Game) Update() error {
	// Update pattern size
	g.background.SetSize(g.size)
	g.ui.Update(g.size)
	return g.scenes.Peek().Update(g)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colors.Background)
	g.ui.Draw(screen)
	g.scenes.Peek().Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	h := outsideHeight
	w := int(math.Min(float64(outsideWidth), float64(h)*(9.0/16.0)))
	g.size = image.Pt(w, h)
	return w, h
}

func (g *Game) GoTo(scene Scene) {
	g.scenes.Push(scene)
}

func (g *Game) GoBack() {
	g.scenes.Pop()
}

func generateSeed() int64 {
	return time.Now().Truncate(time.Minute).Unix()
}
