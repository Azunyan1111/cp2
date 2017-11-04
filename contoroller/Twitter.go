package contoroller

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
	"github.com/labstack/echo"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"github.com/Azunyan1111/cp/model"
)

var credential *oauth.Credentials

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

		// uuidを貰い受ける
		api := anaconda.NewTwitterApi(tokens.Token, tokens.Secret)
		response, err := api.GetSelf(url.Values{})
		if err != nil {
			//帰ってきたエラーメッセージがそのまま出力されるよ！
			return err
		}
		// TODO:ここでログイン処理＆サインアップ処理
		if response.Id == model.TestUser.TwitterId{
			return c.String(http.StatusOK, "Hello " + model.TestUser.Name)
		}else{
			return c.String(http.StatusOK, "プロフィールを登録しよう！")
		}

		return c.String(http.StatusOK, "Your Twitter ID is "+strconv.FormatInt(response.Id, 10))
	}
}
