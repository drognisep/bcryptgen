package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/drognisep/bcryptgen/data"
	"github.com/drognisep/bcryptgen/ui"
)

func main() {
	app := app.New()
	win := app.NewWindow("BCrypt Generator")
	win.SetMaster()
	data.MainWindow = win
	copyPassBtn := widget.NewButton("Copy password", func() {
		win.Clipboard().SetContent(data.Pass.GetState())
	})
	copyHashBtn := widget.NewButton("Copy hash", func() {
		win.Clipboard().SetContent(data.Hash.GetState())
	})

	win.SetContent(widget.NewVBox(
		ui.NewPasswordField(),
		ui.NewBcryptField(),
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			copyPassBtn,
			copyHashBtn,
		),
	))
	win.Resize(fyne.NewSize(640, 300))

	win.ShowAndRun()
}
