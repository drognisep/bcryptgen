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

var max = 64
var min = 8
var popup *widget.PopUp

func strLengthOptions(min, max int) []string {
	olen := max - min + 1
	opts := make([]string, olen)
	for i := 0; i < olen; i++ {
		opts[i] = strconv.Itoa(min + i)
	}
	return opts
}

type generatePasswordModal struct {
	passEntry         *widget.Entry
	numDigitsSelect   *widget.Select
	numSymbolsSelect  *widget.Select
	passLengthSelect  *widget.Select
	upperAlphaCheck   *widget.Check
	allowRepeatsCheck *widget.Check
	popup             *widget.PopUp
	genBtn            *widget.Button
	doneBtn           *widget.Button
}

// PasswordComponent encapsulates the state needed to feed the view.
type PasswordComponent struct {
	passEntry      *widget.Entry
	generateButton *widget.Button
	modal          *generatePasswordModal
	content        fyne.CanvasObject
}

// Content returns rendered content for this component.
func (p *PasswordComponent) Content() fyne.CanvasObject {
	if p.content == nil {
		p.content = fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			widget.NewLabel("Password"),
			fyne.NewContainerWithLayout(
				&FieldLineLayout{},
				p.passEntry,
				p.generateButton,
			),
		)
	} else {
		p.content.Refresh()
	}
	return p.content
}

// NewPasswordField creates a new password field ready to be used in the main UI
func NewPasswordField() *PasswordComponent {
	generateButton := widget.NewButton("Generate Password", func() {})
	passEntry := widget.NewPasswordEntry()
	passEntry.OnChanged = func(newPass string) {
		data.Pass.SetStateNoBroadcast(newPass)
	}
	data.Pass.Attach(func(newPass string) {
		passEntry.SetText(newPass)
	})

	component := &PasswordComponent{
		passEntry:      passEntry,
		generateButton: generateButton,
		modal:          nil,
	}
	generateButton.OnTapped = component.showModal
	return component
}

func (p *PasswordComponent) showModal() {
	p.createGenPassModal()
	p.modal.popup.Show()
}

func (p *PasswordComponent) createGenPassModal() {
	if p.modal == nil {
		passEntry := widget.NewEntry()
		passEntry.OnChanged = func(newVal string) {
			data.Pass.SetState(newVal)
		}
		passEntry.ReadOnly = true

		numDigitsSelect := widget.NewSelect(strLengthOptions(0, max), func(string) {})
		numDigitsSelect.SetSelected(strconv.Itoa(min / 3))
		numSymbolsSelect := widget.NewSelect(strLengthOptions(0, max), func(string) {})
		numSymbolsSelect.SetSelected(strconv.Itoa(min / 3))
		passLengthSelect := widget.NewSelect(strLengthOptions(min, max), func(newLen string) {
			newLeni, _ := strconv.Atoi(newLen)
			numDigitsSelect.SetSelected(strconv.Itoa(newLeni / 3))
			numSymbolsSelect.SetSelected(strconv.Itoa(newLeni / 3))
		})
		passLengthSelect.SetSelected(strconv.Itoa(min))
		upperAlphaCheck := widget.NewCheck("Use mixed-case letters", func(b bool) {})
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

		genBtn := widget.NewButton("Generate", gen)
		doneBtn := widget.NewButton("Done", func() {
			p.modal.popup.Hide()
		})

		var popup *widget.PopUp
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
				genBtn,
				doneBtn,
			),
		)
		popup = widget.NewModalPopUp(content, data.MainWindow.Canvas())

		modal := &generatePasswordModal{
			passEntry:         passEntry,
			numDigitsSelect:   numDigitsSelect,
			numSymbolsSelect:  numSymbolsSelect,
			passLengthSelect:  passLengthSelect,
			upperAlphaCheck:   upperAlphaCheck,
			allowRepeatsCheck: allowRepeatsCheck,
			popup:             popup,
			genBtn:            genBtn,
			doneBtn:           doneBtn,
		}

		p.modal = modal
	}
}

func generatePassword(length, numDigits, numSymbols int, useUpAlpha, allowRepeats bool) string {
	pass, err := password.Generate(length, numDigits, numSymbols, !useUpAlpha, allowRepeats)
	if err != nil {
		dialog.ShowError(err, data.MainWindow)
		return ""
	}
	return pass
}
