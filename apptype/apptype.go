package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

type BrushType = int

type PxCanvasConfig struct {
	DrawingArea    fyne.Size
	CanvasOffset   fyne.Position
	PxRows, PxCols int
	PxSize         int
}

type State struct {
	BrushColor     color.Color
	BrushType      int
	SwatchSelected int
	FilePath       string
}

// to set file path
func (s *State) SetFilePath(path string) {
	s.FilePath = path
}

type Brushable interface {
	SetColor(c color.Color, x, y int)
	MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int)
}
