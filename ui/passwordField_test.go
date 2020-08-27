package ui

import (
	"testing"

	"fyne.io/fyne"
	"fyne.io/fyne/test"
	"github.com/drognisep/bcryptgen/data"
)

var app fyne.App
var win fyne.Window
var pass *PasswordComponent

func setup() {
	data.Pass = data.NewStringSubject("")
	data.Hash = data.NewStringSubject("")
	data.MainWindow = nil
	app = test.NewApp()
	win = app.NewWindow("test window")
	data.MainWindow = win

	pass = NewPasswordField()
	win.SetContent(pass.Content())
}

func cleanup() {
	app.Quit()
}

func TestPasswordGeneration(t *testing.T) {
	setup()
	defer cleanup()
	var passText string

	pass.showModal()
	modal := pass.modal
	passText = modal.passEntry.Text
	test.Tap(modal.genBtn)
	newPassText := modal.passEntry.Text
	if newPassText == passText {
		t.Errorf("Pass text entries should be different")
	}
	if newPassText == "" {
		t.Errorf("Pass text should not be empty")
	}
	if data.Pass.GetState() != newPassText {
		t.Errorf("State should be propagated to data.Pass state")
	}
}

func TestPasswordNotEmpty(t *testing.T) {
	setup()
	defer cleanup()

	pass.showModal()
	modal := pass.modal
	test.Tap(modal.genBtn)
	newPassText := modal.passEntry.Text
	if newPassText == "" {
		t.Errorf("Pass text should not be empty")
	}
}

func TestPasswordState(t *testing.T) {
	setup()
	defer cleanup()

	pass.showModal()
	modal := pass.modal
	test.Tap(modal.genBtn)
	newPassText := modal.passEntry.Text
	if data.Pass.GetState() != newPassText {
		t.Errorf("State should be propagated to data.Pass state")
	}
}
