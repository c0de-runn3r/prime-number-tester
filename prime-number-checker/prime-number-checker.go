package primenumberchecker

import (
	"math"
	"math/big"
)

var CFG Config

type Config struct {
	CustomPrimeChecker bool `env:"CUSTOM_PRIME_CHECKER" env-default:"true"`
}

// Function takes slice of integers, checks every int is it prime and returns
// slice of booleans (true for prime number and false for not).
func CheckNumbers(nums []int) []bool {
	res := make([]bool, len(nums))

	for i, v := range nums {
		switch CFG.CustomPrimeChecker { // by default uses custom prime number checker - faster for small values.
		case true:
			if isPrime(v) {
				res[i] = true
			}
		case false:
			if isPrimeStd(v) {
				res[i] = true
			}
		}
	}
	return res
}

// Standart prime number checker from 'big' package. Uses Miller-Rabin test with n pseudorandomly chosen bases
// and a Baillie-PSW test. Test is 100% accurate for inputs less than 2⁶⁴, then ¼ⁿ probability to make mistake.
func isPrimeStd(num int) bool {
	return big.NewInt(int64(num)).ProbablyPrime(0)
}

// Custom simple prime number checker. 100% accurate, faster than standart one for small numbers, but slower for big ones.
// Uses 6k +/- 1 optimization. The most fast from the simple and 100% accurate prime number checkers.
func isPrime(num int) bool {
	if num <= 3 {
		return num > 1
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	limit := int(math.Pow(float64(num), 0.5))

	for i := 5; i <= limit+1; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}
