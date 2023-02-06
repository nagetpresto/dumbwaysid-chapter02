
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
    IsIcon1     bool
    IsIcon2     bool
    IsIcon3     bool
    IsIcon4     bool
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
        IsIcon1:         true,
        IsIcon2:         true,
        IsIcon3:         true,
        IsIcon4:         true,
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

    // getting icon
    tech := []string{r.Form.Get("c1"),r.Form.Get("c2"),r.Form.Get("c3"),r.Form.Get("c4")}

    IsIcon1 := false
    IsIcon2 := false
    IsIcon3 := false
    IsIcon4 := false

    // conditional for tech icon
    if tech[0] != "" {
        IsIcon1 = true
    }
    if tech[1] != "" {
        IsIcon2 = true
    }
    if tech[2] != "" {
        IsIcon3 = true
    }
    if tech[3] != "" {
        IsIcon4 = true
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
        IsIcon1:        IsIcon1,
        IsIcon2:        IsIcon2,
        IsIcon3:        IsIcon3,
        IsIcon4:        IsIcon4,

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
                IsIcon1:        data.IsIcon1,
                IsIcon2:        data.IsIcon2,
                IsIcon3:        data.IsIcon3,
                IsIcon4:        data.IsIcon4,
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

    // getting icon
    tech := []string{r.Form.Get("c1"),r.Form.Get("c2"),r.Form.Get("c3"),r.Form.Get("c4")}

    IsIcon1 := false
    IsIcon2 := false
    IsIcon3 := false
    IsIcon4 := false

    // conditional for tech icon
    if tech[0] != "" {
        IsIcon1 = true
    }
    if tech[1] != "" {
        IsIcon2 = true
    }
    if tech[2] != "" {
        IsIcon3 = true
    }
    if tech[3] != "" {
        IsIcon4 = true
    }

    // updating data/value
    ProjectCard[ID[0].Id].Pname       = pName
    ProjectCard[ID[0].Id].Sdate       = sDateForm
    ProjectCard[ID[0].Id].SdateDetail = sDateFinal
    ProjectCard[ID[0].Id].Edate       = eDateForm
    ProjectCard[ID[0].Id].EdateDetail = eDateFinal
    ProjectCard[ID[0].Id].Duration    = "3 Month(s)"
    ProjectCard[ID[0].Id].Description = description
    ProjectCard[ID[0].Id].IsIcon1     = IsIcon1
    ProjectCard[ID[0].Id].IsIcon2     = IsIcon2
    ProjectCard[ID[0].Id].IsIcon3     = IsIcon3
    ProjectCard[ID[0].Id].IsIcon4     = IsIcon4

    // redirect to home
    http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

/*--------------------------------------------------------------
# Delete Card Routing Function
--------------------------------------------------------------*/
func deleteCard(w http.ResponseWriter, r *http.Request){
    // var for storing & converting type of the id
    id, _ := strconv.Atoi(mux.Vars(r)["id"])

    // update value of id from Data (global var)
    ID[0].Id = id

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
                IsIcon1:        data.IsIcon1,
                IsIcon2:        data.IsIcon2,
                IsIcon3:        data.IsIcon3,
                IsIcon4:        data.IsIcon4,
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