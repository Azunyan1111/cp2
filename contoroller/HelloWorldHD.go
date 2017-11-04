package contoroller

import (
	"github.com/labstack/echo"
	"net/http"
)

func HelloWorldHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		return c.String(http.StatusOK, "Hello World")
	}
}
