var url = "http://localhost:8008/crontalk/translate";

$(document).ready(function(){
  $("#expression").unbind();
  $("#expression").bind("keyup",function(){
    var expr = $("#expression").val();
    $.post(url , JSON.stringify({expression: expr}) , function(response , status){
      $("#result").val('');
      $("#result").val(response);
    }).fail(function(jqXHR , textStatus , errorThrown){
      $("#result").val('');
      $("#result").val(jqXHR.responseText);
    });
  });
  return;
})
