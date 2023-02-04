/*--------------------------------------------------------------
# Add Project Form
--------------------------------------------------------------*/


//global array to store data
let project = []

//to get data
function addProject () {

    //to store data in variable
    let title = document.getElementById("pName").value
    let startDate = document.getElementById("sDate").value
    let endDate = document.getElementById("eDate").value
    let description = document.getElementById("description").value
    let c1 = document.getElementById("c1").checked
    let c2 = document.getElementById("c2").checked
    let c3 = document.getElementById("c3").checked
    let c4 = document.getElementById("c4").checked
    let image = document.getElementById('file').files

    let icon1 = (c1 == true) ? '<i class="fab fa-php"></i>' : '';
    let icon2 = (c2 == true) ? '<i class="fab fa-js-square"></i>' : '';
    let icon3 = (c3 == true) ? '<i class="fab fa-python"></i>' : '';
    let icon4 = (c4 == true) ? '<i class="fab fa-html5"></i>' : '';

    //to animate alert
    const elem = document.getElementById("alert");   
    let pos = 20;
    let id = setInterval(frame, 5);
    function frame() {
        if (pos == 53) {
        clearInterval(id);
        elem.style.opacity = 1;
        } else {
        pos++; 
        elem.style.opacity = 0.3;
        elem.style.top = pos + "px";  
        }
    }

    //to initialize alert color
    document.getElementById("alert").style.backgroundColor = "rgb(252, 179, 179)";
    document.getElementById("alert").style.color = "#965656";

    //to validate data
    if( title == ""){
        //to call alert
        document.getElementById("alert").innerHTML = "Project Name cannot be empty";
    }

    else if (startDate == ""){
        //to call alert
        document.getElementById("alert").innerHTML = "Start Date cannot be empty";
    }

    else if (endDate == ""){
        //to call alert
        document.getElementById("alert").innerHTML = "End Date cannot be empty";
    }

    else if (description == ""){
        //to call alert
        document.getElementById("alert").innerHTML = "Description cannot be empty";
    }

    else if (icon1 == false && icon2 == false && icon3 == false && icon4 == false){
        //to call alert
        document.getElementById("alert").innerHTML = "Technology cannot be empty";
    }

    else if (image.length == 0){
        //to call alert
        document.getElementById("alert").innerHTML = "Image cannot be empty";
    }

    else {
        // Convert spesific image to blob object
        image = URL.createObjectURL(image[0])

        // to store icon
        let icon = [
            icon1,
            icon2,
            icon3,
            icon4
        ]
        
        // to get duration
        let startDate = new Date(document.getElementById("sDate").value)
        let endDate = new Date(document.getElementById("eDate").value)
        let diff = endDate.getTime() - startDate.getTime() //in milisec
        console.log(diff)
        console.log(startDate)
        console.log(endDate)

        if (diff < 0) {
            //to call alert
            document.getElementById("alert").innerHTML = "Duration is invalid";
            duration = ""
        }
        else {
            let durYear = Math.floor(diff / (12 * 30 * 24* 3600 * 1000)) //in year
            let durMonth = Math.floor(diff / (30 * 24* 3600 * 1000)) //in month
            let durWeek = Math.floor(diff / (7 * 24* 3600 * 1000)) //in week
            let durDay = Math.floor(diff / (24* 3600 * 1000)) //in day

            if (durYear > 0) {
                duration = durYear + " Year(s)"
            }
            else if (durMonth > 0 ) {
                duration = durMonth + " Month(s)"
            }
            else if (durWeek > 0 ) {
                duration = durWeek + " Week(s)"
            }
            else if (durDay > 0 ) {
                duration = durDay + " Day(s)"
            }

            // to filter icon
            let iconFil = icon.filter(elem => elem[1]).reverse()
            console.log(iconFil)

            // to store object
            let dataBlog = {
                image,
                title,
                description,
                iconFil,
                duration
            }

            // to add object to the global array
            project.push(dataBlog)
            console.log(project)

            //to call function
            showProject()
        }

    }
}

//to add project
function showProject () {   

    document.getElementById("project-section").innerHTML = ''
    for(let i=0; i < project.length; i++) {

        if (project[i].iconFil[0] == undefined) {
            project[i].iconFil[0] = ""
        }
        if (project[i].iconFil[1] == undefined) {
            project[i].iconFil[1] = ""
        }
        if (project[i].iconFil[2] == undefined) {
            project[i].iconFil[2] = ""
        }
        if (project[i].iconFil[3] == undefined) {
            project[i].iconFil[3] = ""
        }
        if (project[i].iconFil[4] == undefined) {
            project[i].iconFil[4] = ""
        }

        // for (let a=0; a < 4; a++){
        //     if (project[i].iconFil[a] == undefined) {
        //         project[i].iconFil[a] = ""
        //     }
        // }


        document.getElementById("project-section").innerHTML += `
            <div class="card-project">
            <!-- Project Image -->
            <div class="image-container">
                <img class="image" src="${project[i].image}" alt="Project Image">
            </div>
            <!-- End Project Image -->

            <!-- Project Content -->
            <div class="card-project-content">
                <div class="upper-content">
                    <!-- Project Title -->
                <div class="p-title" id="title">
                    <a href="project-blog.html">${project[i].title}</a>
                </div>
                <!-- Project Title -->

                <!-- Project Period -->
                <div class="p-detail" id="period">
                    Duration: ${project[i].duration}
                </div>
                <!-- End Project Period -->

                <!-- Project Description -->
                <div class="p-desc" id="description">${project[i].description}</div>
                <!-- Project Description -->
            </div>

            <div class="bottom-content">
                <!-- Project Icon -->
                <span class="p-icon">
                    <div class="c1" id="c1">${project[i].iconFil[0]}</div>
                    <div class="c2" id="c2">${project[i].iconFil[1]}</div>
                    <div class="c3" id="c3">${project[i].iconFil[2]}</div>
                    <div class="c4" id="c4">${project[i].iconFil[3]}</div>
                </span>
                <!-- End Project Icon  -->

                <!-- Project Button -->
                <div class="p-button">
                    <button type="button" name="edit" id="edit" class="btn-project">Edit</button>
                    <button type="button" name="delete" id="delete" class="btn-project">Delete</button>
                </div>
                <!-- End Project Button -->
            </div>
            
            </div>
            <!-- End Project Content -->
            
        </div>
        `
    }

    //to change alert color
    document.getElementById("alert").style.backgroundColor = "rgb(185, 245, 170)";
    document.getElementById("alert").style.color = "rgb(75, 110, 66)";

    //to all alert
    document.getElementById("alert").innerHTML = "Project has been added"

    //to reset form only without the whole page
    let form = document.getElementById("form");
    form.reset();

}