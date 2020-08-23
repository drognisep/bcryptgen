package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/drognisep/bcryptgen/data"
	"golang.org/x/crypto/bcrypt"
)

// NewBcryptField creates a new container with the content initialized.
func NewBcryptField() *fyne.Container {
	entry := widget.NewEntry()
	entry.OnChanged = func(newHash string) {
		data.Hash.SetStateNoBroadcast(newHash)
	}
	btn := widget.NewButton("Generate Hash", genBcrypt)
	data.Hash.Attach(func(s string) {
		entry.SetText(s)
	})
	data.Pass.Attach(func(s string) {
		entry.SetText("")
	})

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

func genBcrypt() {
	pass := data.Pass.GetState()
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		dialog.ShowError(err, data.MainWindow)
		data.Hash.SetState("")
		return
	}
	data.Hash.SetState(string(hash))
}
