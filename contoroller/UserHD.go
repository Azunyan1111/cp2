package contoroller

import (
	"github.com/Azunyan1111/cp2/model"
	"github.com/labstack/echo"
	"net/http"
)

func GetUserHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
		userData, err := model.SelectUserDataById(c.Param("userId"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, model.Error{Status: http.StatusBadRequest, Message: err.Error()}) //Message:"Not found user"})
		} else {
			return c.JSON(http.StatusOK, userData)
		}
	}
}

func SetUserHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		return c.String(http.StatusOK, "You name is "+c.Param("userId"))
	}
}

func GetUserPointsHD() echo.HandlerFunc{
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
		points, err := model.SelectUserPointsById(c.Param("userId"))
		if err != nil{
			return c.JSON(http.StatusBadRequest, model.Error{Status:http.StatusBadRequest, Message:err.Error()})
		}
		return c.JSON(http.StatusOK, model.UserPointJson{Status:http.StatusOK, Points:points})
	}
}