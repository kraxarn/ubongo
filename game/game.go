package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/game/colors"
	"github.com/kraxarn/ubongo/res"
	"github.com/kraxarn/ubongo/util/vec2"
	"github.com/kraxarn/ubongo/widget"
	"time"
)

type Game struct {
	ui         *widget.Ui
	size       vec2.Vector2[int]
	background *widget.RepeatImage
	scenes     *SceneManager
	seed       int64
}

func NewGame() *Game {
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
	ui.AddDebugOverlay()

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
	g.size = vec2.New(outsideWidth, outsideHeight)
	return outsideWidth, outsideHeight
}

func (g *Game) GoTo(scene Scene) {
	g.scenes.Push(scene)
}

func generateSeed() int64 {
	return time.Now().Truncate(time.Minute).Unix()
}
