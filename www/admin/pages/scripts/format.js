function format(){

//var args = document.getElementById('ccarguments').value;
//var author = document.getElementById('author');
var args = document.getElementById('ccarguments'); 
var author = document.getElementById('author').value;
var project = document.getElementById('project').value; 
var user = document.getElementById('username').value; 

var today = new Date();
var dd = String(today.getDate()).padStart(2, '0');
var mm = String(today.getMonth() + 1).padStart(2, '0'); //January is 0!
var yyyy = today.getFullYear();

today = `${mm}/${dd}/${yyyy}`;

            //args.value = `["CAD1","Pavel","06/06/2020","Test1"]`
            args.value =  `["${project}","${user}","${today}","Success"]`;

//return args;
//document.getElementById("ccarguments") = "[\"CAD1\",\"Pavel\",\"06/06/19\",\"Test\"]";

}

format()