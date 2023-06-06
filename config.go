package Totois

import (
	"embed"
	"html/template"
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

// Our Routing
/* Routes{
	"route": functionfunc(w http.ResponseWriter, r *http.Request) {
		Body----
	}
}
*/

// controller
type Controller struct {
	views     VViews
	templates VTemplates
}
type VControllers []Controller

type VRoutes map[string]func(w http.ResponseWriter, r *http.Request)
type VViews map[string]string

type VTemplate struct {
	template *template.Template
	vars     map[string]interface{}
}
type VTemplates map[string]VTemplate

var serverIP = "localhost"
var serverPort = "8080"
var staticFolder embed.FS
var Controllers VControllers
var Routes VRoutes
