package service

import (
	"fmt"
	"github.com/unrolled/render"
	"github.com/weirdsnap/Blog/entities"
	"net/http"
	"strconv"
)

func apiGetByIdHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id := req.Form["id"]
		fmt.Println("id:", id)
		// todo :
		// check the id
		intid, _ := strconv.Atoi(id[0])
		// search the id in db, return the content of this id
		if a := entities.ArticleInfoService.FindByID(intid); a != nil {
			formatter.JSON(w, http.StatusOK, struct {
				CONTENT string `json:"content"`
				ID      string `json:"id"`
			}{CONTENT: a.Content, ID: id[0]})
		}

		// if (len(id) != 0) {
		// 	formatter.JSON(w, http.StatusOK, struct {
		// 		CONTENT string `json:"content"`
		// 		ID string `json:"id"`
		// 	}{CONTENT: "xxxxxxxxxxxxxxxxx",ID: id[0]})
		// }
	}
}


func apiGetAllHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// func (*ArticleInfoAtomicService) FindAll() []ArticleInfo {}
		if articles := entities.ArticleInfoService.FindAll(); articles != nil {
			formatter.JSON(w, http.StatusOK, struct {
				ARTICLES struct `json:"articles"`
			}{ARTICLES: a})
		}
	}
}