package view

import (
	"html/template"
	"log"
	"net/http"
)

type CalculatorView struct {
	Result        int
	FirstOperand  int
	SecondOperand int
	Operator      string
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
	//calcView := &CalculatorView{Result: 0}

	t, err := template.ParseFiles("view/calculator_view.html")

	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, self)

	if err != nil {
		log.Print("template executing error: ", err)
	}
}
