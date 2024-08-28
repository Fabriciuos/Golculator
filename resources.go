package main

import ( 
	"strconv"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"github.com/Knetic/govaluate"
)

var (
	res = ""
	Calc = widget.NewEntry()
	Btn_1 = widget.NewButton("1", func() {
		Calc.Append("1")
		res += "1"
	})
	Btn_2 = widget.NewButton("2", func() {
		Calc.Append("2")
		res += "2"
	})
	Btn_3 = widget.NewButton("3", func() {
		Calc.Append("3")
		res += "3"
	})
	Btn_4 = widget.NewButton("4", func() {
		Calc.Append("4")
		res += "4"
	})
	Btn_5 = widget.NewButton("5", func() {
		Calc.Append("5")
		res += "5"
	})
	Btn_6 = widget.NewButton("6", func() {
		Calc.Append("6")
		res += "6"
	})
	Btn_7 = widget.NewButton("7", func() {
		Calc.Append("7")
		res += "7"
	})
	Btn_8 = widget.NewButton("8", func() {
		Calc.Append("8")
		res += "8"
	})
	Btn_9 = widget.NewButton("9", func() {
		Calc.Append("9")
		res += "9"
	})
	Btn_0 = widget.NewButton("0", func() {
		Calc.Append("0")
		res += "0"
	})
	Btn_AC = widget.NewButton("AC", func() {
		Calc.SetText("")
		res = ""
	})
	Btn_Undo = widget.NewButton("⌫", func() {
		if len(Calc.Text) >= 1 {
		Calc.SetText(Calc.Text[:len(Calc.Text)-1])
		res = res[:len(res)-1]
	}
	})
	Btn_M = widget.NewButton("×", func() {
		Calc.Append("×")
		res += "*"
	})
	Btn_D = widget.NewButton("÷", func() {
		Calc.Append("÷")
		res += "/"
	})
	Btn_I = widget.NewButton("+", func() {
		Calc.Append("+")
		res += "+"
	})
	Btn_L = widget.NewButton("−", func() {
		Calc.Append("−")
		res += "-"
	})
	Btn_Z = widget.NewButton(".", func() {
		Calc.Append(".")
		res += "."
	})
	Btn_K = widget.NewButton("√", func() {
		l, err := strconv.ParseFloat(res, 64)
		if err != nil {
			Calc.SetText("ERROR")
		}
		res = strconv.FormatFloat(math.Sqrt(l), 'g', -1, 64)
		Calc.SetText(res)
		
	})
	Btn_P = widget.NewButton("x²", func() {
		l, err := strconv.ParseFloat(res, 64)
		if err != nil {
			Calc.SetText("ERROR")
		}
		res = strconv.FormatFloat(math.Pow(l, 2), 'g', -1, 64)
		Calc.SetText(res)
	})
    Btn_Equal = widget.NewButton("=", func() {
        expression, err := govaluate.NewEvaluableExpression(res)
        if err == nil {
            result, err := expression.Evaluate(nil)
            if err == nil {
                res = strconv.FormatFloat(result.(float64), 'f', -1, 64)
                Calc.SetText(strconv.FormatFloat(result.(float64), 'f', -1, 64))
            } else {
                res = err.Error()
                Calc.SetText("ERROR")
            }
        } else {
            res = err.Error()
            Calc.SetText("ERROR")
        }
    })
)

func Sized() *fyne.Container {
	Calc.Resize(fyne.NewSize(217, 40))
	Calc.Move(fyne.NewPos(5, 0))

	Btn_1.Resize(fyne.NewSize(50, 50))
	Btn_1.Move(fyne.NewPos(5, 215))

	Btn_2.Resize(fyne.NewSize(50, 50))
	Btn_2.Move(fyne.NewPos(60, 215))

	Btn_3.Resize(fyne.NewSize(50, 50))
	Btn_3.Move(fyne.NewPos(115, 215))

	Btn_4.Resize(fyne.NewSize(50, 50))
	Btn_4.Move(fyne.NewPos(5, 160))

	Btn_5.Resize(fyne.NewSize(50, 50))
	Btn_5.Move(fyne.NewPos(60, 160))

	Btn_6.Resize(fyne.NewSize(50, 50))
	Btn_6.Move(fyne.NewPos(115, 160))

	Btn_7.Resize(fyne.NewSize(50, 50))
	Btn_7.Move(fyne.NewPos(5, 105))

	Btn_8.Resize(fyne.NewSize(50, 50))
	Btn_8.Move(fyne.NewPos(60, 105))

	Btn_9.Resize(fyne.NewSize(50, 50))
	Btn_9.Move(fyne.NewPos(115, 105))

	Btn_0.Resize(fyne.NewSize(105, 50))
	Btn_0.Move(fyne.NewPos(5, 270))

	Btn_AC.Resize(fyne.NewSize(50, 50))
	Btn_AC.Move(fyne.NewPos(5, 50))

	Btn_Undo.Resize(fyne.NewSize(100, 50))
	Btn_Undo.Move(fyne.NewPos(60, 105))

	Btn_M.Resize(fyne.NewSize(50, 50))
	Btn_M.Move(fyne.NewPos(170, 105))

	Btn_D.Resize(fyne.NewSize(50, 50))
	Btn_D.Move(fyne.NewPos(170, 50))

	Btn_I.Resize(fyne.NewSize(50, 50))
	Btn_I.Move(fyne.NewPos(170, 215))

	Btn_L.Resize(fyne.NewSize(50, 50))
	Btn_L.Move(fyne.NewPos(170, 160))

	Btn_Z.Resize(fyne.NewSize(50, 50))
	Btn_Z.Move(fyne.NewPos(115, 270))

	Btn_K.Resize(fyne.NewSize(50, 50))
	Btn_K.Move(fyne.NewPos(60, 50))

	Btn_P.Resize(fyne.NewSize(50, 50))
	Btn_P.Move(fyne.NewPos(115, 50))

	Btn_Equal.Resize(fyne.NewSize(50, 50))
	Btn_Equal.Move(fyne.NewPos(170, 270))

	return container.NewWithoutLayout(Calc, Btn_AC, Btn_K, Btn_P, Btn_D, Btn_7, Btn_8, Btn_9, Btn_M, Btn_4, Btn_5, Btn_6, Btn_L, Btn_1, Btn_2, Btn_3, Btn_I, Btn_0, Btn_Z, Btn_Equal)
}