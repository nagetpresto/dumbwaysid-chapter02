/*--------------------------------------------------------------
# General
--------------------------------------------------------------*/
package main

import (
    "fmt"
	"net/http"
    "html/template"
	"strconv"
    "time"

    "github.com/gorilla/mux"
)

/*--------------------------------------------------------------
# Global var (object to send and be displayed in html)
--------------------------------------------------------------*/
var Data = map[string]interface{}{
    "IsLogin": true,
}

type Project struct {
    Id          int
    Pname       string
    Sdate       string
    SdateDetail string
    Edate       string
    EdateDetail string
    Duration    string
    Description string
    Icon1       string
    Icon1PClass string
    Icon1P      string
    Icon1Status string
    Icon2       string
    Icon2PClass string
    Icon2P      string
    Icon2Status string
    Icon3       string
    Icon3PClass string
    Icon3P      string
    Icon3Status string
    Icon4       string
    Icon4PClass string
    Icon4P      string
    Icon4Status string
}

var ID = []Project{
    {   
        Id:             0,
    },
}

var ProjectCard = []Project{
    {
        Pname:          "CRUD WEB",
        Sdate:          "2023-02-02",
        SdateDetail:    "02 February 2023",
        Edate:          "2023-09-02",
        EdateDetail:    "02 September 2023",
        Duration:       "3 Month(s)",
        Description:    "asdas asdsad asdas asd",
        Icon1:          "fab fa-php me-3",
        Icon1PClass:    "d-flex align-items-center col-6",
        Icon1P:         "PHP",
        Icon1Status:    "checked",
        Icon2:          "fab fa-js-square me-3",
        Icon2PClass:    "d-flex align-items-center col-6",
        Icon2P:         "Javascript",
        Icon2Status:    "checked",
        Icon3:          "fab fa-python me-3",
        Icon3PClass:    "d-flex align-items-center col-6",
        Icon3P:         "Python",
        Icon3Status:    "checked",
        Icon4:          "fab fa-html5 me-3",
        Icon4PClass:    "d-flex align-items-center col-6",
        Icon4P:         "HTML",
        Icon4Status:    "checked",
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
    route.HandleFunc("/update-card", updateCard).Methods("POST")
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
	// parsing file
    var tmpl, err = template.ParseFiles("views/index.html")
	// (in response) if parsing error
    if err != nil {
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }

    // mapping data to be displayed
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
	// (in request) output data error >> parsing form html
    err := r.ParseForm()

	// (in response) if parsing error
    if err != nil {
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }

	// getting data from input form
    pName       := r.PostForm.Get("pName")
    description := r.PostForm.Get("description")
    c1          := r.Form.Get("c1")
    Icon1PClass := ""
    Icon1P      := ""
    Icon1Status := ""
    c2          := r.Form.Get("c2")
    Icon2PClass := ""
    Icon2P      := ""
    Icon2Status := ""
    c3          := r.Form.Get("c3")
    Icon3PClass := ""
    Icon3P      := ""
    Icon3Status := ""
    c4          := r.Form.Get("c4")
    Icon4PClass := ""
    Icon4P      := ""
    Icon4Status := ""

    // parsing date (change format)
    layoutFormat := "2006-01-02"
    sDateForm    := r.PostForm.Get("sDate")
    eDateForm    := r.PostForm.Get("eDate")

    sDate, _ := time.Parse(layoutFormat, sDateForm)
    sDateDay := strconv.Itoa(sDate.Day())
    sDateMonth := sDate.Month().String()
    sDateYear := strconv.Itoa(sDate.Year())
    sDateFinal := sDateDay + " " + sDateMonth + " " + sDateYear

    eDate, _ := time.Parse(layoutFormat, eDateForm)
    eDateDay := strconv.Itoa(eDate.Day())
    eDateMonth := eDate.Month().String()
    eDateYear := strconv.Itoa(eDate.Year())
    eDateFinal := eDateDay + " " + eDateMonth + " " + eDateYear

    // conditional for tech icon
    if c1 == "on" {
        c1          = "fab fa-php me-3"
        Icon1PClass = "d-flex align-items-center col-6"
        Icon1P      = "PHP"
        Icon1Status = "checked"
    }else {
        c1          = ""
        Icon1PClass = ""
        Icon1P      = ""
        Icon1Status = ""
    }

    if c2 == "on" {
        c2          = "fab fa-js-square me-3"
        Icon2PClass = "d-flex align-items-center col-6"
        Icon2P      = "Javascript"
        Icon2Status = "checked"
    }else {
        c2          = ""
        Icon2PClass = ""
        Icon2P      = ""
        Icon2Status = ""
    }

    if c3 == "on" {
        c3          = "fab fa-python me-3"
        Icon3PClass = "d-flex align-items-center col-6"
        Icon3P      = "Python"
        Icon3Status = "checked"
    }else {
        c3          = ""
        Icon3PClass = ""
        Icon3P      = ""
        Icon3Status = ""
    }

    if c4 == "on" {
        c4          = "fab fa-html5 me-3"
        Icon4PClass = "d-flex align-items-center col-6"
        Icon4P      = "HTML"
        Icon4Status = "checked"
    }else {
        c4          = ""
        Icon4PClass = ""
        Icon4P      = ""
        Icon4Status = ""
    }

    // create temporary variable for storing data from input form (array)
    var newProjectCard = Project{
        Pname:          pName,
        Sdate:          sDateForm,
        SdateDetail:    sDateFinal,
        Edate:          eDateForm,
        EdateDetail:    eDateFinal,
        Duration:       "3 Month(s)",
        Description:    description,
        Icon1:          c1,
        Icon1PClass:    Icon1PClass,
        Icon1P:         Icon1P,
        Icon1Status:    Icon1Status,   
        Icon2:          c2,
        Icon2PClass:    Icon2PClass,
        Icon2P:         Icon2P,
        Icon2Status:    Icon2Status, 
        Icon3:          c3,
        Icon3PClass:    Icon3PClass,
        Icon3P:         Icon3P,
        Icon3Status:    Icon3Status, 
        Icon4:          c4,
        Icon4PClass:    Icon4PClass,
        Icon4P:         Icon4P,
        Icon4Status:    Icon4Status, 
    }

    // adding to the global variable
    ProjectCard = append(ProjectCard, newProjectCard)

    // redirect to home
    http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

/*--------------------------------------------------------------
# Edit Card Routing Function
--------------------------------------------------------------*/
func editCard(w http.ResponseWriter, r *http.Request){
    // parsing file
    var tmpl, err = template.ParseFiles("views/project-update.html")

	// (in response) if parsing error
    if err != nil {
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }

    // var for storing & converting type of the id
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

    // update value of id from Data (global var)
    ID[0].Id = id
    
    // defining temporary variable for storing data (array)
    editCard := Project{}

    // looping to find data matches the id
    for i, data := range ProjectCard {
        if i == id {        
            editCard = Project{
                Pname:          data.Pname,
                Sdate:          data.Sdate,
                SdateDetail:    data.SdateDetail,
                Edate:          data.Edate,
                EdateDetail:    data.EdateDetail,
                Duration:       "3 Month(s)",
                Description:    data.Description,
                Icon1:          data.Icon1,
                Icon1PClass:    data.Icon1PClass,
                Icon1P:         data.Icon1P,
                Icon1Status:    data.Icon1Status,
                Icon2:          data.Icon2,
                Icon2PClass:    data.Icon2PClass,
                Icon2P:         data.Icon2P,
                Icon2Status:    data.Icon2Status,
                Icon3:          data.Icon3,
                Icon3PClass:    data.Icon3PClass,
                Icon3P:         data.Icon3P,
                Icon3Status:    data.Icon3Status,
                Icon4:          data.Icon4,
                Icon4PClass:    data.Icon4PClass,
                Icon4P:         data.Icon4P,
                Icon4Status:    data.Icon4Status,
            }
        }
    }

	// mapping data to be displayed
    resp := map[string]interface{}{
        "Data":  Data,
        "Project": editCard,
    }
    
    // execute file
    tmpl.Execute(w, resp)

}

/*--------------------------------------------------------------
# Update Card Routing Function
--------------------------------------------------------------*/
func updateCard(w http.ResponseWriter, r *http.Request){
    // (in request) output data error >> parsing form html
    err := r.ParseForm()

	// (in response) if parsing error
    if err != nil {
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }

    // getting data from input form
    pName       := r.PostForm.Get("pName")
    description := r.PostForm.Get("description")
    c1          := r.Form.Get("c1")
    Icon1PClass := ""
    Icon1P      := ""
    Icon1Status := ""
    c2          := r.Form.Get("c2")
    Icon2PClass := ""
    Icon2P      := ""
    Icon2Status := ""
    c3          := r.Form.Get("c3")
    Icon3PClass := ""
    Icon3P      := ""
    Icon3Status := ""
    c4          := r.Form.Get("c4")
    Icon4PClass := ""
    Icon4P      := ""
    Icon4Status := ""

    // parsing date (change format)
    layoutFormat := "2006-01-02"
    sDateForm    := r.PostForm.Get("sDate")
    eDateForm    := r.PostForm.Get("eDate")

    sDate, _ := time.Parse(layoutFormat, sDateForm)
    sDateDay := strconv.Itoa(sDate.Day())
    sDateMonth := sDate.Month().String()
    sDateYear := strconv.Itoa(sDate.Year())
    sDateFinal := sDateDay + " " + sDateMonth + " " + sDateYear

    eDate, _ := time.Parse(layoutFormat, eDateForm)
    eDateDay := strconv.Itoa(eDate.Day())
    eDateMonth := eDate.Month().String()
    eDateYear := strconv.Itoa(eDate.Year())
    eDateFinal := eDateDay + " " + eDateMonth + " " + eDateYear

    // conditional for tech icon
    if c1 == "on" {
        c1          = "fab fa-php me-3"
        Icon1PClass = "d-flex align-items-center col-6"
        Icon1P      = "PHP"
        Icon1Status = "checked"
    }else {
        c1          = ""
        Icon1PClass = ""
        Icon1P      = ""
        Icon1Status = ""
    }

    if c2 == "on" {
        c2          = "fab fa-js-square me-3"
        Icon2PClass = "d-flex align-items-center col-6"
        Icon2P      = "Javascript"
        Icon2Status = "checked"
    }else {
        c2          = ""
        Icon2PClass = ""
        Icon2P      = ""
        Icon2Status = ""
    }

    if c3 == "on" {
        c3          = "fab fa-python me-3"
        Icon3PClass = "d-flex align-items-center col-6"
        Icon3P      = "Python"
        Icon3Status = "checked"
    }else {
        c3          = ""
        Icon3PClass = ""
        Icon3P      = ""
        Icon3Status = ""
    }

    if c4 == "on" {
        c4          = "fab fa-html5 me-3"
        Icon4PClass = "d-flex align-items-center col-6"
        Icon4P      = "HTML"
        Icon4Status = "checked"
    }else {
        c4          = ""
        Icon4PClass = ""
        Icon4P      = ""
        Icon4Status = ""
    }

    // create temporary variable for storing data from input form (array)
    var updateProjectCard = Project{
        Pname:          pName,
        Sdate:          sDateForm,
        SdateDetail:    sDateFinal,
        Edate:          eDateForm,
        EdateDetail:    eDateFinal,
        Duration:       "3 Month(s)",
        Description:    description,
        Icon1:          c1,
        Icon1PClass:    Icon1PClass,
        Icon1P:         Icon1P,
        Icon1Status:    Icon1Status,   
        Icon2:          c2,
        Icon2PClass:    Icon2PClass,
        Icon2P:         Icon2P,
        Icon2Status:    Icon2Status, 
        Icon3:          c3,
        Icon3PClass:    Icon3PClass,
        Icon3P:         Icon3P,
        Icon3Status:    Icon3Status, 
        Icon4:          c4,
        Icon4PClass:    Icon4PClass,
        Icon4P:         Icon4P,
        Icon4Status:    Icon4Status, 
    }

    // updating data/value
    ProjectCard[ID[0].Id].Pname       = updateProjectCard.Pname
    ProjectCard[ID[0].Id].Sdate       = updateProjectCard.Sdate
    ProjectCard[ID[0].Id].SdateDetail = updateProjectCard.SdateDetail
    ProjectCard[ID[0].Id].Edate       = updateProjectCard.Edate
    ProjectCard[ID[0].Id].EdateDetail = updateProjectCard.EdateDetail
    ProjectCard[ID[0].Id].Duration    = updateProjectCard.Duration
    ProjectCard[ID[0].Id].Description = updateProjectCard.Description
    ProjectCard[ID[0].Id].Icon1       = updateProjectCard.Icon1
    ProjectCard[ID[0].Id].Icon1PClass = updateProjectCard.Icon1PClass
    ProjectCard[ID[0].Id].Icon1P      = updateProjectCard.Icon1P
    ProjectCard[ID[0].Id].Icon1Status = updateProjectCard.Icon1Status   
    ProjectCard[ID[0].Id].Icon2       = updateProjectCard.Icon2
    ProjectCard[ID[0].Id].Icon2PClass = updateProjectCard.Icon2PClass
    ProjectCard[ID[0].Id].Icon2P      = updateProjectCard.Icon2P
    ProjectCard[ID[0].Id].Icon2Status = updateProjectCard.Icon2Status 
    ProjectCard[ID[0].Id].Icon3       = updateProjectCard.Icon3
    ProjectCard[ID[0].Id].Icon3PClass = updateProjectCard.Icon3PClass
    ProjectCard[ID[0].Id].Icon3P      = updateProjectCard.Icon3P
    ProjectCard[ID[0].Id].Icon3Status = updateProjectCard.Icon3Status 
    ProjectCard[ID[0].Id].Icon4       = updateProjectCard.Icon4
    ProjectCard[ID[0].Id].Icon4PClass = updateProjectCard.Icon4PClass
    ProjectCard[ID[0].Id].Icon4P      = updateProjectCard.Icon4P
    ProjectCard[ID[0].Id].Icon4Status = updateProjectCard.Icon4Status

    // redirect to home
    http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

/*--------------------------------------------------------------
# Delete Card Routing Function
--------------------------------------------------------------*/
func deleteCard(w http.ResponseWriter, r *http.Request){
    // var for storing & converting type of the id
    id, _ := strconv.Atoi(mux.Vars(r)["id"])

    // selecting data to be displayed (... for excluding)
    ProjectCard = append(ProjectCard[:id], ProjectCard[id+1:]...)
    http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

/*--------------------------------------------------------------
# Project Blog Routing Function
--------------------------------------------------------------*/
func projectBlog(w http.ResponseWriter, r *http.Request) {
	// parsing file
	var tmpl, err = template.ParseFiles("views/project-blog.html")

	// (in response) if parsing error
    if err != nil {
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }

	// var for storing & converting type of the id
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

    // update value of id from Data (global var)
    ID[0].Id = id

    // defining temporary variable for storing data (array)
    projectBlog := Project{}

    // looping to find data matches the id
    for i, data := range ProjectCard {
        if i == id {        
            projectBlog = Project{
                Pname:          data.Pname,
                Sdate:          data.Sdate,
                SdateDetail:    data.SdateDetail,
                Edate:          data.Edate,
                EdateDetail:    data.EdateDetail,
                Duration:       "3 Month(s)",
                Description:    data.Description,
                Icon1:          data.Icon1,
                Icon1PClass:    data.Icon1PClass,
                Icon1P:         data.Icon1P,
                Icon1Status:    data.Icon1Status,
                Icon2:          data.Icon2,
                Icon2PClass:    data.Icon2PClass,
                Icon2P:         data.Icon2P,
                Icon2Status:    data.Icon2Status,
                Icon3:          data.Icon3,
                Icon3PClass:    data.Icon3PClass,
                Icon3P:         data.Icon3P,
                Icon3Status:    data.Icon3Status,
                Icon4:          data.Icon4,
                Icon4PClass:    data.Icon4PClass,
                Icon4P:         data.Icon4P,
                Icon4Status:    data.Icon4Status,
            }
        }
    }

	// mapping data to be displayed
    resp := map[string]interface{}{
        "Data":  Data,
        "Project": projectBlog,
    }
    // execute file
    tmpl.Execute(w, resp)

}

/*--------------------------------------------------------------
# Contact Me Routing Function
--------------------------------------------------------------*/
func contactMe(w http.ResponseWriter, r *http.Request) {
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