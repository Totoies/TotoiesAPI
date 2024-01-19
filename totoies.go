package Totois

import (
	"embed"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

/*
This function will Start the server
*/
func Buid() {
	// Start the server
	CreateRoutes(Routes)
	BuildControllers()

	fmt.Printf("Server starting on http://%s:%s", serverIP, serverPort)
	log.Fatal(http.ListenAndServe(serverIP+":"+serverPort, nil))
}

/*
Add Routing to out Web Application

	Ex. CreateRoutes()
*/
func CreateRoutes(__routes VRoutes) {
	for route, function := range __routes {
		fmt.Println(route)
		http.HandleFunc(route, function)
	}
}

/*
Configure the ServerIp which most of the cases going to be localhost
and ServerPort
*/
func ServerConfig(_ServerIP string, _ServerPort string) {
	serverIP = _ServerIP
	serverPort = _ServerPort
}

/*
Initialise all the controllers and there Required Views
Gettings all the html views name and read the file and create template out of it
user have to give the View name Example if the Html view is in /Static/Views/Home/index.html
*/
func BuildControllers() {

	viewDirectory := "Static/views/"
	// __controller
	for __cname, __controller := range Controllers {
		fmt.Println(__cname)
		// Creating local Routes
		CreateRoutes(__controller.Routes)

		// loop through all views
		for __name, __view := range __controller.Views {
			fmt.Println(__view)
			// holding the index view
			mainViewFilePath := viewDirectory + __name + "/" + __view + ".html"
			ControllerViewFolder := viewDirectory + __name + "/" // folder path of the Current controller view

			// Get the main view first
			tempV, _ := ReadStaticFile(mainViewFilePath)
			mainView := string(tempV)
			if mainView == "" {
				mainView = "There are some issue reading the view file please check the server log"
				__controller.Templates[__name], _ = template.New("template").Parse(string(mainView))
				return
			}
			AllFilesInView, _ := readFilesInDirectory(ControllerViewFolder)

			for _filename, _value := range AllFilesInView {
				mainView = strings.Replace(mainView, "{{@include("+_filename+")}}", string(_value), -1)
			}
			// gettign Footer and header file
			tempV, _ = ReadStaticFile(ControllerViewFolder + "header.html")
			header := string(tempV)
			tempV, _ = ReadStaticFile(ControllerViewFolder + "footer.html")
			footer := string(tempV)
			mainView = header + mainView + footer

			__controller.Templates[__name], _ = template.New("template").Parse(mainView)
		}
	}
}

// function to read files from static directory
func ReadStaticFile(_filePath string) ([]byte, error) {
	__file, __err := staticFolder.ReadFile(_filePath)
	if __err != nil {
		log.Fatal(__err.Error(), http.StatusInternalServerError)
		return nil, __err
	}

	return __file, nil

}

// reads all the files in that directory and returns "folder name" : and folder information in byte
func readFilesInDirectory(directoryPath string) (map[string][]byte, error) {
	fileMap := make(map[string][]byte)

	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		filePath := filepath.Join(directoryPath, file.Name())

		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		fileMap[file.Name()] = content
	}

	return fileMap, nil
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

func PrintCurrentDir() {
	currentdir, err := os.Getwd()

	if err != nil {
		fmt.Println("No Informtaion for the current directory")
		return
	}

	fmt.Println(currentdir)
}
