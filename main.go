package main

import (
	"errors"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/drognisep/bcryptgen/data"
	"github.com/drognisep/bcryptgen/ui"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	myApp := app.New()
	myApp.SetIcon(ui.ResourceIconPng)
	newBcryptGen(myApp)
}

func newBcryptGen(app fyne.App) {
	win := app.NewWindow("BCrypt Generator")
	win.SetMaster()
	data.MainWindow = win
	copyPassBtn := widget.NewButton("Copy password", func() {})
	copyPass := func() {
		win.Clipboard().SetContent(data.Pass.GetState())
		copyPassBtn.SetIcon(theme.ConfirmIcon())
		time.AfterFunc(time.Second, func() {
			copyPassBtn.SetIcon(nil)
		})
	}
	copyPassBtn.OnTapped = copyPass
	copyHashBtn := widget.NewButton("Copy hash", func() {})
	copyHash := func() {
		win.Clipboard().SetContent(data.Hash.GetState())
		copyHashBtn.SetIcon(theme.ConfirmIcon())
		time.AfterFunc(time.Second, func() {
			copyHashBtn.SetIcon(nil)
		})
	}
	copyHashBtn.OnTapped = copyHash
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
	bcryptComponent := ui.NewBcryptField(win)

	win.SetContent(container.NewVBox(
		passComponent.Content(),
		bcryptComponent.Content(),
		layout.NewSpacer(),
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
