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
	assert(t, CheckNumbers([]int{11, 13, 17, 19, 23}), []bool{true, true, true, true, true})
	assert(t, CheckNumbers([]int{100, 1000, 10000, 100000, 1000000}), []bool{false, false, false, false, false})
	assert(t, CheckNumbers([]int{0, -1, -2, -3, -4, -5}), []bool{false, false, false, false, false, false})
}

func TestIsPrime(t *testing.T) {
	assert(t, isPrime(-10), false)
	assert(t, isPrime(0), false)
	assert(t, isPrime(1), false)
	assert(t, isPrime(2), true)
	assert(t, isPrime(3), true)
	assert(t, isPrime(4), false)
	assert(t, isPrime(5), true)
}

func TestIsPrimeStd(t *testing.T) {
	assert(t, isPrime(-10), false)
	assert(t, isPrimeStd(0), false)
	assert(t, isPrimeStd(1), false)
	assert(t, isPrimeStd(2), true)
	assert(t, isPrimeStd(3), true)
	assert(t, isPrimeStd(4), false)
	assert(t, isPrimeStd(5), true)
}

// These benchmarks shows the efficienty of my own prime number checker
// comparing to one from big package on small values
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

// These benchmarks shows the efficienty of prime number checker from big package
// comparing to custom one if we are processing big values
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
