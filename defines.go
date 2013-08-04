package main

import (
	"net"
	"bufio"
)

var Config struct {
	port			int
	portproxy	int
	precmd		string
	rootdir		string
	watchdir	string
	ignores		string
	// pipchan		chan bool
}

const VERSION = "orange/2.4"

// Store client buffer handle
type Client struct {
	bufrw  *bufio.ReadWriter
	conn net.Conn
}

var Clients = make([]Client, 0)

var TmplHeader = `
<!DOCTYPE html>
<html>
<head>
<meta charset='utf-8'>
<title></title>
<style type="text/css">
body {margin: 0; padding: 80px 100px; font: 13px "Helvetica Neue", "Lucida Grande", "Arial"; background: #ECE9E9 -webkit-gradient(linear, 0% 0%, 0% 100%, from(#fff), to(#ECE9E9)); background: #ECE9E9 -moz-linear-gradient(top, #fff, #ECE9E9); background-repeat: no-repeat; color: #555; -webkit-font-smoothing: antialiased; }
h1, h2, h3 {margin: 0; font-size: 22px; color: #343434; }
h1 em, h2 em {padding: 0 5px; font-weight: normal;}
h1 {font-size: 60px;text-indent: 8px;}
h2 {margin-top: 10px;}
h3 {margin: 5px 0 10px 0; padding-bottom: 5px; border-bottom: 1px solid #eee; font-size: 18px; }
ul {margin: 0; padding: 0; }
ul li {margin: 5px 0; padding: 3px 8px; list-style: none; }
ul li:hover {cursor: pointer; color: #2e2e2e; }

ul li:first-child .path {padding-left: 0; }
p { line-height: 1.5; }
a { color: #555; text-decoration: none; }
a:hover {color: #303030; }

#stacktrace { margin-top: 15px; }
.directory h1 {margin-bottom: 15px; font-size: 18px; }
pre { white-space:normal;clear:both;margin:0;padding:0; }
pre a { padding: 0; width: 24%; float:left; }
pre a, a#goback { display:inline-block;margin:1px;height:25px; line-height: 25px; text-indent: 8px; float: left; border: 1px solid transparent; -webkit-border-radius: 5px; -moz-border-radius: 5px; border-radius: 5px; overflow: hidden; text-overflow: ellipsis; }
pre a:focus, pre a:hover, a#goback:hover { outline: none; background: rgba(255,255,255,0.65); border: 1px solid #ececec; }
pre a.highlight {-webkit-transition: background .4s ease-in-out; background: #ffff4f; border-color: #E9DC51; }
a#goback { width:100%; }
#search {display: block; position: fixed; top: 20px; right: 20px; width: 90px; -webkit-transition: width ease 0.2s, opacity ease 0.4s; -moz-transition: width ease 0.2s, opacity ease 0.4s; -webkit-border-radius: 32px; -moz-border-radius: 32px; -webkit-box-shadow: inset 0px 0px 3px rgba(0, 0, 0, 0.25), inset 0px 1px 3px rgba(0, 0, 0, 0.7), 0px 1px 0px rgba(255, 255, 255, 0.03); -moz-box-shadow: inset 0px 0px 3px rgba(0, 0, 0, 0.25), inset 0px 1px 3px rgba(0, 0, 0, 0.7), 0px 1px 0px rgba(255, 255, 255, 0.03); -webkit-font-smoothing: antialiased; text-align: left; font: 13px "Helvetica Neue", Arial, sans-serif; padding: 4px 10px; border: none; background: transparent; margin-bottom: 0; outline: none; opacity: 0.7; color: #888; }
#search:focus { width: 120px; opacity: 1.0; }
</style>
<script type="text/javascript">
function $(id){
  var el = 'string' == typeof id ? document.getElementById(id) : id;

  el.on = function(event, fn){
    if ('content loaded' == event) { event = window.attachEvent ? "load" : "DOMContentLoaded"; }
    el.addEventListener ? el.addEventListener(event, fn, false) : el.attachEvent("on" + event, fn);
  };

  el.all = function(selector){ return $(el.querySelectorAll(selector)); };
  el.each = function(fn){ for(var i = 0, len = el.length; i < len; ++i) {fn($(el[i]), i); } };
  el.getClasses = function(){
      return this.getAttribute('class') ? this.getAttribute('class').split(/\s+/) : [];
  };

  el.addClass = function(name){
    var classes = this.getAttribute('class');
    el.setAttribute('class', classes ? classes + ' ' + name : name);
  };

  el.removeClass = function(name){
    var classes = this.getClasses().filter(function(curr){
      return curr != name;
    });
    this.setAttribute('class', classes);
  };
  return el;
}

function search() {
  var str = $('search').value;
  var links = $('wrapper').all('pre a');

  links.each(function(link){
    var text = link.textContent;
    if ('..' == text) return;
    if (str.length && ~text.indexOf(str)) {
      link.addClass('highlight');
    } else {
      link.removeClass('highlight');
    }
  });
}

$(window).on('content loaded', function(){
  $('search').on('keyup', search);
});
</script>
</head>
  <body class="directory">
    <input id="search" type="text" placeholder="Search" autocomplete="off" />
    <div id="wrapper">
`
var TmplFooter = `
    </div>
  </body>
</html>
`
var Tmplpolljs = `
<script style="text/javascript">
window.onload = function(){
  setTimeout(function(){
    var js = document.createElement('script');
    js.src = "/_longpolling.js";
    document.getElementsByTagName("head")[0].appendChild(js);
    if(window.console && console.log) {
      console.log("["+ (new Date()).toLocaleString() +"]orange watcher js is working.");
    }
  }, 800);
}
</script>
`
