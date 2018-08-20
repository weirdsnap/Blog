package service

import (
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

	// deal with no os environment variable "WEBROOT"
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory!")
		} else {
			webRoot = root
		}
	}
	// api 
	mx.HandleFunc("/api/getbyid/", apiGetByIdHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/getall", apiGetAllHandler(formatter)).Methods("GET")
	// mx.HandleFunc("/api/getall", apiTestHandler(formatter)).Methods("GET")

	// static file server , and dir redirected to webRoot/assets/
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))

	// mx.HandleFunc("/{[a-zA-z]+}", unknowHandler())
	//	mx.HandleFunc("/hello/{id}", testHandler(formatter)).Methods("GET")
}
