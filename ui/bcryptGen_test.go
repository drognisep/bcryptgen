package ui

import (
	"testing"

	"fyne.io/fyne/test"
	"github.com/drognisep/bcryptgen/data"
	testify "github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestBcryptGeneration(t *testing.T) {
	app := test.NewApp()
	win := app.NewWindow("testing")
	bcryptGen := NewBcryptField(win)
	win.SetContent(bcryptGen.Content())
	data.MainWindow = win

	data.Pass.SetState("1234")
	test.Tap(bcryptGen.generateButton)

	err := comparePassAndHash(t)
	if err != nil {
		t.Errorf("Pass and hash do not match: %v", err)
	}
}

func TestNoPasswordNoHash(t *testing.T) {
	assert := testify.New(t)
	app := test.NewApp()
	win := app.NewWindow("testing")
	bcryptGen := NewBcryptField(win)
	win.SetContent(bcryptGen.Content())
	data.MainWindow = win

	data.Pass.SetState("")
	test.Tap(bcryptGen.generateButton)

	assert.Equal("", data.Hash.GetState(), "Hash should not be generated with empty password")
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
