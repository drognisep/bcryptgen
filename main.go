package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/drognisep/bcryptgen/data"
	"github.com/drognisep/bcryptgen/ui"
)

func main() {
	app := app.New()
	win := app.NewWindow("BCrypt Generator")
	win.SetMaster()
	data.MainWindow = win
	ui.InitErrorMessage(win)

	win.SetContent(widget.NewVBox(
		ui.NewPasswordField(),
		ui.NewBcryptField(),
	))
	win.Resize(fyne.NewSize(640, 300))

	win.ShowAndRun()
}
