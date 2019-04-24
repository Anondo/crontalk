function translate()
{
  var url = "http://localhost:8008/crontalk/translate"
  var expr = document.getElementById("expression").value
  document.getElementById("result").innerHTML = ""
  if (expr == "")
  {
    expr = "* * * * *"
  }
  var xhttp = new XMLHttpRequest();
  xhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {
      document.getElementById("result").innerHTML = this.responseText;
    }else if (this.readyState == 4 && this.status == 400){
      document.getElementById("result").innerHTML = this.responseText;
    }
  };
  xhttp.open("POST", url, true);
  // xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  xhttp.send(JSON.stringify({expression:expr}));
}
