package service

import (
	"github.com/unrolled/render"
	"net/http"
	"fmt"
)

func apiTestHandler(formatter *render.Render) http.HandlerFunc {

	// return func(w http.ResponseWriter, req *http.Request) {
	// 	formatter.JSON(w, http.StatusOK, struct {
	// 		ID      string `json:"id"`
	// 		Content string `json:"content"`
	// 	}{ID: "8675309", Content: "Hello from Go!"})
	// 	// formatter.JSON(w, http.StatusOK, "!somestruct")
	// }

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id := req.Form["id"]
		// password := req.Form["password"]
		fmt.Println("id:", id)
		// if () {
		formatter.JSON(w, http.StatusOK, struct {
			CONTENT string `json:"content"`
		}{CONTENT: "tebiechangdeyiduanhua"})
		// }
	}
}
