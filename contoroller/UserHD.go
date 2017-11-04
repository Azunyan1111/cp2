package contoroller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/Azunyan1111/cp/model"
)

func GetUserHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		userData, err := model.SelectUserDataById(c.Param("userId"))
		if err != nil{
			return c.JSON(http.StatusBadRequest, model.Error{Status:http.StatusBadRequest, Message:err.Error()})//Message:"Not found user"})
		}else{
			return c.JSON(http.StatusOK, userData)
		}
	}
}

func SetUserHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		return c.String(http.StatusOK, "You name is "+c.Param("userId"))
	}
}
