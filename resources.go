package main

import ( 
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
)

var (
	Calc = widget.NewEntry()
	Btn_1 = widget.NewButton("1", func() {
		Calc.Append("1")
	})
	Btn_2 = widget.NewButton("2", func() {
		Calc.Append("2")
	})
	Btn_3 = widget.NewButton("3", func() {
		Calc.Append("3")
	})
	Btn_4 = widget.NewButton("4", func() {
		Calc.Append("4")
	})
	Btn_5 = widget.NewButton("5", func() {
		Calc.Append("5")
	})
	Btn_6 = widget.NewButton("6", func() {
		Calc.Append("6")
	})
	Btn_7 = widget.NewButton("7", func() {
		Calc.Append("7")
	})
	Btn_8 = widget.NewButton("8", func() {
		Calc.Append("8")
	})
	Btn_9 = widget.NewButton("9", func() {
		Calc.Append("9")
	})
	Btn_0 = widget.NewButton("0", func() {
		Calc.Append("0")
	})
	Btn_AC = widget.NewButton("AC", func() {
		Calc.SetText("")
	})
	Btn_Undo = widget.NewButton("AC", func() {
		Calc.Undo()
	})
)

func Sized() *fyne.Container {
	Calc.Resize(fyne.NewSize(380, 40))
	Calc.Move(fyne.NewPos(5, 0))

	Btn_1.Resize(fyne.NewSize(50, 50))
	Btn_1.Move(fyne.NewPos(5, 50))

	Btn_2.Resize(fyne.NewSize(50, 50))
	Btn_2.Move(fyne.NewPos(60, 50))

	Btn_3.Resize(fyne.NewSize(50, 50))
	Btn_3.Move(fyne.NewPos(115, 50))

	Btn_4.Resize(fyne.NewSize(50, 50))
	Btn_4.Move(fyne.NewPos(5, 50))

	Btn_5.Resize(fyne.NewSize(50, 50))
	Btn_5.Move(fyne.NewPos(60, 50))

	Btn_6.Resize(fyne.NewSize(50, 50))
	Btn_6.Move(fyne.NewPos(115, 50))

	Btn_7.Resize(fyne.NewSize(50, 50))
	Btn_7.Move(fyne.NewPos(5, 50))

	Btn_8.Resize(fyne.NewSize(50, 50))
	Btn_8.Move(fyne.NewPos(60, 50))

	Btn_9.Resize(fyne.NewSize(50, 50))
	Btn_9.Move(fyne.NewPos(115, 50))

	Btn_0.Resize(fyne.NewSize(50, 50))
	Btn_0.Move(fyne.NewPos(5, 50))

	Btn_AC.Resize(fyne.NewSize(50, 50))
	Btn_AC.Move(fyne.NewPos(5, 105))

	Btn_Undo.Resize(fyne.NewSize(100, 50))
	Btn_Undo.Move(fyne.NewPos(60, 105))

	return container.NewWithoutLayout(Calc, Btn_AC, Btn_Undo, Btn_1, Btn_2, Btn_3)
}