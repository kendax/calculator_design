package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/rachitt96/calculator_design/model"
	"github.com/rachitt96/calculator_design/view"
)

type CalculatorController struct {
	CalcModel model.Calculator
	CalcView  view.CalculatorView
}

/*func (self *CalculatorController) UpdateCalcView(w http.ResponseWriter, r *http.Request) {
	finalResult := self.CalcModel.GetResult()

	self.CalcView.Result = finalResult
	self.CalcView.ShowCalc(w, r)
}*/

func (self *CalculatorController) ResetCalculator(w http.ResponseWriter, r *http.Request) {
	self.CalcModel.SetResult(0)
	self.CalcView.FirstOperand = 0
	self.CalcView.SecondOperand = 0
	self.CalcView.Result = 0
	self.CalcView.Operator = ""

	self.CalcView.ShowCalc(w, r)
}

func (self *CalculatorController) PerformOperation(w http.ResponseWriter, r *http.Request) {

	firstOperand := ""
	secondOperand := ""
	operator := ""

	if r.Method == http.MethodPost {
		firstOperand = r.FormValue("operand1")
		secondOperand = r.FormValue("operand2")
		operator = r.FormValue("operator")
	}

	firstNumber, _ := strconv.Atoi(firstOperand)
	secondNumber, _ := strconv.Atoi(secondOperand)

	fmt.Println()
	fmt.Println("Requested Operation:", firstNumber, operator, secondNumber)

	switch operator {
	case "+":
		self.CalcModel.Add(firstNumber, secondNumber)

	case "-":
		self.CalcModel.Subtract(firstNumber, secondNumber)
	}

	finalAnswer := self.CalcModel.GetResult()
	fmt.Println("Result after operation", finalAnswer)
	fmt.Println()

	self.CalcView.FirstOperand = firstNumber
	self.CalcView.SecondOperand = secondNumber
	self.CalcView.Operator = operator
	self.CalcView.Result = self.CalcModel.GetResult()

	self.CalcView.ShowCalc(w, r)

}
