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

type VRoutes map[string]func(w http.ResponseWriter, r *http.Request)
type VViews map[string]string

/*
map[string]*template.Template
*/
type VTemplates map[string]*template.Template

// controller is consist of views and templates
// Views are strning data of the html file associated with the controller
type Controller struct {
	Views     VViews
	Templates VTemplates
	Routes    VRoutes
}
type VControllers map[string]Controller

type VData map[string]interface{}

var serverIP = "localhost"
var serverPort = "8080"
var staticFolder embed.FS
var Controllers = make(VControllers)
var Routes = make(VRoutes)

func CreateController(_view VViews) Controller {
	return Controller{
		Views:     _view,
		Templates: make(VTemplates),
		Routes:    make(VRoutes),
	}
}

func init() {
	Controllers = make(VControllers)
	for _, c := range Controllers {
		c.Templates = make(VTemplates)
	}
	Routes = make(VRoutes)

	// RegisterRoutes()
	// RegisterControllers()
}

// func RegisterRoutes() {

// }

// func RegisterControllers() {

// }
