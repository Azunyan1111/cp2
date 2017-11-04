package main

import (
	"github.com/Azunyan1111/cp/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/", model.HelloWorldHD())            	// Hello World と表示するだけのチエック用
	//  ユーザー
	e.GET("/login", model.LoginHD())            	// ログイン兼サインアップ（ＰＯＳＴ）
	e.GET("/user/:userId", model.GetUserHD())    	// プロフィール取（ＧＥＴ）
	e.PUT("/user/:userId", model.SetUserHD())    // プロフィール更新（ＰＵＴ）
	// ポイント
	e.GET("/request", model.RequestHD())        	// ポイント請求ＵＲＬ＆ＱＲコード生成（ＰＯＳＴ）
	e.GET("/payment/:param", model.PaymentHD()) 	// ポイント支払い

	// ツイッターログイン
	e.GET("/request_token", model.RequestTokenHD()) //apiを利用する時に使うリクエスト
	e.GET("/access_token", model.AccessTokenHD())   //apiを利用する時に使うアクストークン

	// サーバー起動
	e.Start(":" + os.Getenv("PORT")) //ポート番号指定してね
}
