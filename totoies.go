package Totois

import (
	"embed"
	"fmt"
	"log"
	"net/http"
)

// Our Routing
/* Routes{
	"route": functionfunc(w http.ResponseWriter, r *http.Request) {
		Body----
	}
}
*/
type Routes map[string]func(w http.ResponseWriter, r *http.Request)

/*
This function will Start the server
*/
func Buid() {
	// Start the server
	fmt.Printf("Server starting on http://%s:%s", serverIP, serverPort)
	log.Fatal(http.ListenAndServe(serverIP+":"+serverPort, nil))
}

/*
Add Routing to out Web Application

	Ex. CreateRoutes(Routes{
	            "/": func (w http.ResponseWriter, r *http.Request) {
	                fmt.Fprint(w, "Hello, World!")
	            }
	        })
*/
func CreateRoutes(_routes map[string]func(w http.ResponseWriter, r *http.Request)) {
	for route, function := range _routes {
		http.HandleFunc(route, function)
	}
}

/*
Configure the ServerIp which most of the cases going to be localhost
and ServerPort
*/
func ConfigApplication(_ServerIP string, _ServerPort string) {
	serverIP = _ServerIP
	serverPort = _ServerPort
}

/*
initialise the static folder for later use
*/
func InitStaticFolder(_staticFolder embed.FS) {
	staticFolder = _staticFolder
}
