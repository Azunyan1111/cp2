package main

import (
	"github.com/Azunyan1111/cp/contoroller"
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
	e.GET("/", contoroller.HelloWorldHD()) // Hello World と表示するだけのチエック用
	//  ユーザー
	e.GET("/login", contoroller.LoginHD())          // ログイン兼サインアップ（ＰＯＳＴ）
	e.GET("/user/:userId", contoroller.GetUserHD()) // プロフィール取（ＧＥＴ）
	e.PUT("/user/:userId", contoroller.SetUserHD()) // プロフィール更新（ＰＵＴ）
	// ポイント
	e.GET("/request", contoroller.RequestHD())        // ポイント請求ＵＲＬ＆ＱＲコード生成（ＰＯＳＴ）
	e.GET("/payment/:param", contoroller.PaymentHD()) // ポイント支払い

	// ツイッターログイン
	e.GET("/request_token", contoroller.RequestTokenHD()) //apiを利用する時に使うリクエスト
	e.GET("/access_token", contoroller.AccessTokenHD())   //apiを利用する時に使うアクストークン

	// サーバー起動
	e.Start(":" + os.Getenv("PORT")) //ポート番号指定してね
}
