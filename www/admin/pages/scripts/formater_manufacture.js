function formatmanu(){

var select = document.getElementById('fcn');
var fcn = select.options[select.selectedIndex].value;

var creator = document.getElementById('creator').value;
var contributor = document.getElementById('contributor').value;
var publisher = document.getElementById('publisher').value;
var title = document.getElementById('title').value; 
var date = document.getElementById('date').value;
var language = document.getElementById('language').value.toLowerCase();
var formats = document.getElementById('formats').value;
var subject = document.getElementById('subject').value;
var description = document.getElementById('description').value;
var relation = document.getElementById('relation').value;
var source = document.getElementById('source').value;
var type = document.getElementById('type').value.toLowerCase();
var coverage = document.getElementById('coverage').value.toUpperCase();
var rights = document.getElementById('rights').value;
var cell = document.getElementById('cell').value;
var ecm = document.getElementById('ecm').value;
var osolution = document.getElementById('osolution').value;
var csolution = document.getElementById('csolution').value;
var nozzle = document.getElementById('nozzle').value;



var today = new Date();
var dd = String(today.getDate()).padStart(2, '0');
var mm = String(today.getMonth() + 1).padStart(2, '0'); //January is 0!
var yyyy = today.getFullYear();

today = `${mm}/${dd}/${yyyy}`;

ccarguments.value = `["${fcn}","${creator}","${contributor}","${publisher}","${title}","${date}","${language}","${formats}","${subject}","${description}","${relation}","${source}","${type}","${coverage}","${rights}","${cell}","${ecm}","${osolution}","${csolution}","${nozzle}"]`;



}

