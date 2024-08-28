package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Golculator")

	myWindow.Resize(fyne.NewSize(233, 330))

	cont := Sized()

	myWindow.SetContent(cont)
	myWindow.ShowAndRun()
}