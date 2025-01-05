package brush

import (
	"go_dojo/pixl/apptype"

	"fyne.io/fyne/v2/driver/desktop"
)

const (
	Pixel = iota
)

func TryBrush(appstate *apptype.State, canvas apptype.Brushable, ev *desktop.MouseEvent) bool {
	switch {
	case appstate.BrushType == Pixel:
		return TryPaintPixel(appstate, canvas, ev)
	default:
		return false
	}
}

func TryPaintPixel(appstate *apptype.State, canvas apptype.Brushable, ev *desktop.MouseEvent) bool {
	x, y := canvas.MouseToCanvasXY(ev)

	if x != nil && y != nil && ev.Button == desktop.MouseButtonPrimary {
		canvas.SetColor(appstate.BrushColor, *x, *y)
		return true
	}

	return false
}
