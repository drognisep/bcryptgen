package ui

import (
	"crypto/rand"
	"math/big"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/drognisep/bcryptgen/data"
)

const (
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits       = "0123456789"
	symbols      = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

type dataDrivenEntry struct {
	widget.Entry
}

func newDataDrivenEntry() *dataDrivenEntry {
	entry := &dataDrivenEntry{}
	entry.ExtendBaseWidget(entry)

	return entry
}

// NewPasswordField creates a new password field ready to be used in the main UI
func NewPasswordField() *fyne.Container {
	passEntry := widget.NewPasswordEntry()
	obs := func(newPass string) {
		passEntry.SetText(newPass)
	}
	data.Pass.Attach(obs)
	field := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		widget.NewLabel("Password"),
		fyne.NewContainerWithLayout(
			&FieldLineLayout{},
			passEntry,
			widget.NewButton("Generate Password", showModal),
		),
	)
	return field
}

func showModal() {
	var popup *widget.PopUp

	passEntry := widget.NewEntry()
	passEntry.OnChanged = func(newVal string) {
		data.Pass.SetState(newVal)
	}
	upperAlphaCheck := widget.NewCheck("Use uppercase alpha characters", func(b bool) {})
	upperAlphaCheck.SetChecked(true)
	lowerAlphaCheck := widget.NewCheck("Use lowercase alpha characters", func(b bool) {})
	lowerAlphaCheck.SetChecked(true)
	numCheck := widget.NewCheck("Use numeric characters", func(b bool) {})
	numCheck.SetChecked(true)
	specCheck := widget.NewCheck("Use special characters", func(b bool) {})
	specCheck.SetChecked(true)

	gen := func() {
		s := generatePassword(10, upperAlphaCheck.Checked, lowerAlphaCheck.Checked, numCheck.Checked, specCheck.Checked)
		passEntry.SetText(s)
	}

	content := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		passEntry,
		upperAlphaCheck,
		lowerAlphaCheck,
		numCheck,
		specCheck,
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			widget.NewButton("Generate", gen),
			widget.NewButton("Done", func() {
				popup.Hide()
			}),
		),
	)

	popup = widget.NewModalPopUp(content, data.MainWindow.Canvas())
	popup.Show()
}

func generatePassword(length int, useUpAlpha bool, useLowAlpha bool, useNum bool, useSpecial bool) string {
	buf := make([]rune, length)
	var strings string

	if useUpAlpha {
		strings += upperLetters
	}
	if useLowAlpha {
		strings += lowerLetters
	}
	if useNum {
		strings += digits
	}
	if useSpecial {
		strings += symbols
	}
	choices := []rune(strings)

	max := len(choices)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
		if err != nil {
			ShowErrorMessage("Failed to get random data for password")
			return ""
		}
		buf[i] = choices[int(num.Int64())]
	}

	return string(buf)
}
