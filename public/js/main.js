AOS.init()

/*--------------------------------------------------------------
# Contact Me Form
--------------------------------------------------------------*/

//to get data
function getData () {

    //to store data in variable
    let name = document.getElementById("name").value
    let email = document.getElementById('email').value
    let phone = document.getElementById('phone').value
    let subject = document.getElementById('subject').value
    let comments = document.getElementById('comments').value

    //to animate alert
    const elem = document.getElementById("alert");   
    let pos = 0;
    let id = setInterval(frame, 5);
    function frame() {
        if (pos == 20) {
        clearInterval(id);
        elem.style.opacity = 1;
        } else {
        pos++; 
        elem.style.opacity = 0.3;
        elem.style.top = pos + "px";  
        }
    }

    //to initialize alert color
    document.getElementById("alert").className = "alert alert-danger";
   

    //to validate data
    if( name == ""){
        //to all alert
        document.getElementById("alert").innerHTML = "Name cannot be empty";
    }

    else if( email == ""){
        //to all alert
        document.getElementById("alert").innerHTML = "Email cannot be empty";
    }

    else if( phone == ""){
        //to all alert
        document.getElementById("alert").innerHTML = "Phone number cannot be empty";
    }

    else if( subject == "--Subject--"){
        //to all alert
        document.getElementById("alert").innerHTML = "Subject cannot be empty";
    }

    else if( comments == ""){
        //to all alert
        document.getElementById("alert").innerHTML = "Comments cannot be empty";
    }

    else {

        //to mail data
        const defEmail = "bilqistaz27@gmail.com"
        let ebody = "Dear Ms Bilqist," + "%0D%0A" + "%0D%0A" + "My name is " + name + "." + "%0D%0A" + "%0D%0A" + comments + "%0D%0A" + "%0D%0A" + "You can contact me here: " + phone

        console.log(ebody)
        window.location.href = `mailto:${defEmail}?subject=${subject}&body=${ebody}`

        //to change alert color
        document.getElementById("alert").className = "alert alert-success";

        //to all alert
        document.getElementById("alert").innerHTML = "Email has been sent"

        //to reset form only without the whole page
        let form = document.getElementById("form");
        form.reset();

    }

}

