package service

import (
	"github.com/unrolled/render"
	"net/http"
	"time"
)

func homeHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "index", struct {
			// Number string `json:"Number"`
			Time   string `json:"Time"`
		}{ Time: string(time.Now().Year())})
	}
}
