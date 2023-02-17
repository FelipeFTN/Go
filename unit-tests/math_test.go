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
	assert.Equal(mathService.Add(), 4, "they should be equal")

	assert.NotEqual(mathService.Add(), 5, "they should not be equal")
}

func TestSubtract(t *testing.T) {
	assert := assert.New(t)
	math := new(MathContract)

	mathService := NewMathService(*math, 3, 1)
	assert.Equal(mathService.Subtract(), 2, "they should be equal")

	assert.NotEqual(mathService.Subtract(), 4, "they should be equal")
}

func TestDivide(t *testing.T) {
	assert := assert.New(t)
	math := new(MathContract)

	mathService := NewMathService(*math, 8, 2)
	assert.Equal(mathService.Divide(), 4, "they should be equal")

	assert.NotEqual(mathService.Divide(), 6, "they should be equal")
}

func TestMultiply(t *testing.T) {
	assert := assert.New(t)
	math := new(MathContract)

	mathService := NewMathService(*math, 2, 2)
	assert.Equal(mathService.Multiply(), 4, "they should be equal")

	assert.NotEqual(mathService.Multiply(), 8, "they should be equal")
}

func TestAddAndMultiply(t *testing.T) {
	assert := assert.New(t)
	mockRepo := new(mocks.MathContract)

	mockService := NewMathService(mockRepo, 1, 1)

	mockRepo.On("Multiply").Return(4)
	mockRepo.On("Add").Return(2)

	res := mockService.AddAndMultiply()

	mockRepo.AssertExpectations(t)

	assert.Equal(res, 4)
}
