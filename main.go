package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Golculator")

	myWindow.Resize(fyne.NewSize(400, 400))

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")
	input.Resize(fyne.NewSize(380, 40))
	input.Move(fyne.NewPos(5, 0))

	btn_1 := widget.NewButton("1", func() {
		input.Append("1")
	})
	btn_1.Resize(fyne.NewSize(50, 50))
	btn_1.Move(fyne.NewPos(5, 50))

	btn_2 := widget.NewButton("2", func() {
		input.Append("2")
	})
	btn_2.Resize(fyne.NewSize(50, 50))
	btn_2.Move(fyne.NewPos(60, 50))

	btn_3 := widget.NewButton("3", func() {
		input.Append("3")
	})
	btn_3.Resize(fyne.NewSize(50, 50))
	btn_3.Move(fyne.NewPos(115, 50))

	myWindow.SetContent(container.NewWithoutLayout(input, btn_1, btn_2, btn_3))
	myWindow.ShowAndRun()
}