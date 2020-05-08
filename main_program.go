package main

import (
	"net/http"

	ctl "github.com/rachitt96/calculator_design/controller"
	mdl "github.com/rachitt96/calculator_design/model"
	view "github.com/rachitt96/calculator_design/view"
)

func main() {
	calcModel := mdl.Calculator{}

	operators := []view.OperatorRadioButton{
		view.OperatorRadioButton{"+", false, "+"},
		view.OperatorRadioButton{"-", false, "-"},
	}

	calcView := view.CalculatorView{Operator: operators}

	calcController := &ctl.CalculatorController{CalcModel: calcModel, CalcView: calcView}

	//http.HandleFunc("/", calcController.UpdateCalcView)
	http.HandleFunc("/submit", calcController.PerformOperation)
	http.HandleFunc("/", calcController.ResetCalculator)

	http.ListenAndServe(":9090", nil)

	/*mainLoop:
	for {

		fmt.Println("Enter your choise")
		fmt.Println("1: Next Calculation")
		fmt.Println("2: Exit Program")

		var choise string
		fmt.Scanln(&choise)

		switch choise {

		case "1":

			fmt.Println("Enter Number 1:")
			var operandOne string
			fmt.Scanln(&operandOne)

			fmt.Println("Enter Operator:")
			var operator string
			fmt.Scanln(&operator)

			fmt.Println("Enter Number 2:")
			var operandTwo string
			fmt.Scanln(&operandTwo)

			calcController.PerformOperation(operandOne, operator, operandTwo)

		case "2":
			break mainLoop
		}
	}*/

}
