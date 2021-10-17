package main

import (
	"testing"
)

func TestArray(t *testing.T) {
	exp := [4]int{1, 2, 3, 4}
	act := [4]int{}

	//Copy contents of exp into act
	for i, n := range exp {
		act[i] = n
	}

	//Test to see if the length of act and exp are the same
	if len(exp) != len(act) {
		t.Errorf("Expected lengths to be the same but they are different")
	}

	//Iterate through act and assert that both arrays are equal
	for i, n := range act {
		result := exp[i]

		//Compares values for equality
		if result != n {
			t.Errorf("Expected %d but got %d", n, result)
		}
	}
}

func TestSlice(t *testing.T) {
	exp := []int{1, 2, 3, 4}
	act := []int{}

	//Copy content of exp into act
	for _, n := range exp {
		act = append(act, n)
	}

	//Test to see if the length of act and exp are the same
	if len(exp) != len(act) {
		t.Errorf("Expected lengths to be the same but they are different")
	}

	//Iterate through act and assert that the contents of both slices are equal
	for i, n := range act {
		result := exp[i]

		//Compares values for equality
		if result != n {
			t.Errorf("Expected %d but got %d", n, result)
		}
	}
}

func TestMap(t *testing.T) {
	exp := map[string]int{
		"John":     10,
		"Samantha": 12,
		"Bob":      15,
		"Stacy":    17,
	}
	act := map[string]int{}

	//Copy content of exp into act
	for i, n := range exp {
		act[i] = n
	}

	//Test to see if the length of act and exp are the same
	if len(exp) != len(act) {
		t.Errorf("Expected lengths of the maps to be the same but they are not")
	}

	//Iterate through act and assert that the contents of both maps are equal
	for k := range act {
		expected := exp[k]
		actual, ok := act[k]

		//Checks if key exists in the map
		if !ok {
			t.Errorf("Key not found: %s", k)
		}

		//Compares values of both maps to see if they match
		if actual != expected {
			t.Errorf("Expected %d but got %d", expected, actual)
		}
	}
}
