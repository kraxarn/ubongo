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

	dialog.setTitle(randomMessage(total))
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

func randomMessage(total time.Duration) string {
	var messages []string

	if total < time.Second*10 {
		messages = []string{
			"Amazing!",
			"Incredible!",
			"That was quick!",
		}
	} else if total > time.Minute {
		messages = []string{
			"neat",
			"Puzzle Defeated",
			"Just in time!",
			"Phew!",
		}
	} else {
		messages = []string{
			"You won!",
			"You did it!",
			"Cool!",
			"Congratulations!",
		}
	}

	// Seed should be reset when loading the next level,
	// so it should be safe to change it here
	rand.Seed(time.Now().Unix())
	return messages[rand.Intn(len(messages))]
}
