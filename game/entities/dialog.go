package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/game/colors"
	"github.com/kraxarn/ubongo/res"
	"github.com/kraxarn/ubongo/widget"
	"image"
)

type Dialog struct {
	rect           image.Rectangle
	ui             *widget.Ui
	background     *widget.NinePatch
	backdrop       *ebiten.Image
	title          *widget.Label
	info           *widget.Label
	onLeftPressed  func()
	onRightPressed func()
}

func newDialog(screenSize image.Point) (*Dialog, error) {
	ui, err := widget.NewUi()
	if err != nil {
		return nil, err
	}

	w := screenSize.X - widget.ScreenPadding*4
	h := screenSize.Y / 4
	x := widget.ScreenPadding * 2
	y := screenSize.Y/2 - h/2

	backdrop := ebiten.NewImage(screenSize.X, screenSize.Y)
	backdrop.Fill(colors.DialogBackdrop)

	background := ui.AddNinePatch(res.BackgroundDialog, x, y, w, h)

	dialog := &Dialog{
		rect:       widget.Rect(x, y, w, h),
		ui:         ui,
		background: background,
		backdrop:   backdrop,
	}

	return dialog, err
}

func (d *Dialog) Update(screenSize image.Point) {
	d.ui.Update(screenSize)
}

func (d *Dialog) Draw(dst *ebiten.Image) {
	dst.DrawImage(d.backdrop, nil)
	d.ui.Draw(dst)
}

func (d *Dialog) setTitle(text string) {
	if d.title == nil {
		d.title = d.ui.AddLabel(0, 0, text)
	} else {
		d.title.SetText(text)
	}

	titleSize := d.title.Size()
	pos := d.rect.Min
	titleX := pos.X + (d.rect.Dx() / 2) - (titleSize.X / 2)
	titleY := pos.Y + widget.ScreenPadding + titleSize.Y
	d.title.SetPosition(titleX, titleY)
}

func (d *Dialog) setInfo(text string) {
	if d.info == nil {
		d.info = d.ui.AddDebugLabel(0, 0, text)
	} else {
		d.info.SetText(text)
	}

	var titleSize image.Point
	var titlePos image.Point
	if d.title == nil {
		titleSize = d.title.Size()
		titlePos = d.title.GetPosition()
	}

	infoSize := d.info.Size()
	infoX := titlePos.X + (titleSize.X / 2) - (infoSize.X / 2)
	infoY := titlePos.Y + titleSize.Y + widget.ScreenPadding/2
	d.info.SetPosition(infoX, infoY)
}

func (d *Dialog) addButton(i int, icon res.UiImageType, text string) *widget.ImageButton {
	part := d.rect.Dx() / 3
	centerX := (part * i) + (part / 2)

	label := d.ui.AddDebugLabel(0, 0, text)
	labelSize := label.Size()
	labelX := centerX - labelSize.X/2
	labelY := d.rect.Max.Y - widget.ScreenPadding
	label.SetPosition(labelX, labelY)

	buttonX := centerX - 25
	buttonY := labelY - 50
	return d.ui.AddImageButton(icon, buttonX, buttonY, 50, 50)
}

func (d *Dialog) addLeftButton(icon res.UiImageType, text string) {
	d.addButton(0, icon, text).SetOnPressed(func(*widget.Button) {
		if d.onLeftPressed != nil {
			d.onLeftPressed()
		}
	})
}

func (d *Dialog) addRightButton(icon res.UiImageType, text string) {
	d.addButton(2, icon, text).SetOnPressed(func(*widget.Button) {
		if d.onRightPressed != nil {
			d.onRightPressed()
		}
	})
}
