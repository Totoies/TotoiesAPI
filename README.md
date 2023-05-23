# Totois
A file which will be responsible Creating servers, Updating Data Base

# Creating your First Application

    `go get github.com/Totoies/Totoies`

        
    import( totoies "github.com/Totoies/Totoies")

    func main() {

	    totoies.CreateRoutes(totoies.Routes{
            "/": func (w http.ResponseWriter, r *http.Request) {
                fmt.Fprint(w, "Hello, World!")
            }
        })

        totoies.Buid()
    }

# How to add a static Directory

        -> Add this line in your application [static/*] is the directory
        //go:embed static/*
        var staticDir embed.FS

        totoies.InitStaticFolder(staticDir)