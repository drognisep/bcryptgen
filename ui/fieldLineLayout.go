package ui

import (
	"fyne.io/fyne"
)

// FieldLineLayout will layout two objects side by side, with the first taking up all available width.
type FieldLineLayout struct{}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (l *FieldLineLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := 0, 0
	for _, o := range objects {
		childSize := o.MinSize()
		w += childSize.Width
		h = intMax(h, childSize.Height)
	}
	return fyne.NewSize(w, h)
}

func (l *FieldLineLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	elemHeight := l.MinSize(objects).Height
	firstElemWidth := containerSize.Width - objects[1].MinSize().Width

	objects[0].Resize(fyne.NewSize(firstElemWidth, elemHeight))
	objects[1].Resize(fyne.NewSize(objects[1].MinSize().Width, elemHeight))
	objects[1].Move(fyne.NewPos(firstElemWidth, 0))
}
