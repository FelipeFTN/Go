package math

import (
	"testing"

	"github.com/FelipeFTN/unit-tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	math := new(MathContract)

	mathService := NewMathService(*math, 3, 1)
	assert.Equal(4, mathService.Add(), "they should be equal")

	assert.NotEqual(5, mathService.Add(), "they should not be equal")
}

func TestSubtract(t *testing.T) {
	assert := assert.New(t)
	math := new(MathContract)

	mathService := NewMathService(*math, 3, 1)
	assert.Equal(2, mathService.Subtract(), "they should be equal")

	assert.NotEqual(4, mathService.Subtract(), "they should be equal")
}

func TestDivide(t *testing.T) {
	assert := assert.New(t)
	math := new(MathContract)

	mathService := NewMathService(*math, 8, 2)
	assert.Equal(4, mathService.Divide(), "they should be equal")

	assert.NotEqual(6, mathService.Divide(), "they should be equal")
}

func TestMultiply(t *testing.T) {
	assert := assert.New(t)
	math := new(MathContract)

	mathService := NewMathService(*math, 2, 2)
	assert.Equal(4, mathService.Multiply(), "they should be equal")

	assert.NotEqual(8, mathService.Multiply(), "they should be equal")
}

func TestAddAndMultiply(t *testing.T) {
	assert := assert.New(t)
	mockRepo := new(mocks.MathContract)

	mockService := NewMathService(mockRepo, 1, 1)

	mockRepo.On("Multiply").Return(4)
	mockRepo.On("Add").Return(2)

	res := mockService.AddAndMultiply()

	mockRepo.AssertExpectations(t)

	assert.Equal(4, res)
}
