package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

// IMPLEMENT MOUSE SCROLLABLE INTERFACE
func (pxCanvas *PxCanvas) Scrolled(ev *fyne.ScrollEvent) {
	pxCanvas.scale(int(ev.Scrolled.DY))
	pxCanvas.Refresh()
}

// called,  anytime mouse movess
func (pxCanvas *PxCanvas) MouseMoved(ev *desktop.MouseEvent) {
	pxCanvas.TryPan(pxCanvas.mouseState.previousCoord, ev)
	pxCanvas.Refresh()
	// update after TryPan completes, else, try pan will operate on same coords
	pxCanvas.mouseState.previousCoord = &ev.PointEvent
}

func (pxCanvas *PxCanvas) MouseIn(ev *desktop.MouseEvent) {}
func (pxCanvas *PxCanvas) MouseOut()                      {}
