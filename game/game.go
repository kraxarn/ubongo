package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/game/colors"
	"github.com/kraxarn/ubongo/game/resources"
	"github.com/kraxarn/ubongo/util/vec2"
	"github.com/kraxarn/ubongo/widget"
)

type Game struct {
	ui         *widget.Ui
	size       vec2.Vector2[int]
	background *widget.RepeatImage
	scenes     *SceneManager
}

func NewGame() *Game {
	ui, err := widget.NewUi()
	if err != nil {
		panic(err)
	}

	imgPattern, err := resources.Image("pattern")
	if err != nil {
		panic(err)
	}

	// Background pattern
	bg := ui.AddRepeatImage(imgPattern, 0, 0, 0, 0)

	// Debug overlay
	ui.AddDebugOverlay()

	// Load title scene
	titleScene, err := NewTitle()
	if err != nil {
		panic(err)
	}

	sceneManager := NewSceneManager()
	sceneManager.Push(titleScene)

	return &Game{
		ui:         ui,
		background: bg,
		scenes:     sceneManager,
	}
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
