package main

import "testing"
import "reflect"

func TestSum(t *testing.T){
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1,2,3,4,5}

		got := Sum(numbers)
		want := 15

		if got != want{
			t.Errorf("got %d want %d give %c", got,want, numbers)	
		}
	})
	t.Run("collection of any size", func(t * testing.T){
		numbers := []int{1,2,3}

		got := Sum(numbers)
		want :=6


		if got != want{
			t.Errorf("got %d want %d give %c", got,want, numbers)	
		}
	})
}

func TestSumAll(t *testing.T){
	got := SumAll([]int{1,2},  []int{0,9})
	want := []int{3,9}


	if !reflect.DeepEqual(got, want){
		t.Errorf("got %v want %v)", got, want)
	}

}

func TestSumAllTail(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int){
		t.Helper()
		if !reflect.DeepEqual(got, want){
			t.Errorf("Got %v, want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T){
		got := SumAllTails([]int{1,2},[]int{0,9})
		want := []int{2,9}

		checkSums(t, got, want)

	})
	t.Run("Safely sum empty slices", func(t *testing.T){
		got := SumAllTails([]int{},[]int{0,9})
		want := []int{0,9}

		checkSums(t, got, want)
	})
}
