package main

import (
	"net/http"
	"github.com/Azunyan1111/cp/model"
	"os"
)


func main() {
	http.HandleFunc("/", model.HelloWolrdHD) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}
