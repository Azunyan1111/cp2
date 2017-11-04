package main

import (
	"github.com/Azunyan1111/cp/contoroller"
	"github.com/Azunyan1111/cp/model"
	"github.com/Azunyan1111/anaconda"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// データベース接続
	model.DataBaseInit()
	// anaconda 設定
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/", contoroller.MainPage())    // メインページ。なんでも動作確認できる。
	e.GET("/all", contoroller.DataBase()) // メインページ。なんでも動作確認できる。
	//  ユーザー
	e.GET("/login", contoroller.LoginHD())          // ログイン兼サインアップのHTML
	e.GET("/user/:userId", contoroller.GetUserHD()) // プロフィール取得（ＧＥＴ）
	e.PUT("/user/:userId", contoroller.SetUserHD()) // プロフィール更新（ＰＵＴ）
	// ポイント
	e.GET("/request", contoroller.RequestHD()) // ポイント請求ＵＲＬ＆ＱＲコード生成（ＰＯＳＴ）
	e.GET("/payment", contoroller.PaymentHD()) // ポイント支払い

	// ツイッターログイン
	e.GET("/request_token", contoroller.RequestTokenHD()) //apiを利用する時に使うリクエスト（TwitterログインのURL）
	e.GET("/access_token", contoroller.AccessTokenHD())   //apiを利用する時に使うアクストークン

	// サーバー起動
	e.Start(":" + os.Getenv("PORT")) //ポート番号指定してね
}
