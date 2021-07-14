var add_modal_open = document.getElementById("add_task");
var add_modal_close = document.getElementsByClassName("add_modal_close");
var add_modal = document.getElementById("add_modal");

var contents = document.getElementsByClassName("contents");
var task_id = document.getElementsByClassName("task_id")
function set_task_id(value){
  for (var i=0; i < task_id.length; i++){
    task_id[i].value = value
  }
}

add_modal_open.addEventListener("click", function(){
  // task_id.value = "";
  set_task_id("")
  add_modal.style.display = "block";
});

for (var i=0; i < add_modal_close.length; i++){
  add_modal_close[i].addEventListener("click", function() {
    add_modal.style.display = "none";
  });
}

for (var i=0; i < contents.length; i++){
  contents[i].addEventListener("click", function(){
    // task_id.value = this.id.split("_")[1];
    set_task_id(this.id.split("_")[1])
    add_modal.style.display = "block";
  });
}


