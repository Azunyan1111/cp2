package model

import (
	"github.com/labstack/echo"
	"net/http"
)

func LoginHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		return c.String(http.StatusOK, "Login URL")
	}
}
