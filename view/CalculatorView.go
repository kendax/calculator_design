package view

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ValidationMessage struct {
	FirstOperandError  string
	SecondOperandError string
	OperatorError      string
}

type OperatorRadioButton struct {
	Value     string
	IsChecked bool
	Text      string
}

type CalculatorView struct {
	Result        float64
	FirstOperand  float64
	SecondOperand float64
	Operator      []OperatorRadioButton
	ShowValue     bool
	ErrorMsg      ValidationMessage
}

/*func (self *CalculatorView) submitButton(w http.ResponseWriter, r *http.Request) *CalculatorView {

	if r.Method == http.MethodPost {
		self.FirstOperand = r.FormValue("operand1")
		self.SecondOperand = r.FormValue("operand2")
		self.Operator = r.FormValue("operator")
	}

	return self
}*/

func (self *CalculatorView) ShowCalc(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("view/calculator_view.html")

	if err != nil {
		log.Print("template parsing error: ", err)
	}

	fmt.Println(self)

	err = t.Execute(w, self)

	if err != nil {
		log.Print("template executing error: ", err)
	}
}
