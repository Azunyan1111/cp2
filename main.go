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
	e.GET("/", model.HelloWorldHD())            // Hello World と表示するだけのチエック用
	e.GET("/login", model.LoginHD())            // ログイン兼サインアップ（ＰＯＳＴ）
	e.GET("/user/:userId", model.UserHD())      // プロフィール取（ＧＥＴ）・プロフィール更新（ＰＵＴ）
	e.GET("/request", model.RequestHD())        // ポイント請求ＵＲＬ＆ＱＲコード生成（ＰＯＳＴ）
	e.GET("/payment/:param", model.PaymentHD()) // ポイント支払い

	// サーバー起動
	e.Start(":" + os.Getenv("PORT")) //ポート番号指定してね
}
