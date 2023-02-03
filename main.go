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
    "IsLogin": true,
}

type Project struct {
    Pname       string
    Sdate       string
    Edate       string
    Duration    string
    Description string
    Icon1       string
    Icon2       string
    Icon3       string
    Icon4       string
}

var ProjectCard = []Project{
    {
        Pname:          "CRUD WEB",
        Sdate:          "02/02/2023",
        Edate:          "02/09/2023",
        Duration:       "3 Month(s)",
        Description:    "asdas asdsad asdas asd",
        Icon1:          "fab fa-php me-4",
        Icon2:          "fab fa-js-square me-4",
        Icon3:          "fab fa-python me-4",
        Icon4:          "fab fa-html5 me-4",
    },
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
    route.HandleFunc("/edit-card/{id}", editCard).Methods("GET")
    route.HandleFunc("/delete-card/{id}", deleteCard).Methods("GET")
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

    respCard := map[string]interface{}{
        "Data":  Data,
        "ProjectCard": ProjectCard,
    }
	// execute file
    tmpl.Execute(w, respCard)
}

/*--------------------------------------------------------------
# Project Routing Function
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
    pName := r.PostForm.Get("pName")
    sDate := r.PostForm.Get("sDate")
    eDate := r.PostForm.Get("eDate")
    description := r.PostForm.Get("description")
    c1 := r.Form.Get("c1")
    c2 := r.Form.Get("c2")
    c3 := r.Form.Get("c3")
    c4 := r.Form.Get("c4")

    if c1 == "on" {
        c1 = "fab fa-php me-4"
    }else {
        c1 = ""
    }

    if c2 == "on" {
        c2 = "fab fa-js-square me-4"
    }else {
        c2 = ""
    }

    if c3 == "on" {
        c3 = "fab fa-python me-4"
    }else {
        c3 = ""
    }

    if c4 == "on" {
        c4 = "fab fa-html5 me-4"
    }else {
        c4 = ""
    }

    var newProjectCard = Project{
        Pname:          pName,
        Sdate:          sDate,
        Edate:          eDate,
        Duration:       "3 Month(s)",
        Description:    description,
        Icon1:          c1,
        Icon2:          c2,
        Icon3:          c3,
        Icon4:          c4,
    }

    ProjectCard = append(ProjectCard, newProjectCard)
    fmt.Println(ProjectCard)
    http.Redirect(w, r, "/add-project", http.StatusMovedPermanently)
}

/*--------------------------------------------------------------
# Edit Card Routing Function
--------------------------------------------------------------*/
func editCard(w http.ResponseWriter, r *http.Request){

    // var for storing id
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

    editCard := Project{}

    for i, data := range ProjectCard {
        if i == id {
            editCard = Project{
                Pname:          data.Pname,
                Sdate:          data.Sdate,
                Edate:          data.Edate,
                Duration:       "3 Month(s)",
                Description:    data.Description,
                Icon1:          data.Icon1,
                Icon2:          data.Icon2,
                Icon3:          data.Icon3,
                Icon4:          data.Icon4,
            }
        }

    }

    // joining id to global var
    resp := map[string]interface{}{
        "Data":  Data,
        "EditCard": editCard,
    }

    fmt.Println(resp)
}

/*--------------------------------------------------------------
# Delete Card Routing Function
--------------------------------------------------------------*/
func deleteCard(w http.ResponseWriter, r *http.Request){
    id, _ := strconv.Atoi(mux.Vars(r)["id"])

    ProjectCard = append(ProjectCard[:id], ProjectCard[id+1:]...)
    http.Redirect(w, r, "/", http.StatusMovedPermanently)

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

    projectBlog := Project{}

    for i, data := range ProjectCard {
        if i == id {
            projectBlog = Project{
                Pname:          data.Pname,
                Sdate:          data.Sdate,
                Edate:          data.Edate,
                Duration:       "3 Month(s)",
                Description:    data.Description,
                Icon1:          data.Icon1,
                Icon2:          data.Icon2,
                Icon3:          data.Icon3,
                Icon4:          data.Icon4,
            }
        }

    }

	// joining id to global var
    resp := map[string]interface{}{
        "Data":  Data,
        "Project": projectBlog,
    }
    // execute file
    fmt.Println(resp)
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