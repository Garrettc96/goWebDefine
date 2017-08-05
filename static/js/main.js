$(function(){
  $("#dropDown")
})
function submitSearch(){
  console.log("called?  ")
  var text = $("#searchBar").val()
  console.log(text)
  $.ajax ({
    url: "/search",
    method: "POST",
    data: {text:text},
    success: function(rawData){
      console.log(rawData)
      var parsed = JSON.parse(rawData)
      if (!parsed) return;
      $("#dropDown ul").empty()
      var defs = parsed.Definition
      if (defs){
        defs.forEach(function(input){
          var li = $("<li>"+input.substring(1)+"</li>")
          li
          $("#dropDown ul").append(li)
        })
      }
      else{
        $("#dropDown ul").append($("<li style = 'color:red'>Word not found</li>"))
      }
    }
  });
}
$(document).on('click','body:not(.container)',function(){
  $("#dropDown ul").empty()
  console.log("test")
})
