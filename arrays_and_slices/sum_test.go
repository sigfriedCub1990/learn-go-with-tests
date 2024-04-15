package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got '%d' want '%d' given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		first := []int{1, 2, 3}
		second := []int{1, 2, 3, 4}

		got := SumAll(first, second)
		want := []int{6, 10}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
    checkSums := func(t testing.TB, got, want []int) {
        t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v'", got, want)
		}
    }

	t.Run("make the sums of some slices", func(t *testing.T) {
		first := []int{1, 2}
		second := []int{1, 9}

		got := SumAllTails(first, second)
		want := []int{2, 9}
        checkSums(t, got, want)
	})

	t.Run("safely sums empty slices", func(t *testing.T) {
		first := []int{1, 2}
		second := []int{}

		got := SumAllTails(first, second)
		want := []int{2, 0}
        checkSums(t, got, want)
	})

}
