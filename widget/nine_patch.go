package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

// NinePatch is a nine-patch image
type NinePatch struct {
	image      *ebiten.Image
	targetRect image.Rectangle
	sourceRect image.Rectangle
}

func NewNinePatch(image *ebiten.Image, x, y, w, h int) *NinePatch {
	return &NinePatch{
		image:      image,
		targetRect: rect(x, y, w, h),
		sourceRect: rect(0, 0, w, h),
	}
}

func (n *NinePatch) Update(*Ui) {
}

func (n *NinePatch) Draw(dst *ebiten.Image) {
	srcX := n.sourceRect.Min.X
	srcY := n.sourceRect.Min.Y
	srcW := n.sourceRect.Dx()
	srcH := n.sourceRect.Dy()

	dstX := n.targetRect.Min.X
	dstY := n.targetRect.Min.Y
	dstW := n.targetRect.Dx()
	dstH := n.targetRect.Dy()

	op := &ebiten.DrawImageOptions{}
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			op.GeoM.Reset()

			sx := srcX
			sy := srcY
			sw := srcW / 4
			sh := srcH / 4
			dx := 0
			dy := 0
			dw := sw
			dh := sh

			switch i {
			case 1:
				sx = srcX + srcW/4
				sw = srcW / 2
				dx = srcW / 4
				dw = dstW - 2*srcW/4

			case 2:
				sx = srcX + 3*srcW/4
				dx = dstW - srcW/4
			}

			switch j {
			case 1:
				sy = srcY + srcH/4
				sh = srcH / 2
				dy = srcH / 4
				dh = dstH - 2*srcH/4

			case 2:
				sy = srcY + 3*srcH/4
				dy = dstH - srcH/4
			}

			op.GeoM.Scale(float64(dw)/float64(sw), float64(dh)/float64(sh))
			op.GeoM.Translate(float64(dx), float64(dy))
			op.GeoM.Translate(float64(dstX), float64(dstY))

			dst.DrawImage(n.image.SubImage(image.Rect(sx, sy, sx+sw, sy+sh)).(*ebiten.Image), op)
		}
	}
}

func (n *NinePatch) SetSourceRect(rect image.Rectangle) {
	n.sourceRect = rect
}

func (n *NinePatch) SetTargetRect(rect image.Rectangle) {
	n.targetRect = rect
}
