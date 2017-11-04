package contoroller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/Azunyan1111/cp/model"
	"strconv"
)

func GetUserHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		if c.Param("userId") == strconv.FormatInt(model.TestUser.Id,10){
			return c.String(http.StatusOK, "You name is " + model.TestUser.UserName)
		}else{
			return c.String(http.StatusBadRequest, c.Param("userId") + " is not found user.")
		}
	}
}

func SetUserHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		return c.String(http.StatusOK, "You name is "+c.Param("userId"))
	}
}
