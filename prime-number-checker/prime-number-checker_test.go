package primenumberchecker

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
	assert(t, CheckNumbers([]int{1, 2, 3, 4, 5}), []bool{false, true, true, false, true})
}

func BenchmarkStandartPrimeSmallRandomNumbers(b *testing.B) {
	for n := 0; n < b.N; n++ {
		isPrimeStd(n)
	}
}
func BenchmarkCustomPrimeSmallRandomNumbers(b *testing.B) {
	for n := 0; n < b.N; n++ {
		isPrime(n)
	}
}

func BenchmarkStandartPrimeBigNumber(b *testing.B) {
	for n := 0; n < b.N; n++ {
		isPrimeStd(2147483647)
	}
}
func BenchmarkCustomPrimeBigNumber(b *testing.B) {
	for n := 0; n < b.N; n++ {
		isPrime(2147483647)
	}
}
