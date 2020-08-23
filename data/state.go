package data

import "fyne.io/fyne"

// Pass is the observable state of the password entered by the user or generated.
var Pass *StringSubject = NewStringSubject("")

// Hash is the observable state of the bcrypt hash generated from the password.
var Hash *StringSubject = NewStringSubject("")

// MainWindow is held as global state to enable easy reference with dialogs.
var MainWindow fyne.Window
