package service

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	n := negroni.Classic()

	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	// get the os environment variable "WEBROOT", means the web server home
	webRoot := os.Getenv("WEBROOT")

	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory!")
		} else {
			webRoot = root
		}
	}
	// api test
	mx.HandleFunc("/api/test", apiTestHandler(formatter)).Methods("GET")
	//
	mx.HandleFunc("/", homeHandler(formatter)).Methods("GET")

	mx.HandleFunc("/login", loginHandler(formatter)).Methods("GET", "post")

	mx.HandleFunc("/{[a-zA-z]+}", unknowHandler())

	// static file server , and dir redirected to webRoot/assets/
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))

	//	mx.HandleFunc("/hello/{id}", testHandler(formatter)).Methods("GET")
}

func unknowHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(501)
		fmt.Fprint(w, "Not implemented")
	}
}

/*
func testHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Hello " + id})
	}
}
*/
