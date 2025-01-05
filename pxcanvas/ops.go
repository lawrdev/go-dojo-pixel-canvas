package pxcanvas

import "fyne.io/fyne/v2"

// to get location to move the canavas to
func (pxCanvas *PxCanvas) Pan(previousCoord, currentCoord fyne.PointEvent) {
	xDiff := currentCoord.Position.X - previousCoord.AbsolutePosition.X
	yDiff := currentCoord.Position.Y - previousCoord.AbsolutePosition.Y

	pxCanvas.CanvasOffset.X += xDiff
	pxCanvas.CanvasOffset.Y += yDiff

	// referesh after changing offset to update canvas onscreen
	pxCanvas.Refresh()
}

// to scale the canvas up/down
func (pxCanvas *PxCanvas) scale(direction int) {
	switch {
	case direction > 0:
		pxCanvas.PxSize += 1

	case direction < 0:
		// not to decrease below size 1
		if pxCanvas.PxSize > 2 {
			pxCanvas.PxSize -= 1
		}

	default:
		pxCanvas.PxSize = 10
	}
}
