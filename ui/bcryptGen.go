package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/drognisep/bcryptgen/data"
	"golang.org/x/crypto/bcrypt"
)

// BcryptComponent encapsulates the state and content of the BCrypt controls.
type BcryptComponent struct {
	entry          *widget.Entry
	generateButton *widget.Button
	content        fyne.CanvasObject
}

// Content returns the content for this component.
func (b *BcryptComponent) Content() fyne.CanvasObject {
	if b.content == nil {
		b.content = fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			widget.NewLabel("Bcrypt Hash"),
			fyne.NewContainerWithLayout(
				&FieldLineLayout{},
				b.entry,
				b.generateButton,
			),
		)
	} else {
		b.content.Refresh()
	}
	return b.content
}

// NewBcryptField creates a new container with the content initialized.
func NewBcryptField(win fyne.Window) *BcryptComponent {
	entry := widget.NewEntry()
	entry.OnChanged = func(newHash string) {
		data.Hash.SetStateNoBroadcast(newHash)
	}
	btn := widget.NewButton("Generate Hash", func() {
		if data.Pass.GetState() == "" {
			dialog.ShowInformation("No Password Set", "Set a password before generating a hash", win)
			return
		}
		genBcrypt()
	})
	data.Hash.Attach(func(s string) {
		entry.SetText(s)
	})
	data.Pass.Attach(func(s string) {
		entry.SetText("")
	})

	return &BcryptComponent{entry: entry, generateButton: btn}
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
