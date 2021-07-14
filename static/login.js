var singin = document.getElementById("singin");
var singup = document.getElementById("singup");

var singin_form = document.getElementById("singin-form");
var singup_form = document.getElementById("singup-form");

var singin_color = singin.style.color;
var singup_color = singup.style.color;

function set_color(){
    singin_color = singin.style.color;
    singup_color = singup.style.color;
    console.log(singup_color)
}

singin.addEventListener("click", function(){
    singin_form.style.display = "block";
    singup_form.style.display = "none";
    singin.style.color = "#ECF0F5";
    singup.style.color = "#6C6C6C";
    set_color();
})

singup.addEventListener("click", function(){
    singin_form.style.display = "none";
    singup_form.style.display = "block";
    singin.style.color = "#6C6C6C";
    singup.style.color = "#ECF0F5";
    set_color();
})

singin.addEventListener("mouseover", function(){
    singin.style.color = "#ECF0F5";
    singup.style.color = "#6C6C6C";
})

singup.addEventListener("mouseover", function(){
    singin.style.color = "#6C6C6C";
    singup.style.color = "#ECF0F5";
})

singin.addEventListener("mouseout", function(){
    singin.style.color = singin_color;
    singup.style.color = singup_color;
})

singup.addEventListener("mouseout", function(){
    singin.style.color = singin_color;
    singup.style.color = singup_color;
})
