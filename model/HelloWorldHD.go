package model

import (
	"net/http"
	"fmt"
)

func HelloWolrdHD(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}
