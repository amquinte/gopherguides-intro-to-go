package week08

import (
	"errors"
	"testing"
)

func Test_CP_TDT(t *testing.T) {
	t.Parallel()
	table := []struct {
		name string
		cp   CompletedProduct
		err  error
	}{
		{name: "invalidEmployee", cp: CompletedProduct{Employee: -1}, err: ErrInvalidEmployee(-1)},
		{name: "invalidMaterials", cp: CompletedProduct{Employee: 1}, err: ErrInvalidMaterials(0)},
		{name: "validCP", cp: CompletedProduct{Product: Product{Materials: Materials{
			Wood: 2,
			Oil:  3,
		}, builtBy: Employee(1)}, Employee: 1}, err: nil},
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
