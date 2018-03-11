package service

import (
	"fmt"
	"github.com/gorilla/schema"
	"github.com/unrolled/render"
	"net/http"
)

type User struct {
	Name     string
	Password string
}

func loginHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		user := new(User)
		if req.Method == "GET" {
			formatter.HTML(w, http.StatusOK, "login", struct {
				Number string `json:"Number"`
				Time   string `json:"Time"`
			}{Number: "1", Time: "2017/11/21"})
		} else {
			req.ParseForm()
			fmt.Println("username:", req.Form["Name"])
			fmt.Println("password:", req.Form["Password"])

			decoder := schema.NewDecoder()
			err := decoder.Decode(user, req.PostForm)
			if err != nil {
				fmt.Println("user data decode error")
				return
			}
			formatter.HTML(w, http.StatusOK, "logined", struct {
				Username string `json:"Username"`
			}{Username: user.Name})

		}
	}
}
