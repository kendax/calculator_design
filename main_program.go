package main

import (
	"fmt"

	ctl "github.com/rachitt96/calculator_design/controller"
	mdl "github.com/rachitt96/calculator_design/model"
)

func main() {
	calcModel := mdl.Calculator{}
	calcModel.SetResult(0)

	calcController := &ctl.CalculatorController{CalcModel: calcModel}

mainLoop:
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
	}

}
