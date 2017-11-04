package contoroller

import (
	"github.com/Azunyan1111/cp/model"
	"github.com/labstack/echo"
	"net/http"
	"net/url"
	"os"
)

func RequestHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		if c.QueryParam("id") == "" || c.QueryParam("cp") == "" {
			return c.JSON(http.StatusBadRequest, model.Error{Status: http.StatusBadRequest, Message: "Not found param"})
		}

		// param
		urlParam := url.Values{}
		urlParam.Add("id", c.QueryParam("id"))
		urlParam.Add("cp", c.QueryParam("cp"))

		// url
		json := model.RequestJson{
			Url: os.Getenv("URL") + "payment" + "?" + urlParam.Encode(),
			QrCode: "https://chart.googleapis.com/chart?chs=300x300&cht=qr&chl=" +
				os.Getenv("URL") + "payment" + "?" + urlParam.Encode(),
		}
		// {http://localhost:8080/payment?cp=100&id=31 https://chart.googleapis.com/chart?chs=300x300&cht=qr&chl=http://localhost:8080/payment?cp=100&id=31}
		return c.JSON(http.StatusOK, json)
	}
}
