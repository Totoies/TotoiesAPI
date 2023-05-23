package Totois

import (
	"embed"
	"net/http"
)

// Storing the Configarations of the Application
// Ex. Enviourment which can be either Dev or Prod

// Enviourment
const (
	Prod = false
	Dev  = true
)

var Enviourment = Dev

func SetEnviourment(_env bool) {
	Enviourment = _env
}

var serverIP = "localhost"
var serverPort = "8080"
var staticFolder embed.FS

// Our Routing
/* Routes{
	"route": functionfunc(w http.ResponseWriter, r *http.Request) {
		Body----
	}
}
*/
type Routes map[string]func(w http.ResponseWriter, r *http.Request)
