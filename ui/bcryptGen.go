package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/drognisep/bcryptgen/data"
	"golang.org/x/crypto/bcrypt"
)

func NewBcryptField() *fyne.Container {
	entry := widget.NewEntry()
	btn := widget.NewButton("Generate Hash", genBcrypt(entry))

	content := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		widget.NewLabel("Bcrypt Hash"),
		fyne.NewContainerWithLayout(
			&FieldLineLayout{},
			entry,
			btn,
		),
	)

	return content
}

func genBcrypt(entry *widget.Entry) func() {
	return func() {
		pass := data.Pass.GetState()
		hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		if err != nil {
			dialog.ShowError(err, data.MainWindow)
			entry.SetText("")
		}
		entry.SetText(string(hash))
	}
}
