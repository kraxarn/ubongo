package entities

import (
	"fmt"
	"github.com/kraxarn/ubongo/res"
	"image"
	"math/rand"
	"time"
)

type WinDialog struct {
	*Dialog
}

func NewWinDialog(total time.Duration, screenSize image.Point) (*WinDialog, error) {
	dialog, err := newDialog(screenSize)
	if err != nil {
		return nil, err
	}

	dialog.setTitle(randomMessage())
	dialog.setInfo(fmt.Sprintf("You completed in %.3f seconds", total.Seconds()))

	dialog.addLeftButton(res.Return, "Back to menu")
	dialog.addRightButton(res.Forward, "Next level")

	return &WinDialog{
		Dialog: dialog,
	}, nil
}

func (d *WinDialog) SetOnBack(pressed func()) {
	d.onLeftPressed = pressed
}

func (d *WinDialog) SetOnNext(pressed func()) {
	d.onRightPressed = pressed
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
