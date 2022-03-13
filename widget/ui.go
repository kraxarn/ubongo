package widget

import "github.com/hajimehoshi/ebiten/v2"

type Ui struct {
	widgets []Widget
}

func NewUi() Ui {
	return Ui{}
}

func (u *Ui) AddWidgets(widgets ...Widget) {
	u.widgets = append(u.widgets, widgets...)
}

func (u *Ui) Update() {
	for _, widget := range u.widgets {
		widget.Update()
	}
}

func (u *Ui) Draw(screen *ebiten.Image) {
	for _, widget := range u.widgets {
		widget.Draw(screen)
	}
}
