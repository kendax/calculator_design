package model

type Calculator struct {
	result float64
}

func (self *Calculator) Add(a float64, b float64) {
	self.result = a + b
}

func (self *Calculator) Subtract(a float64, b float64) {
	self.result = a - b
}

func (self *Calculator) Multiply(a float64, b float64) {
	self.result = a * b
}

func (self *Calculator) Divide(a float64, b float64) {
	self.result = a / b
}

func (self *Calculator) SetResult(resultVal float64) {
	self.result = resultVal
}

func (self *Calculator) GetResult() float64 {
	return self.result
}
