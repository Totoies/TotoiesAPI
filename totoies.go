package Totois

import (
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

// Our Application
type totoies_app struct {
	ServerIP   string
	ServerPort string
	Routes     map[string]func(w http.ResponseWriter, r *http.Request)
}

// app with our default configaration
var App = totoies_app{
	ServerIP:   "localhost",
	ServerPort: "8080",
}

/*
This function will Start the server
*/
func Buid() {
	// Start the server
	fmt.Printf("Server starting on http://%s:%s", App.ServerIP, App.ServerPort)
	log.Fatal(http.ListenAndServe(App.ServerIP+":"+App.ServerPort, nil))
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
	App.Routes = _routes
	for route, function := range App.Routes {
		http.HandleFunc(route, function)
	}
}

/*
Configure the ServerIp which most of the cases going to be localhost
and ServerPort
*/
func ConfigApplication(_ServerIP string, _ServerPort string) {
	App.ServerIP = _ServerIP
	App.ServerPort = _ServerPort
}
