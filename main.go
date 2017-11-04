package main

import (
	"net/http"
	"github.com/Azunyan1111/cp/model"
	"os"
)


func main() {
	http.HandleFunc("/", model.HelloWolrdHD) // Hello World と表示するだけのチエック用
	// ユーザー情報関連
	http.HandleFunc("/login", model.HelloWolrdHD) // ログイン兼サインアップ（ＰＯＳＴ）
	http.HandleFunc("/user/:userId", model.HelloWolrdHD) // プロフィール取（ＧＥＴ）・プロフィール更新（ＰＵＴ）
	// 請求
	http.HandleFunc("/request", model.HelloWolrdHD) // ポイント請求ＵＲＬ＆ＱＲコード生成（ＰＯＳＴ）
	http.HandleFunc("/payment/:param", model.HelloWolrdHD) // ポイント支払い
	http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}
