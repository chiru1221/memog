var add_modal_open = document.getElementById("add_task");
var add_modal_close = document.getElementsByClassName("add_modal_close");
var add_modal = document.getElementById("add_modal");

var del_modal_open = document.getElementsByClassName("contents_del");
var del_modal_close = document.getElementsByClassName("del_modal_close");
var del_modal = document.getElementById("del_modal")

var contents = document.getElementsByClassName("contents");
var task_id = document.getElementsByClassName("task_id");
var modal = document.getElementsByClassName("modal");

function set_task_id(value){
  for (var i=0; i < task_id.length; i++){
    task_id[i].value = value;
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
  contents[i].addEventListener("click", function(e){
    if (e.target != e.currentTarget){
      return ;
    }
    set_task_id(this.id.split("_")[1])
    add_modal.style.display = "block";
  });

}

for (var i=0; i < del_modal_open.length; i++){
  del_modal_open[i].addEventListener("click", function() {
    set_task_id(this.id.split("_")[2])
    del_modal.style.display = "block";
  });
}

for (var i=0; i < del_modal_close.length; i++){
  del_modal_close[i].addEventListener("click", function() {
    del_modal.style.display = "none";
  });
}

// for (var i=0; i < modals.length; i++){
window.addEventListener("click", function(e) {
  console.log(e.target.className)
  if (e.target.className == "modal"){
    add_modal.style.display = "none";
    del_modal.style.display = "none";
  }
})
// }
