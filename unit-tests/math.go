package math

type MathContract interface {
	Add() int
	Subtract() int
	Divide() int
	Multiply() int
	AddAndMultiply() int
}

type MathType struct {
	A            int
	B            int
	MathContract MathContract
}

func NewMathService(mathContract MathContract, a int, b int) *MathType {
	return &MathType{
		a,
		b,
		mathContract,
	}
}

func (m *MathType) Add() int {
	return m.A + m.B
}

func (m *MathType) Subtract() int {
	return m.A - m.B
}

func (m *MathType) Divide() int {
	return m.A / m.B
}

func (m *MathType) Multiply() int {
	return m.A * m.B
}

func (m *MathType) AddAndMultiply() int {

	m.A = m.MathContract.Add()
	m.B = m.MathContract.Add()

	result := m.MathContract.Multiply()

	return result
}
