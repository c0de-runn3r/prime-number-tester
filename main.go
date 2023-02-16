package main

import (
	"log"
	"prime-number-tester/handlers"
	pnc "prime-number-tester/prime-number-checker"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := cleanenv.ReadEnv(&pnc.CFG)
	if err != nil {
		log.Println("no env set. using default values")
	}

	e := echo.New()
	e.HideBanner = true

	e.POST("/", handlers.HandleNumbersRequest)

	e.Logger.Fatal(e.Start(":8000"))
}
