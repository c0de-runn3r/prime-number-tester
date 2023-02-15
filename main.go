package main

import (
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.POST("/", handleRequest)

	e.Start(":8000")
}

func handleRequest(c echo.Context) error {
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Fatal(err)
	}
	reqBody := string(b)

	reqBody = strings.ReplaceAll(reqBody, " ", "")

	reqBody = reqBody[1 : len(reqBody)-1] // TODO make this more beatiful
	strSlc := strings.Split(reqBody, ",")
	intSlc := make([]int, len(strSlc))
	for i, s := range strSlc {
		intSlc[i], _ = strconv.Atoi(s) // TODO check errors
	}

	res := checkNumbers(intSlc)

	return c.JSON(http.StatusOK, res)
}

func checkNumbers(nums []int) []bool {
	res := make([]bool, len(nums))

	for i, v := range nums { // TODO make variable isPrime
		if isPrime(v) {
			res[i] = true
		}
	}

	return res
}

func isPrimeStd(num int) bool {
	return big.NewInt(int64(num)).ProbablyPrime(0)
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	sq_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq_root; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
