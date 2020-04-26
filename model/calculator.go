package model

type Calculator struct {
	result int
}

func (self *Calculator) Add(a int, b int) {
	self.result = a + b
}

func (self *Calculator) Subtract(a int, b int) {
	self.result = a - b
}

func (self *Calculator) SetResult(resultVal int) {
	self.result = resultVal
}

func (self *Calculator) GetResult() int {
	return self.result
}
