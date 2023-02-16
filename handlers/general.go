package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	pnc "prime-number-tester/prime-number-checker"

	"github.com/labstack/echo/v4"
)

// for returning JSON error message for HTTP request with invalid parameters
type ErrorMessage struct {
	Error string `json:"error"`
}

func HandleNumbersRequest(c echo.Context) error {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Fatal(err)
	}
	reqBody := string(b)

	// remove all the whitespaces
	reqBody = strings.ReplaceAll(reqBody, " ", "")

	// remove [ ] brackets if they exist
	if reqBody[0] == '[' && reqBody[len(reqBody)-1] == ']' {
		reqBody = reqBody[1 : len(reqBody)-1]
	}

	strSlc := strings.Split(reqBody, ",")
	intSlc := make([]int, len(strSlc))
	for i, s := range strSlc {
		intSlc[i], err = strconv.Atoi(s)
		if err != nil {
			errMessage := ErrorMessage{Error: fmt.Sprintf("the given input is invalid. Element on index %x is not a number", i)}
			return c.JSON(http.StatusBadRequest, errMessage)
		}
	}

	res := pnc.CheckNumbers(intSlc)

	return c.JSON(http.StatusOK, res)
}
