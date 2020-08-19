package ui

import (
	"errors"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
)

var w fyne.Window

// InitErrorMessage sets the root window to use to display an error message
func InitErrorMessage(window fyne.Window) {
	w = window
}

// ShowErrorMessage uses the default window to display an error dialog.
func ShowErrorMessage(errorMessage string) {
	dialog.ShowError(errors.New(errorMessage), w)
}
