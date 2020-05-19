package ch6

import "testing"

type IntList struct {
	Val  int
	Tail *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Val + list.Tail.Sum()
}

func TestSum(t *testing.T) {
	var testdata = []struct {
		list IntList
		exp  int
	}{
		{
			IntList{10,
				&IntList{5,
					&IntList{3, nil},
				},
			},
			18,
		},
		{
			IntList{10,
				&IntList{5,
					&IntList{3, nil},
				},
			},
			19,
		},
	}

	for _, d := range testdata {
		r := d.list.Sum()
		if r != d.exp {
			t.Error("test failed")
		}
	}
}
