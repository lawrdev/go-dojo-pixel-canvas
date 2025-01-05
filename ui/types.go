package ui

import (
	"go_dojo/pixl/apptype"
	"go_dojo/pixl/pxcanvas"
	"go_dojo/pixl/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
