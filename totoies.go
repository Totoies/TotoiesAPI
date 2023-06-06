package Totois

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

/*
This function will Start the server
*/
func Buid() {
	// Start the server
	CreateRoutes()
	BuildControllers()

	fmt.Printf("Server starting on http://%s:%s", serverIP, serverPort)
	log.Fatal(http.ListenAndServe(serverIP+":"+serverPort, nil))
}

/*
Add Routing to out Web Application

	Ex. CreateRoutes()
*/
func CreateRoutes() {
	for route, function := range Routes {
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
Initialise all the controllers and there Required Views
*/
func BuildControllers() {

	for _, __controller := range Controllers {
		// loop through all views
		for __name, __view := range __controller.views {

			__file, __err := staticFolder.ReadFile(__view)
			if __err != nil {
				log.Fatal("Not able to read ", __view)
			}

			var _template, _ = template.New("template").Parse(string(__file))
			__controller.templates[__name] = VTemplate{
				template: _template,
			}

			// *__controller.templates[__name].template = *_template

			if __err != nil {
				log.Fatal("Not able to create new template")
				log.Fatal(__err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

/*
Execute our template
Exec(w http.ResponseWriter, __template *VTemplate)
*/
func Exec(w http.ResponseWriter, __template *VTemplate) {
	__template.template.Execute(w, __template.vars)
}

/*
initialise the static folder for later use
*/
func InitStaticFolder(_staticFolder embed.FS) {
	staticFolder = _staticFolder

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFolder))))
}

/*
Get the value inside a file
to get the file need to specify the file path
EX. GetFileData("static/file.txt")
*/
func GetFileData(path *string) *[]byte {

	data, err := staticFolder.ReadFile(*path)
	if err != nil {
		return nil
	}
	return &data
}

func GetStaticFolder() *embed.FS {
	return &staticFolder
}
