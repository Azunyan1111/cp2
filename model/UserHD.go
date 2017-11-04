package model

import (
	"github.com/labstack/echo"
	"net/http"
)

func GetUserHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		return c.String(http.StatusOK, "You name is " + c.Param("userId"))
	}
}

func SetUserHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		return c.String(http.StatusOK, "You name is " + c.Param("userId"))
	}
}