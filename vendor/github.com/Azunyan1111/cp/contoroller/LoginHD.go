package contoroller

import (
	"github.com/labstack/echo"
	"net/http"
	"os"
)

func LoginHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		return c.HTML(http.StatusOK, "Login URL.<br><a href='"+os.Getenv("URL")+"request_token'>Twitter Login</a>")
	}
}
