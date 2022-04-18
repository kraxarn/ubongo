package entities

import (
	"github.com/kraxarn/ubongo/res"
	"image"
)

type PauseDialog struct {
	*Dialog
}

func NewPauseDialog(screenSize image.Point) (*PauseDialog, error) {
	dialog, err := newDialog(screenSize)
	if err != nil {
		return nil, err
	}

	dialog.setTitle("Paused")
	dialog.setInfo("Game is paused")

	dialog.addLeftButton(res.Return, "Back to menu")
	dialog.addRightButton(res.Forward, "Resume")

	return &PauseDialog{
		Dialog: dialog,
	}, nil
}

func (p *PauseDialog) SetOnBack(pressed func()) {
	p.onLeftPressed = pressed
}

func (p *PauseDialog) SetOnResume(pressed func()) {
	p.onRightPressed = pressed
}
