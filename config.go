package Totois

import "embed"

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
