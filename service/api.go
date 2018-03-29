package service

import (
	"github.com/unrolled/render"
	"net/http"
	"fmt"
)

func apiTestHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id := req.Form["id"]
		fmt.Println("id:", id)
		// todo :
		// check the id 
		// search the id in db
		// return the content of this id
		if (len(id) != 0) {
			formatter.JSON(w, http.StatusOK, struct {
				CONTENT string `json:"content"`
				ID string `json:"id"`
			}{CONTENT: "# 123 \n ## 321 \n",ID: id[0]})
		}
	}
}
