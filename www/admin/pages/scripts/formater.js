function format(){

var author = document.getElementById('author').value;
var project = document.getElementById('project').value; 

var today = new Date();
var dd = String(today.getDate()).padStart(2, '0');
var mm = String(today.getMonth() + 1).padStart(2, '0'); //January is 0!
var yyyy = today.getFullYear();

today = `${mm}/${dd}/${yyyy}`;

ccarguments.value = `["${project}","${author}","${today}","Success"]`;



}

