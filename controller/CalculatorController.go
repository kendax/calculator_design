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

	for i := 0; i < len(self.CalcView.Operator); i++ {
		self.CalcView.Operator[i].IsChecked = false
	}

	self.CalcView.ShowValue = false

	self.CalcView.ShowCalc(w, r)
}

func contains(calcView view.CalculatorView, str string) bool {
	for _, a := range calcView.Operator {
		if a.Value == str {
			return true
		}
	}
	return false
}

func (self *CalculatorController) PerformOperation(w http.ResponseWriter, r *http.Request) {

	firstOperand := ""
	secondOperand := ""
	operator := ""

	if r.Method == http.MethodPost {
		r.ParseForm()
		firstOperand = r.FormValue("operand1")
		secondOperand = r.FormValue("operand2")
		operator = r.Form.Get("operator")
	}

	firstNumber, operandOneErr := strconv.ParseFloat(firstOperand, 64)
	secondNumber, operandTwoErr := strconv.ParseFloat(secondOperand, 64)

	fmt.Println(operandOneErr)

	if operandOneErr != nil || operandTwoErr != nil || contains(self.CalcView, operator) == false {
		if operandOneErr != nil {
			fmt.Println("in here")
			self.CalcView.ErrorMsg.FirstOperandError = "Enter Valid Number"
		}

		if operandTwoErr != nil {
			self.CalcView.ErrorMsg.SecondOperandError = "Enter Valid Number"
		}

		if contains(self.CalcView, operator) == false {
			self.CalcView.ErrorMsg.OperatorError = "Select an operator"
		}

		self.ResetCalculator(w, r)
		return
	} else {
		self.CalcView.ErrorMsg.FirstOperandError = ""
		self.CalcView.ErrorMsg.SecondOperandError = ""
		self.CalcView.ErrorMsg.OperatorError = ""
	}

	fmt.Println()
	fmt.Println("Requested Operation:", firstNumber, operator, secondNumber)

	switch operator {
	case "+":
		self.CalcModel.Add(firstNumber, secondNumber)

	case "-":
		self.CalcModel.Subtract(firstNumber, secondNumber)

	case "*":
		self.CalcModel.Multiply(firstNumber, secondNumber)

	case "/":
		self.CalcModel.Divide(firstNumber, secondNumber)

	default:
		http.Error(w, "operator not available", 500)
		return
	}

	finalAnswer := self.CalcModel.GetResult()
	fmt.Println("Result after operation", finalAnswer)
	fmt.Println()

	self.CalcView.FirstOperand = firstNumber
	self.CalcView.SecondOperand = secondNumber
	self.CalcView.Result = self.CalcModel.GetResult()
	self.CalcView.ShowValue = true

	for i := 0; i < len(self.CalcView.Operator); i++ {
		current_obj := &self.CalcView.Operator[i]

		if current_obj.Value == operator {
			current_obj.IsChecked = true
		}
	}

	self.CalcView.ShowCalc(w, r)

}
