package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/drognisep/bcryptgen/ui"
)

func main() {
	app := app.New()
	win := app.NewWindow("BCrypt Generator")
	win.SetMaster()

	win.SetContent(widget.NewVBox(
		ui.PasswordField(),
	))
	win.Resize(fyne.NewSize(640, 480))

	win.ShowAndRun()
}
