
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
    "math"

    "chapter02/connection"
    "context"

    "github.com/gorilla/mux"
)

/*--------------------------------------------------------------
# Global var (object to send and be displayed in html)
--------------------------------------------------------------*/
var Data = map[string]interface{}{
    "IsLogin": true,
    "Alert": "sd",
}

type Project struct {
    Id              int
    Pname           string
    Sdate           time.Time
    SdateDetail     string
    Edate           time.Time
    EdateDetail     string
    Duration        string
    Description     string
    Technologies    []string
    Alert           string
}

/*--------------------------------------------------------------
# Main Routing Function
--------------------------------------------------------------*/

func main() {
	// create new router
    route := mux.NewRouter()

    // calling connection
    connection.DatabaseConnect()

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

    // query
    sql := "SELECT id, pname, sdate, edate, description, technologies FROM public.tb_project ORDER BY id ASC;"

    // func to execute query
    rows, _ := connection.Conn.Query(context.Background(), sql)

	var result []Project
    for rows.Next() {
        var each = Project{}

        err := rows.Scan(&each.Id, &each.Pname, &each.Sdate, &each.Edate, &each.Description, &each.Technologies)
        if err != nil {
            fmt.Println(err.Error())
            return
        }

        // parsing date (change format)
        sDateDay := strconv.Itoa(each.Sdate.Day())
        sDateMonth := each.Sdate.Month().String()
        sDateYear := strconv.Itoa(each.Sdate.Year())
        each.SdateDetail = sDateDay + " " + sDateMonth + " " + sDateYear

        eDateDay := strconv.Itoa(each.Edate.Day())
        eDateMonth := each.Edate.Month().String()
        eDateYear := strconv.Itoa(each.Edate.Year())
        each.EdateDetail = eDateDay + " " + eDateMonth + " " + eDateYear

        // get duration
        diff := (each.Edate.Sub(each.Sdate)).Hours() // hour

        DurYear := math.Round(diff / (12 * 30 * 24)) //in Year
        DurMonth := math.Round(diff / (30 * 24)) //in Month
        DurWeek := math.Round(diff / (7 * 24)) //in Day
        DurDay := math.Round(diff / (24)) //in Day

        if DurYear > 0 {
            each.Duration =  strconv.FormatFloat(DurYear, 'f', 0, 64) + " Year(s)"
        }else if DurMonth > 0 {
            each.Duration = strconv.FormatFloat(DurMonth, 'f', 0, 64) + " Month(s)"
        }else if DurWeek > 0 {
            each.Duration = strconv.FormatFloat(DurWeek, 'f', 0, 64) + " Week(s)"
        }else if DurDay > 0 {
            each.Duration = strconv.FormatFloat(DurDay, 'f', 0, 64) + " Days(s)"
        }
           
        result = append(result, each)
    }

    // mapping data to be displayed
    resp := map[string]interface{}{
        "Data":  Data,
        "Project": result,
    }

    fmt.Println(resp)
	// execute file
    tmpl.Execute(w, resp)
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
    sDate       := r.PostForm.Get("sDate")
    eDate       := r.PostForm.Get("eDate")
    description := r.PostForm.Get("description")
    icon        := r.Form["tech"]

    // parsing date (change format)
    layoutFormat := "2006-01-02"
    sDateForm, _ := time.Parse(layoutFormat, sDate)
    eDateForm, _ := time.Parse(layoutFormat, eDate)
    
    // create temporary variable for storing data from input form (array)
    var newProjectCard = Project{
        Pname:          pName,
        Sdate:          sDateForm,
        Edate:          eDateForm,
        Description:    description,
        Technologies:   icon,

    }

    fmt.Println(newProjectCard)

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

	// mapping data to be displayed
    resp := map[string]interface{}{
        "Data":  Data,
        // "Project": editCard,
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

    // redirect to home
    http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

/*--------------------------------------------------------------
# Delete Card Routing Function
--------------------------------------------------------------*/
func deleteCard(w http.ResponseWriter, r *http.Request){
    
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

    // var for storing id
	id := (mux.Vars(r)["id"])
    
    // query
    sql := "SELECT id, pname, sdate, edate, description, technologies FROM public.tb_project WHERE id=$1;"

	row := connection.Conn.QueryRow(context.Background(), sql, id)

	var project Project
	err = row.Scan(&project.Id, &project.Pname, &project.Sdate, &project.Edate, &project.Description, &project.Technologies)
	if err != nil {
		fmt.Println("Unable to retrieve data:", err)
		return
	}

    // parsing date (change format)
    sDateDay := strconv.Itoa(project.Sdate.Day())
    sDateMonth := project.Sdate.Month().String()
    sDateYear := strconv.Itoa(project.Sdate.Year())
    project.SdateDetail = sDateDay + " " + sDateMonth + " " + sDateYear

    eDateDay := strconv.Itoa(project.Edate.Day())
    eDateMonth := project.Edate.Month().String()
    eDateYear := strconv.Itoa(project.Edate.Year())
    project.EdateDetail = eDateDay + " " + eDateMonth + " " + eDateYear

    // get duration
    diff := (project.Edate.Sub(project.Sdate)).Hours() // hour
    DurYear := math.Round(diff / (12 * 30 * 24)) //in Year
    DurMonth := math.Round(diff / (30 * 24)) //in Month
    DurWeek := math.Round(diff / (7 * 24)) //in Day
    DurDay := math.Round(diff / (24)) //in Day

    if DurYear > 0 {
        project.Duration =  strconv.FormatFloat(DurYear, 'f', 0, 64) + " Year(s)"
    }else if DurMonth > 0 {
        project.Duration = strconv.FormatFloat(DurMonth, 'f', 0, 64) + " Month(s)"
    }else if DurWeek > 0 {
        project.Duration = strconv.FormatFloat(DurWeek, 'f', 0, 64) + " Week(s)"
    }else if DurDay > 0 {
        project.Duration = strconv.FormatFloat(DurDay, 'f', 0, 64) + " Days(s)"
    }

	result := []Project{project}

	// mapping data to be displayed
    resp := map[string]interface{}{
        "Data":  Data,
        "Project": result,
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