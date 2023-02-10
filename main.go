
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

    "golang.org/x/crypto/bcrypt"
    "github.com/gorilla/sessions"
    "github.com/gorilla/mux"
)

/*--------------------------------------------------------------
# Global var (object to send and be displayed in html)
--------------------------------------------------------------*/
var Data = map[string]interface{}{
    "IsLogin"   : true,
    "Id"        : 1,
    "Name"      : "Bilkis",
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
}

type Users struct {
    Id       int
    Name     string
    Email    string
    Password string
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
    route.HandleFunc("/edit-card/{id}", updateCard).Methods("POST")
    route.HandleFunc("/delete-card/{id}", deleteCard).Methods("GET")
	route.HandleFunc("/project/{id}", projectBlog).Methods("GET")
	route.HandleFunc("/contact-me", contactMe).Methods("GET")
    route.HandleFunc("/register", register).Methods("GET")
    route.HandleFunc("/register", registerPost).Methods("POST")
    route.HandleFunc("/login", login).Methods("GET")
    route.HandleFunc("/login", loginPost).Methods("POST")

	// starting server
    fmt.Println("Server running on port 5000")
    http.ListenAndServe("localhost:5000", route)
}

/*--------------------------------------------------------------
# Parsing Time Function
--------------------------------------------------------------*/
func ParsingTime(startDate, endDate time.Time) (string, string, error){
    var SdateDetail string
    var EdateDetail string

    sDateDay := strconv.Itoa(startDate.Day())
    sDateMonth := startDate.Month().String()
    sDateYear := strconv.Itoa(startDate.Year())
    SdateDetail = sDateDay + " " + sDateMonth + " " + sDateYear

    eDateDay := strconv.Itoa(endDate.Day())
    eDateMonth := endDate.Month().String()
    eDateYear := strconv.Itoa(endDate.Year())
    EdateDetail = eDateDay + " " + eDateMonth + " " + eDateYear

    return SdateDetail, EdateDetail, nil
}
/*--------------------------------------------------------------
# Get Duration Function
--------------------------------------------------------------*/

func GetDuration(startDate, endDate time.Time) (string, error) {
    diff := (endDate.Sub(startDate)).Hours() // hour

    DurDay := math.Floor(diff / (24)) //in Day
    DurWeek := math.Floor(DurDay / (7)) //in Week
    DurMonth := math.Floor(DurDay / (30)) //in Month
    DurYear := math.Floor(DurDay / (12 * 30)) //in Year

    var Duration string
    if DurYear > 0 {
        Duration = strconv.FormatFloat(DurYear, 'f', 0, 64) + " Year(s)"
    }else if DurMonth > 0 {
        Duration = strconv.FormatFloat(DurMonth, 'f', 0, 64) + " Month(s)"
    }else if DurWeek > 0 {
        Duration = strconv.FormatFloat(DurWeek, 'f', 0, 64) + " Week(s)"
    }else if DurDay > 0 {
        Duration = strconv.FormatFloat(DurDay, 'f', 0, 64) + " Days(s)"
    }

    return Duration, nil
}

