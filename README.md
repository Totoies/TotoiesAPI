# Totois
A file which will be responsible Creating servers, Updating Data Base

# Creating your First Application

    `go get github.com/Totoies/Totoies@main`

        
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

# How to add a static Directory

        -> Add this line in your application [static/*] is the directory
        //go:embed static/*
        var staticDir embed.FS

        totoies.InitStaticFolder(staticDir)

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