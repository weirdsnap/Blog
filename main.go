package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/weirdsnap/Blog/service"
	"os"
)

const (
	// set the default port 80
	PORT string = "80"
)

func main() {
	// use environment variables if exist
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}
	// use flag "-p" to set port if p's value exist
	pPort := flag.StringP("port", "p", port, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}
	fmt.Println("Server prot change to ", port)
	// run the server
	server := service.NewServer()
	server.Run(":" + port)
}
