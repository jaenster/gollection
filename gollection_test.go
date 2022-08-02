package gollection

import (
	"testing"
)

type TestStruct struct {
	X      int
	Y      int
	signal chan struct{}
}

func TestPack(t *testing.T) {

	t.Run("Test gollection", func(t *testing.T) {
		gollection := New[TestStruct]()

		one := &TestStruct{1, 1, make(chan struct{})}
		two := &TestStruct{2, 2, make(chan struct{})}
		three := &TestStruct{3, 3, make(chan struct{})}
		four := &TestStruct{4, 4, make(chan struct{})}

		gollection.Add(one)
		if gollection.Has(one) != true ||
			gollection.Has(two) != false ||
			gollection.Size() != 1 {
			t.Fail()
		}

		gollection.Add(two)
		if gollection.Has(one) != true ||
			gollection.Has(two) != true ||
			gollection.Size() != 2 {
			t.Fail()
		}

		// Test remove first element
		gollection.Remove(one)
		if gollection.Has(one) != false ||
			gollection.Has(two) != true ||
			gollection.Size() != 1 {
			t.Fail()
		}

		// test remove only element
		gollection.Remove(two)
		if gollection.Has(one) != false ||
			gollection.Has(two) != false ||
			gollection.Size() != 0 {
			t.Fail()
		}

		gollection.Add(one)
		gollection.Add(two)
		gollection.Add(three)
		gollection.Add(four)
		if gollection.Has(one) != true ||
			gollection.Has(two) != true ||
			gollection.Has(three) != true ||
			gollection.Has(four) != true ||
			gollection.Size() != 4 {
			t.Fail()
		}

		// test remove middle element
		gollection.Remove(two)
		if gollection.Has(one) != true ||
			gollection.Has(two) != false ||
			gollection.Has(three) != true ||
			gollection.Has(four) != true ||
			gollection.Size() != 3 {
			t.Fail()
		}

		gollection.Remove(four)
		if gollection.Has(one) != true ||
			gollection.Has(two) != false ||
			gollection.Has(three) != true ||
			gollection.Has(four) != false ||
			gollection.Size() != 2 {
			t.Fail()
		}

		gollection.ForEach(func(other *TestStruct) {
			if gollection.Has(other) != true ||
				(other != one && other != three) {
				t.Fail()
			}
		})

		done := make(chan struct{})

		go (func() {

			gollection.GoEach(func(other *TestStruct) {
				other.signal <- struct{}{}
			})

			gollection.ForEach(func(other *TestStruct) {
				<-other.signal
			})

			close(done)
		})()

		<-done
	})
}
