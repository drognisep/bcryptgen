package ui

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// PasswordField creates a new password field ready to be used in the main UI
func PasswordField() *fyne.Container {
	genPassButton := widget.NewButton("Generate Password", generatePassword)
	passEntry := widget.NewPasswordEntry()
	field := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		widget.NewLabel("Password"),
		fyne.NewContainerWithLayout(
			&FieldLineLayout{},
			passEntry,
			genPassButton,
		),
	)
	return field
}

func generatePassword() {
	fmt.Println("'generatePassword' stub called")
}