/*--------------------------------------------------------------
# Technologies Chunk Function
--------------------------------------------------------------*/
func TechChunk(Technologies []string) (bool, bool, bool, bool, error){
    IsPHP := false
    IsJS := false
    IsPYTHON := false
    IsHTML := false
    for _, data := range Technologies {
        if data == "php" {
            IsPHP = true
        }
        if data == "js-square" {
            IsJS = true
        }
        if data == "python" {
            IsPYTHON = true
        }
        if data == "html5" {
            IsHTML = true
        }
    }

    return IsPHP, IsJS, IsPYTHON, IsHTML, nil
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

    sql := ""
    var store = sessions.NewCookieStore([]byte("SESSION_ID"))
    session, _ := store.Get(r, "SESSION_ID")

    if session.Values["IsLogin"] != true {
        Data["IsLogin"] = false
        sql = "SELECT id, pname, sdate, edate, description, technologies FROM public.tb_project ORDER BY id ASC;"
    }else {
        Data["IsLogin"] = session.Values["IsLogin"].(bool)
        Data["Name"] = session.Values["Name"].(string)
        Data["Id"] = session.Values["Id"].(int)
        sql = fmt.Sprintf("SELECT public.tb_project.id, pname, sdate, edate, description, technologies FROM public.tb_users JOIN public.tb_project ON public.tb_users.id = public.tb_project.user_id WHERE public.tb_users.id=%d;",Data["Id"])
            // "SELECT id, pname, sdate, edate, description, technologies FROM public.tb_project WHERE user_id=%d ORDER BY id ASC;", Data["Id"])
        
    }

    // func to execute query
    rows, _ := connection.Conn.Query(context.Background(), sql)

	var result []Project
    for rows.Next() {
        var each = Project{}

        err := rows.Scan(&each.Id, &each.Pname, &each.Sdate, &each.Edate, &each.Description, &each.Technologies)
        if err != nil {
            http.Error(w, "Unable to retrieve data: " + err.Error(), http.StatusBadRequest)
            return
        }

        // output parsing tima function
        SdateDetail, EdateDetail, err := ParsingTime(each.Sdate, each.Edate)
        if err != nil {
            http.Error(w, err.Error(),http.StatusInternalServerError)
            return
        }
        each.SdateDetail = SdateDetail
        each.EdateDetail = EdateDetail

        // output get duration function
        Duration, err := GetDuration(each.Sdate, each.Edate)
        if err != nil {
            http.Error(w, err.Error(),http.StatusInternalServerError)
            return
        }
        each.Duration = Duration
        result = append(result, each)
    }

    // mapping data to be displayed
    resp := map[string]interface{}{
        "Data":  Data,
        "Project": result,
    }

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
    
    var store = sessions.NewCookieStore([]byte("SESSION_ID"))
    session, _ := store.Get(r, "SESSION_ID")

    if session.Values["IsLogin"] != true {
        Data["IsLogin"] = false
        Data["Id"] = session.Values["Id"]
    } else {
        Data["IsLogin"] = session.Values["IsLogin"].(bool)
        Data["Name"] = session.Values["Name"].(string)
        Data["Id"] = session.Values["Id"].(int)
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

    // duration validation
    diff := (eDateForm.Sub(sDateForm)).Hours() // hour

    if diff < 0 {
        http.Error(w, "Duration is Invalid", http.StatusBadRequest)
		return
    }

    sql := "INSERT INTO public.tb_project (pname, sdate, edate, description, technologies, user_id) VALUES ($1, $2, $3, $4, $5, $6)"
	_, errr := connection.Conn.Exec(context.Background(), sql, pName, sDateForm, eDateForm, description, icon, Data["Id"])
	if errr != nil {
        http.Error(w, "Unable to insert data: " + errr.Error(),  http.StatusBadRequest)
        return
    }

	fmt.Println("Data inserted successfully.")

    // redirect 
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

    // var for storing id
	id := (mux.Vars(r)["id"])
    
    // selecting query
    sql := "SELECT id, pname, sdate, edate, description, technologies FROM public.tb_project WHERE id=$1;"
	row := connection.Conn.QueryRow(context.Background(), sql, id)

	var project Project
	err = row.Scan(&project.Id, &project.Pname, &project.Sdate, &project.Edate, &project.Description, &project.Technologies)
	if err != nil {
        http.Error(w, "Unable to retrieve data: " + err.Error(),  http.StatusBadRequest)
        return
    }

    // parsing date (change format)
    sDate := project.Sdate.Format("2006-01-02")
    eDate := project.Edate.Format("2006-01-02")

    // output TechChunk function
    PHP, JS, PYTHON, HTML, err := TechChunk(project.Technologies)
    if err != nil {
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }
    IsPHP      := PHP
    IsJS       := JS
    IsPYTHON   := PYTHON
    IsHTML     := HTML

	result := []Project{project}

	// mapping data to be displayed
    resp := map[string]interface{}{
        "Data":  Data,
        "Id": id,
        "Project": result,
        "sDate": sDate,
        "eDate": eDate,
        "IsPHP" : IsPHP,
        "IsJS" : IsJS,
        "IsPYTHON" : IsPYTHON,
        "IsHTML" : IsHTML,
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

    // var for storing id
	id := (mux.Vars(r)["id"])

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

    // duration validation
    diff := (eDateForm.Sub(sDateForm)).Hours() // hour

    if diff < 0 {
        http.Error(w, "Duration is Invalid", http.StatusBadRequest)
		return
    }

    // query to update
    sql := "UPDATE public.tb_project SET pname=$1, sdate=$2, edate=$3, description=$4, technologies=$5 WHERE id=$6"
	// executing query
    _, errr := connection.Conn.Exec(context.Background(), sql, pName, sDateForm, eDateForm, description, icon, id)
	if errr != nil {
        http.Error(w, "Unable to update data: " + errr.Error(),  http.StatusBadRequest)
        return
    }

	fmt.Println("Data updated successfully.")

    // redirect to home
    http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

/*--------------------------------------------------------------
# Delete Card Routing Function
--------------------------------------------------------------*/
func deleteCard(w http.ResponseWriter, r *http.Request){
    // var for storing id
	id := (mux.Vars(r)["id"])
    
    // query
    sql := "DELETE FROM public.tb_project WHERE id=$1;"
    
	_, errr := connection.Conn.Exec(context.Background(), sql, id)
	if errr != nil {
        http.Error(w, "Unable to delete data: " + errr.Error(),  http.StatusBadRequest)
        return
    }

	fmt.Println("Data deleted successfully.")
    
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

    var store = sessions.NewCookieStore([]byte("SESSION_ID"))
    session, _ := store.Get(r, "SESSION_ID")

    if session.Values["IsLogin"] != true {
        Data["IsLogin"] = false
        Data["Id"] = session.Values["Id"]
    } else {
        Data["IsLogin"] = session.Values["IsLogin"].(bool)
        Data["Name"] = session.Values["Name"].(string)
        Data["Id"] = session.Values["Id"].(int)
    }

    // var for storing id
	id := (mux.Vars(r)["id"])
    
    // query
    sql := "SELECT id, pname, sdate, edate, description, technologies FROM public.tb_project WHERE id=$1;"

	row := connection.Conn.QueryRow(context.Background(), sql, id)

	var project Project
	err = row.Scan(&project.Id, &project.Pname, &project.Sdate, &project.Edate, &project.Description, &project.Technologies)
	if err != nil {
        http.Error(w, "Unable to retrieve data: " + err.Error(),  http.StatusBadRequest)
        return
    }

    // output parsing tima function
    SdateDetail, EdateDetail, err := ParsingTime(project.Sdate, project.Edate)
    if err != nil {
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }
    project.SdateDetail = SdateDetail
    project.EdateDetail = EdateDetail

    // output get duration function
    Duration, err := GetDuration(project.Sdate, project.Edate)
    if err != nil {
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }
    project.Duration = Duration

    // output TechChunk function
    PHP, JS, PYTHON, HTML, err := TechChunk(project.Technologies)
    if err != nil {
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }
    IsPHP      := PHP
    IsJS       := JS
    IsPYTHON   := PYTHON
    IsHTML     := HTML

	result := []Project{project}

	// mapping data to be displayed
    resp := map[string]interface{}{
        "Data":  Data,
        "Project": result,
        "IsPHP" : IsPHP,
        "IsJS" : IsJS,
        "IsPYTHON" : IsPYTHON,
        "IsHTML" : IsHTML,
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

    var store = sessions.NewCookieStore([]byte("SESSION_ID"))
    session, _ := store.Get(r, "SESSION_ID")

    if session.Values["IsLogin"] != true {
        Data["IsLogin"] = false
        Data["Id"] = session.Values["Id"]
    } else {
        Data["IsLogin"] = session.Values["IsLogin"].(bool)
        Data["Name"] = session.Values["Name"].(string)
        Data["Id"] = session.Values["Id"].(int)
    }

	// execute file
    tmpl.Execute(w, Data)
}

/*--------------------------------------------------------------
# Register Routing Function
--------------------------------------------------------------*/
func register(w http.ResponseWriter, r *http.Request) {
	// parsing file
    var tmpl, err = template.ParseFiles("views/register.html")

	// (in response) if parsing error
    if err != nil {
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }

    var store = sessions.NewCookieStore([]byte("SESSION_ID"))
    session, _ := store.Get(r, "SESSION_ID")

    if session.Values["IsLogin"] != true {
        Data["IsLogin"] = false
        Data["Id"] = session.Values["Id"]
    } else {
        Data["IsLogin"] = session.Values["IsLogin"].(bool)
        Data["Name"] = session.Values["Name"].(string)
        Data["Id"] = session.Values["Id"].(int)
    }

	// execute file
    tmpl.Execute(w, Data)
}

/*--------------------------------------------------------------
# Register POST Routing Function
--------------------------------------------------------------*/
func registerPost(w http.ResponseWriter, r *http.Request){
    // (in request) output data error >> parsing form html
    err := r.ParseForm()

	// (in response) if parsing error
    if err != nil {
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }

    // getting data from input form
    name            := r.PostForm.Get("name")
    email           := r.PostForm.Get("email")
    password        := r.PostForm.Get("password")
    passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

    // query
    sql := "INSERT INTO public.tb_users (name, email, password) VALUES ($1, $2, $3)"
	_, errr := connection.Conn.Exec(context.Background(), sql, name, email, passwordHash)
	if errr != nil {
        http.Error(w, "Unable to register data: " + errr.Error(),  http.StatusBadRequest)
        return
    }

	fmt.Println("Data registered successfully.")

    // redirect 
    http.Redirect(w, r, "/login", http.StatusMovedPermanently)
}

/*--------------------------------------------------------------
# Login Routing Function
--------------------------------------------------------------*/
func login(w http.ResponseWriter, r *http.Request) {
	// parsing file
    var tmpl, err = template.ParseFiles("views/login.html")

	// (in response) if parsing error
    if err != nil {
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }

    var store = sessions.NewCookieStore([]byte("SESSION_ID"))
    session, _ := store.Get(r, "SESSION_ID")

    if session.Values["IsLogin"] != true {
        Data["IsLogin"] = false
        Data["Id"] = session.Values["Id"]
    } else {
        Data["IsLogin"] = session.Values["IsLogin"].(bool)
        Data["Name"] = session.Values["Name"].(string)
        Data["Id"] = session.Values["Id"].(int)
    }

	// execute file
    tmpl.Execute(w, Data)
}

/*--------------------------------------------------------------
# Login POST Routing Function
--------------------------------------------------------------*/
func loginPost(w http.ResponseWriter, r *http.Request){
    // (in request) output data error >> parsing form html
    err := r.ParseForm()

	// (in response) if parsing error
    if err != nil {
        http.Error(w, err.Error(),http.StatusInternalServerError)
        return
    }

    // getting data from input form
    email       := r.PostForm.Get("email")
    password    := r.PostForm.Get("password")

    // query for selecting to compare
    sql := "SELECT * FROM public.tb_users WHERE email=$1;"
	row := connection.Conn.QueryRow(context.Background(), sql, email)

    // storing data
	var user Users
	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
    // matching email
    if err != nil {
        http.Error(w, "Email does not match: " + err.Error(), http.StatusBadRequest)
        return
    }

    // matching password
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        http.Error(w, "Password does not match: " + err.Error(), http.StatusBadRequest)
        return
    }

    // setting up the session store (secrey key)
    store := sessions.NewCookieStore([]byte("SESSION_ID"))

    // getting the session map(session name)
    session, _ := store.Get(r, "SESSION_ID")

    // assigning values to session.Values map
    session.Values["IsLogin"] = true
    session.Values["Id"] = user.Id
    session.Values["Name"] = user.Name
    session.Options.MaxAge = 3600 

    // // adding the flash message to the session
    // session.AddFlash("Login success", "message")

    // saving the session
    session.Save(r, w)
    
    http.Redirect(w, r, "/", http.StatusMovedPermanently)
}