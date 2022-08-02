package main

import (
	"fmt"
	"math"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var paramAndSequence = map[string]int{
	"x": 2,
	"y": 5,
	"z": 6,
}

func calculate(param string) int {
	n := paramAndSequence[param]
	return int((math.Pow(float64(n), 3) + 3*math.Pow(float64(n), 2) - 4*float64(n) + 6) / 6)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) (err error) {
		return c.String(http.StatusOK, "health check !")
	})
	e.GET("/find/:param", func(c echo.Context) (err error) {
		param := strings.ToLower(c.Param("param"))
		if _, ok := paramAndSequence[param]; ok {
			result := calculate(param)
			res := fmt.Sprintf("value of %s is %d", param, result)
			return c.String(http.StatusOK, res)
		}
		return c.String(http.StatusBadRequest, "request must be x, y or z")
	})

	e.Logger.Fatal(e.Start(":9000"))
}
