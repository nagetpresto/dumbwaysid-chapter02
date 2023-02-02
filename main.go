/*--------------------------------------------------------------
# General
--------------------------------------------------------------*/
package main

import (
    "fmt"
	"net/http"
    "html/template"
	"strconv"

    "github.com/gorilla/mux"
)

/*--------------------------------------------------------------
# Global var (object to send and be displayed in html)
--------------------------------------------------------------*/
var Data = map[string]interface{}{
    "Title": "Personal Web",
}

/*--------------------------------------------------------------
# Main Routing Function
--------------------------------------------------------------*/

func main() {
	// create new router
    route := mux.NewRouter()

	// accessing static assets (route, funct) >> stripprefix = membungkus >> fileserver = handler
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	
	// routing for each page
    route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/add-project", addProject).Methods("GET")
	route.HandleFunc("/add-card", addCard).Methods("POST")
	route.HandleFunc("/project/{id}", projectBlog).Methods("GET")
	route.HandleFunc("/contact-me", contactMe).Methods("GET")

	// starting server
    fmt.Println("Server running on port 5000")
    http.ListenAndServe("localhost:5000", route)
}

/*--------------------------------------------------------------
# Home Routing Function
--------------------------------------------------------------*/
func home(w http.ResponseWriter, r *http.Request) {
    // w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// parsing file
    var tmpl, err = template.ParseFiles("views/index.html")
	// (in response) if parsing error
    if err != nil {
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }

	// execute file
    tmpl.Execute(w, Data)
}

/*--------------------------------------------------------------
# Add Project Routing Function
--------------------------------------------------------------*/
func addProject(w http.ResponseWriter, r *http.Request) {
    // w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// parsing file
    var tmpl, err = template.ParseFiles("views/project.html")
	// (in response) if parsing error
    if err != nil {
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }

	// execute file
    tmpl.Execute(w, Data)
}

/*--------------------------------------------------------------
# Add Card Routing Function
--------------------------------------------------------------*/
func addCard(w http.ResponseWriter, r *http.Request) {
    // w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// (in request) only var error >> parsing form html
    err := r.ParseForm()
	// (in response) if parsing error
    if err != nil {
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }

	// getting data
    fmt.Println("Project Name : " + r.PostForm.Get("pName"))
    fmt.Println("Start Date : " + r.PostForm.Get("sDate"))
	fmt.Println("End Date : " + r.PostForm.Get("eDate"))
	fmt.Println("Description : " + r.PostForm.Get("description"))
	fmt.Println("Icon1 : " + r.PostForm.Get("c1"))
	fmt.Println("Icon2 : " + r.PostForm.Get("c2"))
	fmt.Println("Icon3 : " + r.PostForm.Get("c3"))
	fmt.Println("Icon4 : " + r.PostForm.Get("c4"))

    http.Redirect(w, r, "/add-project", http.StatusMovedPermanently)
}

/*--------------------------------------------------------------
# Project Blog Routing Function
--------------------------------------------------------------*/
func projectBlog(w http.ResponseWriter, r *http.Request) {
    // w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// parsing file
	var tmpl, err = template.ParseFiles("views/project-blog.html")

	// (in response) if parsing error
    if err != nil {
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }

	// var for storing id
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// joining id to global var
    resp := map[string]interface{}{
        "Data": Data,
        "Id":   id,
    }
    // execute file
    tmpl.Execute(w, resp)

}

/*--------------------------------------------------------------
# Contact Me Routing Function
--------------------------------------------------------------*/
func contactMe(w http.ResponseWriter, r *http.Request) {
    // w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// parsing file
    var tmpl, err = template.ParseFiles("views/contact-me.html")
	// (in response) if parsing error
    if err != nil {
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }

	// execute file
    tmpl.Execute(w, Data)
}