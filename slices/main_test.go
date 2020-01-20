package slices

import (
	"testing"
)

func TestElements_Delete(t *testing.T) {
	tests := []struct {
		name string
		es   Elements
	}{
		{
			name: "empty slice",
			es:   []Element{},
		}, {
			name: "one element",
			es: []Element{
				"one",
			},
		}, {
			name: "many elements",
			es: []Element{
				"one",
				"two",
				"three",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r, err := test.es.Delete(0)
			if err != nil {
				t.Error(err)
			}
			if len(r) != len(test.es)-1 && len(test.es) > 0 {
				t.Error()
			}
			if len(test.es)-1 > 0 && r[0] == "one" {
				t.Error()
			}
		})
	}
}

func TestElements_Insert(t *testing.T) {
	tests := []struct {
		name        string
		es          Elements
		insertKey   int
		insertValue Element
	}{
		{
			name:        "insert in empty slice on zero index",
			es:          []Element{},
			insertKey:   0,
			insertValue: "insertedValue",
		}, {
			name:        "insert in empty slice on greater then zero index",
			es:          []Element{},
			insertKey:   2,
			insertValue: "insertedValue",
		}, {
			name: "insert in non empty slice",
			es: []Element{
				"one",
				"two",
				"three",
			},
			insertKey:   0,
			insertValue: "insertedValue",
		}, {
			name: "insert in non empty slice",
			es: []Element{
				"one",
				"two",
				"three",
			},
			insertKey:   1,
			insertValue: "insertedValue",
		}, {
			name: "insert in non empty slice",
			es: []Element{
				"one",
				"two",
				"three",
			},
			insertKey:   2,
			insertValue: "insertedValue",
		}, {
			name: "insert in non empty slice",
			es: []Element{
				"one",
				"two",
				"three",
			},
			insertKey:   -1,
			insertValue: "insertedValue",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			originalLen := len(test.es)

			result := test.es.Insert(test.insertKey, test.insertValue)

			if originalLen+1 != len(result) {
				t.Error()
			}

			var insertId int
			if test.insertKey > len(test.es) || test.insertKey < 0 {
				insertId = len(test.es)
			} else {
				insertId = test.insertKey
			}

			if result[insertId] != test.insertValue {
				t.Error()
			}
		})
	}
}

func TestSlice(t *testing.T) {
	foo := func(in []string) {
		if len(in) <= 0 {
			return
		}
		key := 0

		if key < len(in)-1 {
			copy(in[key:], in[key+1:])
		}
		//in[len(in)-1] = ""

		//in = in[:len(in)-1]
	}

	//in := []string{"first", "second"}
	in := []string{"first"}
	//out := foo(in)
	foo(in)

	t.Log(in)
	//t.Log(out)
	t.Log(len(in))
	//t.Log(len(out))
	t.Log("")
}
