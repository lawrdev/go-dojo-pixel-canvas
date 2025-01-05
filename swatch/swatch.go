package swatch

import (
	"go_dojo/pixl/apptype"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	clickHandler func(s *Swatch)
}

// set color of swatch
func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	// need to refresh to trigger a redraw of the screen, else color will change but not updated on screen
	s.Refresh()
}

// to make a new swatch
func NewSwatch(state *apptype.State, color color.Color, swatchIndex int, clickHandler func(s *Swatch)) *Swatch {
	// setup swatch with defaults
	swatch := &Swatch{
		Selected:     false,
		Color:        color,
		SwatchIndex:  swatchIndex,
		clickHandler: clickHandler,
	}

	// this supplies all state info to fyne toolkit
	swatch.ExtendBaseWidget(swatch)

	return swatch
}

func (s *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(s.Color)
	objects := []fyne.CanvasObject{square}
	return &SwatchRenderer{square: *square, objects: objects, parent: s}
}
