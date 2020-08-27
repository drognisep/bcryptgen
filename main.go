package main

import (
	"errors"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/drognisep/bcryptgen/data"
	"github.com/drognisep/bcryptgen/ui"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	app := app.New()
	app.SetIcon(ui.ResourceIconPng)
	newBcryptGen(app)
}

func newBcryptGen(app fyne.App) {
	win := app.NewWindow("BCrypt Generator")
	win.SetMaster()
	data.MainWindow = win
	copyPassBtn := widget.NewButton("Copy password", func() {
		win.Clipboard().SetContent(data.Pass.GetState())
	})
	copyHashBtn := widget.NewButton("Copy hash", func() {
		win.Clipboard().SetContent(data.Hash.GetState())
	})
	compareHashBtn := widget.NewButton("Compare password and hash", func() {
		pass := data.Pass.GetState()
		hash := data.Hash.GetState()
		if pass == "" || hash == "" {
			dialog.ShowError(errors.New("Enter a password and hash to compare them"), win)
			return
		}
		err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
		if err != nil {
			dialog.ShowError(err, win)
		} else {
			dialog.ShowInformation("Match!", "The password and hash match", win)
		}
	})

	passComponent := ui.NewPasswordField()
	bcryptComponent := ui.NewBcryptField()

	win.SetContent(widget.NewVBox(
		passComponent.Content(),
		bcryptComponent.Content(),
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			copyPassBtn,
			copyHashBtn,
			compareHashBtn,
		),
	))
	win.Resize(fyne.NewSize(640, 300))

	win.ShowAndRun()
}
