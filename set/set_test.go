package set

import (
	"fmt"
	"testing"
)

var tests_has = []struct {
	item     any
	expected any
}{
	{1, true},
	{2, true},
	{3, true},
	{5, false},
	{9, false},
}

func Test_Has(t *testing.T) {
	set := New(1, 2, 3)
	for _, test := range tests_has {
		fmt.Printf("Item: %v, Expected: %v, Received: %v\n", test.item, test.expected, set.Has(test.item))
		if set.Has(test.item) != test.expected {
			t.Errorf("Unexpected result. Expected: %v, Received: %v\n", test.expected, set.Has(test.item))
		}
	}
}

func Test_Size(t *testing.T) {
	set := New(1, 2, 3)
	fmt.Printf("Expected: %v, Received: %v\n", 3, set.Size())
	if set.Size() != 3 {
		t.Errorf("Wrong size. Expected: %v, Received: %v\n", 3, set.Size())
	}
}

func Test_Delete(t *testing.T) {
	set := New(1, 2, 3)
	for i := range set.Items() {
		set.Delete(i + 1)
		if set.Has(i + 1) {
			t.Errorf("Unexpected result: Expected: %v, Received: %v\n", false, set.Has(i+1))
		}
	}
}

func Test_Clear(t *testing.T) {
	set := New(1, 2, 3)
	set.Clear()
	if set.Size() != 0 {
		t.Errorf("Unexpected result: Expected: %v, Received: %v\n", 0, set.Size())
	}
}

func Test_Xor(t *testing.T) {
	set1 := New(1, 2, 3, 6, 9)
	set2 := New(1, 2, 3, 7, 8, 9)
	set3 := set1.Xor(set2)

	fmt.Println(set3)

	if set3.Size() != 3 {
		t.Errorf("Unexpected result: Expected: %v, Received: %v\n", 3, set3.Size())
	}

	for _, item := range []int{6, 7, 8} {
		fmt.Printf("Expected: %v, Received: %v\n", true, set3.Has(item))
		if !set3.Has(item) {
			t.Errorf("Unexpected result: Expected: %v, Received: %v\n", true, set3.Has(item))
		}
	}
}

func Test_Add(t *testing.T) {
	set := New(1, 2, 3)

	set.Add(1, 2)
	if set.Size() != 3 {
		t.Errorf("Unexpected result: Expected: %v, Received: %v\n", 3, set.Size())
	}

	set.Add(1, 2, 3)
	if set.Size() != 3 {
		t.Errorf("Unexpected result: Expected: %v, Received: %v\n", 3, set.Size())
	}

	set.Add(1, 2, 3, 4)
	if set.Size() != 4 {
		t.Errorf("Unexpected result: Expected: %v, Received: %v\n", 4, set.Size())
	}

	set.Add(5, 6)
	if set.Size() != 6 {
		t.Errorf("Unexpected result: Expected: %v, Received: %v\n", 6, set.Size())
	}

	set.Delete(5)
	set.Delete(6)

	if set.Size() != 4 {
		t.Errorf("Unexpected result: Expected: %v, Received: %v\n", 4, set.Size())
	}
}
