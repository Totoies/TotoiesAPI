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
        var Home = totoies.Controller{
            views: totoies.VViews{
                "templateName1": "path1"
                "templateName2": "path2"
            }
        }

        func (h *totoies.Controller) LoadHome(w http.ResponseWriter, r *http.Request) {
            totoies.Exec(w, Home.templates["TemplateName1"])
        }