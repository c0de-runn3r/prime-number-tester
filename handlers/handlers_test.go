package handlers

import (
	"reflect"
	"testing"
)

func assert(t *testing.T, a, b any) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%+v != %+v", a, b)
	}
}

func TestReqBodyToIntSlice(t *testing.T) {
	slc, err := reqBodyToIntSlice([]byte("[1,2,3,4,5]"))
	assert(t, slc, []int{1, 2, 3, 4, 5})
	assert(t, err.Error == "", true)

	slc, err = reqBodyToIntSlice([]byte("[-1,-2,-3,-4,-5]"))
	assert(t, slc, []int{-1, -2, -3, -4, -5})
	assert(t, err.Error == "", true)

	slc, err = reqBodyToIntSlice([]byte("[1  ,      2,3 , 4,5]"))
	assert(t, slc, []int{1, 2, 3, 4, 5})
	assert(t, err.Error == "", true)

	slc, err = reqBodyToIntSlice([]byte("1,2,3,4,5"))
	assert(t, slc, []int{1, 2, 3, 4, 5})
	assert(t, err.Error == "", true)

	slc, err = reqBodyToIntSlice([]byte("1  ,      2,3 , 4,5"))
	assert(t, slc, []int{1, 2, 3, 4, 5})
	assert(t, err.Error == "", true)

	_, err = reqBodyToIntSlice([]byte("[a1,2,3,4,5]"))
	assert(t, err.Error == "", false)

	_, err = reqBodyToIntSlice([]byte("[a1,sdadasdad2, ,4,5]"))
	assert(t, err.Error == "", false)

	_, err = reqBodyToIntSlice([]byte("1,2,3,4,5]"))
	assert(t, err.Error == "", false)

	_, err = reqBodyToIntSlice([]byte("[1,2,3,4,5"))
	assert(t, err.Error == "", false)
}
