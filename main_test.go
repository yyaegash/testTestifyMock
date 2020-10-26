package main

import (
	// "fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

type NumbersMock struct {
	mock.Mock
}

func (m *NumbersMock) addition() int {
	args := m.Called()
	return args.Int(0)
}

func TestCompute(t *testing.T) {
	addMock := new(NumbersMock)
	addMock.On("addition").Return(5)

	addMock.addition()

	res := compute()
	if res != 3 {
		t.Errorf("res = %v, expected 3\n", res)
	}
	addMock.AssertExpectations(t)
}

func TestMain(t *testing.T) {
	// testObj := new(NumbersMock)
	// testObj.On("addition").Return(10)
	// test := testObj.addition()
	// testObj.AssertExpectations(t)

	if false {
		t.Errorf("failing")
	}
}

// func TestAddition(t *testing.T) {
// 	tt := []struct {
// 		name string
// 		a    int
// 		b    int
// 		sum  int
// 	}{
// 		{"1 + 2", 1, 2, 3},
// 		{"1 - 1", 1, -1, 0},
// 	}

// 	for _, tc := range tt {
// 		t.Run(tc.name, func(t *testing.T) {
// 			res := addition(tc.a, tc.b)
// 			if res != tc.sum {
// 				t.Errorf("res = %v, expected was %v \n", res, tc.sum)
// 			}
// 		})
// 	}
// }
