# Totois
A file which will be responsible Creating servers, Updating Data Base

# Table content

- [Totois](#totois)
- [Table content](#table-content)
- [Setup Enviourment](#setup-enviourment)
  - [Download and Install Golang](#download-and-install-golang)
  - [Download and Install Git](#download-and-install-git)
- [Create your First Application](#create-your-first-application)
- [How to add a static Directory](#how-to-add-a-static-directory)
- [Example Controller](#example-controller)
- [include one Html inside another html](#include-one-html-inside-another-html)

# Setup Enviourment

## Download and Install Golang

    1. Download Golang from [this link](https://go.dev/doc/install)
    2. After Downloading just click on the downloaded executable and install it
## Download and Install Git
    1. Download Git form [this link](https://git-scm.com/downloads)
    2. After Download just click on the downloaded executable and install it

# Create your First Application

Run 
```powershell
go get github.com/Totoies/Totoies@main
```
        


```golang
import( totoies "github.com/Totoies/Totoies")

func main() {

	totoies.Routes{
        "/": func (w http.ResponseWriter, r *http.Request) {
            fmt.Fprint(w, "Hello, World!")
        }
    }

    totoies.Controllers {
        Controllers.Home
    }
    totoies.Buid()
}
```

# How to add a static Directory

```golang
// Add this line in your application [static/*] is the directory
//go:embed static/*
var staticDir embed.FS

// in main Function before Build
totoies.InitStaticFolder(staticDir)
```

# Example Controller 

View will in Static/Views/ControllerName/viewname.html
```golang
package Controller

import (
	"net/http"

	totoies "github.com/Totoies/Totoies"
)

var Home = totoies.CreateController(totoies.VViews{
	// "Home": "Static/Views/Home/home.html",
	"Home": "home",
})

/*
LoadHome() - Function load our controller and other settings
*/
func LoadHome(w http.ResponseWriter, r *http.Request) {

	totoies.Exec(w, &Home.Templates, "Home", totoies.VData{
		"WelcomeTxt": "Welcome All",
		"Title":      "First Application",
	})
}
```


# include one Html inside another html

syntax is 

@includ(htmlfilename)

it will include what ever is written in that file