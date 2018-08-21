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
	// get one article by id, return an article {aid,title,class,content}
	mx.HandleFunc("/api/getbyid/", apiGetByIdHandler(formatter)).Methods("GET")
	// get all article, return all [{aid,title,class,content},...]
	mx.HandleFunc("/api/getall", apiGetAllHandler(formatter)).Methods("GET")
	// get articles by class, return those articles with that class [{aid,title,class,content},...]
	// mx.HandleFunc("/api/getbyclass", apiGetByClassHandler(formatter)).Methods("GET")


	// mx.HandleFunc("/api/somefunction", apisonefunctionHandler(formatter)).Methods("GETorPOST")

	// static file server , and dir redirected to webRoot/assets/
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))

	// other function undone
	// mx.HandleFunc("/{[a-zA-z]+}", unknowHandler())
	//	mx.HandleFunc("/hello/{id}", testHandler(formatter)).Methods("GET")
}
