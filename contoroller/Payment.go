package contoroller

import (
	"github.com/labstack/echo"
	"net/http"
)

func PaymentHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		return c.String(http.StatusOK, "Payment Param is "+c.Param("param"))
	}
}
