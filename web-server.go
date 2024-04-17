package main

import (
	"fmt"
	"net/http"
	"os"
)

//Create a folder (that I called Resources) inside your directory.
//Put all your images, JS scripts and CSS stylesheets into the folder "Resources", (keep the rest you want to serve dynamically such as the html files within the root folder).

const ResourceDir = "/Resources/"

func homeGetHandler(w http.ResponseWriter, r *http.Request) {
	html, err := os.ReadFile("index.html")
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "%s", html)
}



func main() {

	//Resources - Statically served
	var HandlerResources = http.StripPrefix(ResourceDir,
		http.FileServer(http.Dir("."+ResourceDir)),
	)
	http.Handle(ResourceDir, HandlerResources)

	//Homepage - Dynamically served
	http.HandleFunc("/", homeGetHandler)

	http.ListenAndServe(":8080", nil)

}
