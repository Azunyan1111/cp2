package contoroller

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
	"github.com/labstack/echo"
	"net/http"
	"net/url"
	"os"
	"github.com/Azunyan1111/cp/model"
)

var credential *oauth.Credentials

type Authentication struct {
	Token   string `json:"token"`
	Secret string `json:"secret"`
}


func RequestTokenHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		// set keys
		anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
		anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))

		//リクエストしてユーザーを飛ばすURLとか貰う
		url, tmpCred, err := anaconda.AuthorizationURL(os.Getenv("URL") + "access_token")
		if err != nil {
			return err
		}
		credential = tmpCred
		//ツイッターの認証ページに飛ばす
		http.Redirect(c.Response(), c.Request(), url, http.StatusFound)

		return c.String(http.StatusOK, "Login URL")
	}
}

func AccessTokenHD() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する

		//アクセストークンとシークレットをもらってくる便利な奴
		tokens, _, err := anaconda.GetCredentials(credential, c.Request().URL.Query().Get("oauth_verifier"))
		if err != nil {
			return err
		}

		// クッキーに貼り付ける
		c.SetCookie(&http.Cookie{
			Name:  "Token",      // ここにcookieの名前を記述
			Value: tokens.Token, // ここにcookieの値を記述
		})
		c.SetCookie(&http.Cookie{
			Name:  "Secret",      // ここにcookieの名前を記述
			Value: tokens.Secret, // ここにcookieの値を記述
		})


		// uuidを貰い受ける
		api := anaconda.NewTwitterApi(tokens.Token, tokens.Secret)
		response, err := api.GetSelf(url.Values{})
		if err != nil {
			//帰ってきたエラーメッセージがそのまま出力されるよ！
			return err
		}

		if model.IsUserExistByTwitter(response.Id){
			// TODO:既存
			return c.String(http.StatusOK, "ようこそ！")
		}else{
			// TODO:新規
			model.InsertNewUserByTwitter(response.Id)
			return c.String(http.StatusOK, "プロフィールを登録しよう！")

		}
	}
}
