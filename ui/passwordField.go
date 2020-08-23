package ui

import (
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/drognisep/bcryptgen/data"
	"github.com/sethvargo/go-password/password"
)

const (
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits       = "0123456789"
	symbols      = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

var max = 64
var min = 8

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

func strLengthOptions(min, max int) []string {
	olen := max - min + 1
	opts := make([]string, olen)
	for i := 0; i < olen; i++ {
		opts[i] = strconv.Itoa(min + i)
	}
	return opts
}

func showModal() {
	var popup *widget.PopUp

	passEntry := widget.NewEntry()
	passEntry.OnChanged = func(newVal string) {
		data.Pass.SetState(newVal)
	}

	passLengthSelect := widget.NewSelect(strLengthOptions(min, max), func(string) {})
	passLengthSelect.SetSelected(string(min))
	numDigitsSelect := widget.NewSelect(strLengthOptions(0, max), func(string) {})
	numDigitsSelect.SetSelected(string(min / 3))
	numSymbolsSelect := widget.NewSelect(strLengthOptions(0, max), func(string) {})
	numSymbolsSelect.SetSelected(string(min / 3))
	upperAlphaCheck := widget.NewCheck("Use uppercase letters", func(b bool) {})
	upperAlphaCheck.SetChecked(true)
	allowRepeatsCheck := widget.NewCheck("Allow repeated characters", func(b bool) {})
	allowRepeatsCheck.SetChecked(false)

	gen := func() {
		passLength, _ := strconv.Atoi(passLengthSelect.Selected)
		numDigits, _ := strconv.Atoi(numDigitsSelect.Selected)
		numSymbols, _ := strconv.Atoi(numSymbolsSelect.Selected)
		s := generatePassword(
			passLength,
			numDigits,
			numSymbols,
			upperAlphaCheck.Checked,
			allowRepeatsCheck.Checked,
		)
		passEntry.SetText(s)
	}

	content := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		passEntry,
		fyne.NewContainerWithLayout(
			&FieldLineLayout{},
			widget.NewLabel("Password Length"),
			passLengthSelect,
		),
		fyne.NewContainerWithLayout(
			&FieldLineLayout{},
			widget.NewLabel("Number of digits"),
			numDigitsSelect,
		),
		fyne.NewContainerWithLayout(
			&FieldLineLayout{},
			widget.NewLabel("Number of symbols"),
			numSymbolsSelect,
		),
		upperAlphaCheck,
		allowRepeatsCheck,
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

func generatePassword(length, numDigits, numSymbols int, useUpAlpha, allowRepeats bool) string {
	pass, err := password.Generate(length, numDigits, numSymbols, !useUpAlpha, allowRepeats)
	if err != nil {
		dialog.ShowError(err, data.MainWindow)
		return ""
	}
	return pass
}
