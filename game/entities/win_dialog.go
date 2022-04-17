package entities

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/game/colors"
	"github.com/kraxarn/ubongo/res"
	"github.com/kraxarn/ubongo/widget"
	"image"
	"math/rand"
	"time"
)

type WinDialog struct {
	rect       image.Rectangle
	ui         *widget.Ui
	background *widget.NinePatch
	backdrop   *ebiten.Image
	info       *widget.Label
	onBack     func()
	onNext     func()
}

func NewWinDialog(total time.Duration, screenSize image.Point, x, y, w, h int) (*WinDialog, error) {
	ui, err := widget.NewUi()
	if err != nil {
		return nil, err
	}

	background := ui.AddNinePatch(res.DialogBackground, x, y, w, h)

	title := ui.AddLabel(0, 0, randomMessage())
	titleSize := title.Size()
	titleX := x + (w / 2) - (titleSize.X / 2)
	titleY := y + widget.ScreenPadding + titleSize.Y
	title.SetPosition(titleX, titleY)

	infoText := fmt.Sprintf("You completed in %.3f seconds", total.Seconds())
	info := ui.AddDebugLabel(0, 0, infoText)
	infoSize := info.Size()
	infoX := x + (w / 2) - (infoSize.X / 2)
	infoY := titleY + titleSize.Y + widget.ScreenPadding/2
	info.SetPosition(infoX, infoY)

	backdrop := ebiten.NewImage(screenSize.X, screenSize.Y)
	backdrop.Fill(colors.DialogBackdrop)

	padding := int(float64(w) * 0.15)
	backX := x + padding
	labelY := y + h - widget.ScreenPadding
	backLabel := ui.AddDebugLabel(backX, labelY, "Back to menu")
	backSize := backLabel.Size()

	nextLabel := ui.AddDebugLabel(0, 0, "Next level")
	nextSize := nextLabel.Size()
	nextX := x + w - nextSize.X - padding
	nextLabel.SetPosition(nextX, labelY)

	winDialog := &WinDialog{
		rect:       widget.Rect(x, y, w, h),
		ui:         ui,
		background: background,
		backdrop:   backdrop,
	}

	buttonY := labelY - 50 - widget.ScreenPadding/2
	back := ui.AddImageButton(res.Return, backX+backSize.X/2-25, buttonY, 50, 50)
	back.SetOnPressed(func(*widget.Button) {
		winDialog.onBack()
	})

	next := ui.AddImageButton(res.Forward, nextX+nextSize.X/2-25, buttonY, 50, 50)
	next.SetOnPressed(func(*widget.Button) {
		winDialog.onNext()
	})

	return winDialog, nil
}

func (d *WinDialog) Update(screenSize image.Point) {
	d.ui.Update(screenSize)
}

func (d *WinDialog) Draw(dst *ebiten.Image) {
	dst.DrawImage(d.backdrop, nil)
	d.ui.Draw(dst)
}

func (d *WinDialog) SetOnBack(pressed func()) {
	d.onBack = pressed
}

func (d *WinDialog) SetOnNext(pressed func()) {
	d.onNext = pressed
}

func randomMessage() string {
	messages := []string{
		"You won!",
		"You did it!",
		"Amazing!",
		"neat",
		"Cool!",
		"Congratulations!",
		"gz",
		"Puzzle Defeated",
		"Just in time!",
	}

	// Seed should be reset when loading the next level,
	// so it should be safe to change it here
	rand.Seed(time.Now().Unix())
	return messages[rand.Intn(len(messages))]
}
