package service

import (
	"fmt"
	"time"
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
				ARTICLES []entities.ArticleInfo `json:"articles"`
			}{ARTICLES: articles})
		}
	}
}

func apiSetOneHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		// todo: create the unique id
		// t := time.Now()
		// fmt.Println(t.Unix()) // 1970 s 
		// timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
		// fmt.Println(timestamp) // more 
		id := time.Now().Unix(); 
		t := req.PostFormValue("title")
		c := req.PostFormValue("class")
		s := req.PostFormValue("content")
		a := entities.NewArticleInfo(entities.ArticleInfo{AID: int(id), Title: t, Class: c, Content: s})
		entities.ArticleInfoService.Save(a)
	}
}
