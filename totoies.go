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
		fmt.Println(route)
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

	// __controller
	for __cname, __controller := range Controllers {
		fmt.Println(__cname)
		// loop through all views
		for __name, __view := range __controller.Views {
			fmt.Println(__view)
			__file, __err := staticFolder.ReadFile(__view)
			if __err != nil {
				log.Fatal("Not able to read ", __view)
			}

			__controller.Templates[__name], _ = template.New("template").Parse(string(__file))
			// __controller.Templates[__name] = _template

			// // *__controller.templates[__name].template = *_template

			if __err != nil {
				log.Fatal("Not able to create new template")
				log.Fatal(__err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

/*
Execute our template
Exec(w http.ResponseWriter, __templates *VTemplates, __templateName string, data any)
totoies.Exec(w, &Home.Templates, "Home")
*/
func Exec(w http.ResponseWriter, __templates *VTemplates, __templateName string, data VData) {
	// for _, v := range (*__templates)[__templateName].Vars {
	// 	fmt.Println(v)
	// }
	(*__templates)[__templateName].Execute(w, data)
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
