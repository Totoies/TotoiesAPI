package Totois

import (
	"fmt"
	"log"
	"net/http"
)

// Enviourment
const (
	Prod = false
	Dev  = true
)

// Our Application
type totoies_app struct {
	ServerIP    string
	ServerPort  string
	Enviourment bool
	Routes      map[string]func(w http.ResponseWriter, r *http.Request)
}

// app with our default configaration
var App = totoies_app{
	ServerIP:    "localhost",
	ServerPort:  "8080",
	Enviourment: Dev,
}

/*
This function will Start the server
*/
func (a *totoies_app) Buid() {

	// Start the server
	fmt.Printf("Server starting on http://%s:%s", a.ServerIP, a.ServerPort)
	log.Fatal(http.ListenAndServe(a.ServerIP+":"+a.ServerPort, nil))
}

func (a *totoies_app) AddRoutes(_routes map[string]func(w http.ResponseWriter, r *http.Request)) {
	a.Routes = _routes
	for route, function := range a.Routes {
		http.HandleFunc(route, function)
	}
}
