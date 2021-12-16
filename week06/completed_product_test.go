package week06

import (
	"errors"
	"testing"
)

// func Test_CP_IsValid(t *testing.T) {
// 	e := -1
// 	cp := CompletedProduct{Employee: Employee(e)}
// 	act := cp.IsValid()
// 	exp := ErrInvalidEmployee(e)
// 	if errors.Is(act, exp) == false {
// 		t.Fatalf("expect %v but got %v instead", exp, act)
// 	}
// }

// func Test_CP_IsValid2(t *testing.T) {
// 	e := 1
// 	cp := CompletedProduct{Employee: Employee(e)}
// 	act := cp.IsValid()
// 	exp := ErrInvalidQuantity(0)
// 	if errors.Is(act, exp) == false {
// 		t.Fatalf("expect %v but got %v instead", exp, act)
// 	}
// }

// func Test_CP_IsValid3(t *testing.T) {
// 	cp := CompletedProduct{Product: Product{Quantity: 1, builtBy: Employee(1)}, Employee: 1}
// 	act := cp.IsValid()

// 	if act != nil {
// 		t.Fatalf("expect nil but got %v instead", act)
// 	}
// }

//This acomplishes the same tests as the ones above
//Just re-wrote the tests using TDT
func Test_CP_TDT(t *testing.T) {
	t.Parallel()
	table := []struct {
		name string
		cp   CompletedProduct
		err  error
	}{
		{name: "invalidEmployee", cp: CompletedProduct{Employee: -1}, err: ErrInvalidEmployee(-1)},
		{name: "invalidQuantity", cp: CompletedProduct{Employee: 1}, err: ErrInvalidQuantity(0)},
		{name: "validCP", cp: CompletedProduct{Product: Product{Quantity: 1, builtBy: Employee(1)}, Employee: 1}, err: nil},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.cp.IsValid()

			if errors.Is(tt.err, act) == false {
				t.Fatalf("expect %v but got %v instead", tt.err, act)
			}
		})
	}
}
