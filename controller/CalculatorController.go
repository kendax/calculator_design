package controller

import (
	"fmt"
	"strconv"

	"github.com/rachitt96/calculator_design/model"
)

type CalculatorController struct {
	CalcModel model.Calculator
}

func (self *CalculatorController) PerformOperation(firstOperand string, operator string, secondOperand string) {
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

}
