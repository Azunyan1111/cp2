package contoroller

import (
	"github.com/Azunyan1111/cp2/model"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		return c.HTML(http.StatusOK, ""+
			"<a href='/login'>login</a><br>"+
			"<a href='/user/31'>ユーザー31のユーザープロフィール</a><br>"+
			"<a href='/request?cp=100&id=31'>ユーザー31が100CP請求する</a><br>"+
			"<a href='/payment?rqid=41&rqcp=100&myid=31'>ユーザー41にログインしているユーザーが100CP送る</a><br>"+
			"<a href='/all'>全てのユーザーデータ</a>")
	}
}
func DataBase() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		allUser := model.SelectAllUserLIMIT100()
		log.Println(allUser)
		return c.JSON(http.StatusOK, allUser)
	}
}
