package contoroller

import (
	"github.com/Azunyan1111/cp2/model"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
)

func PaymentHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		rqId := c.QueryParam("rqid")
		rqCp := c.QueryParam("rqcp")
		myId := c.QueryParam("myid")
		//パラメーター確認
		if rqId == "" || rqCp == "" || myId == ""{
			return c.JSON(http.StatusBadRequest, model.Error{Status: http.StatusBadRequest, Message: "Not found param"})
		}
		// 受信するユーザーが存在しない場合
		if !model.IsUserExistById(rqId) {
			return c.JSON(http.StatusBadRequest, model.Error{Status: http.StatusBadRequest, Message: "Not found receive user"})
		}
		// 送信するユーザーが存在しない場合
		if !model.IsUserExistById(myId) {
			return c.JSON(http.StatusBadRequest, model.Error{Status: http.StatusBadRequest, Message: "Not found send user"})
		}

		// クッキーから自分の情報アクセストークンを取得する。
		//token, err := c.Cookie("Token")
		//if err != nil {
		//	token = nil
		//}
		//secret, err := c.Cookie("Secret")
		//if err != nil {
		//	secret = nil
		//}

		//api := anaconda.NewTwitterApi(token.Value, secret.Value)
		//response, err := api.GetSelf(url.Values{})
		//if err != nil {
		//	log.Println(err)
		//	return c.JSON(http.StatusBadRequest, model.Error{Status: http.StatusBadRequest, Message: "Please re login. Error By Twitter API"})
		//}
		// 自分が存在するか確認
		//if !model.IsUserExistByTwitter(response.Id) {
		//	return c.JSON(http.StatusBadRequest, model.Error{Status: http.StatusBadRequest, Message: "Please sign up. Error By Not Found Request User"})
		//}

		// 自分にポイントが存在するか確認
		myPoint, err := model.SelectUserPointById(myId)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, model.Error{Status: http.StatusBadRequest, Message: "Please Retry. Error By SQL"})
		}
		// 自分のポイントが存在しない場合
		intCp, err := strconv.ParseInt(rqCp, 10, 64)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, model.Error{Status: http.StatusInternalServerError, Message: "Please retry. Error By ParseInt"})
		}
		if myPoint <= intCp {
			return c.JSON(http.StatusBadRequest, model.Error{Status: http.StatusBadRequest, Message: "Please change Point. Error By Points missing"})
		}

		// 自分のポイントを減らす
		if err := model.UpdatePointSubById(myId, intCp); err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, model.Error{Status: http.StatusInternalServerError, Message: "Please retry. Error can not sub request user."})
		}
		// 相手のポイントを増やす
		if err := model.UpdatePointAddById(rqId, intCp); err != nil {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, model.Error{Status: http.StatusInternalServerError, Message: "Please retry. Error can not add."})
		}
		return c.JSON(http.StatusOK,  "ok")
	}
}
