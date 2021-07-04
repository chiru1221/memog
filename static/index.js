var modal_open = document.getElementById("add_task");
var modal_close = document.getElementsByClassName("modal_close");
var add_modal = document.getElementById("add_modal");

var contents = document.getElementsByClassName("contents");
var task_id = document.getElementById("task_id")

modal_open.addEventListener("click", function(){
  task_id.value = "";
  add_modal.style.display = "block";
});

for (var i=0; i < modal_close.length; i++){
  modal_close[i].addEventListener("click", function() {
    add_modal.style.display = "none";
  });
}

for (var i=0; i < contents.length; i++){
  contents[i].addEventListener("click", function(){
    task_id.value = this.id.split("_")[1];
    add_modal.style.display = "block";
  });
}

