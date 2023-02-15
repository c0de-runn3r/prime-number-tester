package main

import (
	"reflect"
	"testing"
)

func assert(t *testing.T, a, b any) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%+v != %+v", a, b)
	}
}

func TestCheckNumbers(t *testing.T) {
	assert(t, checkNumbers([]int{1, 2, 3, 4, 5}), []bool{false, true, true, false, true})
}
