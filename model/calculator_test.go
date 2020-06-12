package model

import (
	"testing"
)

func TestAddition(t *testing.T) {
	calcObj := &Calculator{}
	calcObj.Add(5.5, 5)
	expectedResult := float64(10.5)

	if calcObj.result != expectedResult {
		t.Errorf("Sum was incorrect, got: %f, want: %f.", calcObj.result, expectedResult)
	}
}

func TestSubtraction(t *testing.T) {
	calcObj := &Calculator{}
	calcObj.Subtract(5.5, 5)
	expectedResult := float64(0.5)

	if calcObj.result != expectedResult {
		t.Errorf("Sum was incorrect, got: %f, want: %f.", calcObj.result, expectedResult)
	}
}

func TestMultiplication(t *testing.T) {
	calcObj := &Calculator{}
	calcObj.Multiply(5.5, 5)
	expectedResult := float64(27.5)

	if calcObj.result != expectedResult {
		t.Errorf("Sum was incorrect, got: %f, want: %f.", calcObj.result, expectedResult)
	}
}

func TestDivision(t *testing.T) {
	calcObj := &Calculator{}
	calcObj.Divide(5.5, 5)
	expectedResult := float64(1.1)

	if calcObj.result != expectedResult {
		t.Errorf("Sum was incorrect, got: %f, want: %f.", calcObj.result, expectedResult)
	}
}
