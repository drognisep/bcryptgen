package ui

import (
	"testing"

	"fyne.io/fyne/test"
	"github.com/drognisep/bcryptgen/data"
	"golang.org/x/crypto/bcrypt"
)

func TestBcryptGeneration(t *testing.T) {
	app := test.NewApp()
	win := app.NewWindow("testing")
	bcryptGen := NewBcryptField()
	win.SetContent(bcryptGen.Content())
	data.MainWindow = win

	data.Pass.SetState("1234")
	test.Tap(bcryptGen.generateButton)

	err := comparePassAndHash(t)
	if err != nil {
		t.Errorf("Pass and hash do not match: %v", err)
	}
}

func comparePassAndHash(t *testing.T) error {
	pass := data.Pass.GetState()
	hash := data.Hash.GetState()
	if pass == "" {
		t.Errorf("Password state is blank")
	}
	if hash == "" {
		t.Errorf("Hash state is blank")
	}
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
}
