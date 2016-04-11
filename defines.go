package main

import (
	"bufio"
	"net"
)

var Config struct {
	http     string
	proxy    string
	precmd   string
	rootdir  string
	watchdir []string
	openURL  string
	ignores  string
}

const VERSION = "orange/3.5"

// Store client buffer handle
type Client struct {
	bufrw *bufio.ReadWriter
	conn  net.Conn
}

var Clients = make([]Client, 0)

var TmplHeader = `
<!DOCTYPE html>
<html>
<head>
<meta charset='UTF-8'>
<title></title>
<style type="text/css">
ul a:focus, ul a:hover, a#goback:hover { outline: none; background: rgba(255,255,255,0.65); }
ul a.highlight {-webkit-transition: background .4s ease-in-out; background: #ffff4f; border-color: #E9DC51; }
#search {display: block; position: fixed; top: 20px; right: 20px; width: 90px; transition: width ease 0.2s, opacity ease 0.4s; -moz-transition: width ease 0.2s, opacity ease 0.4s; -webkit-border-radius: 16px; -moz-border-radius: 16px; -webkit-box-shadow: inset 0px 0px 3px rgba(0, 0, 0, 0.25), inset 0px 1px 3px rgba(0, 0, 0, 0.7), 0px 1px 0px rgba(255, 255, 255, 0.03); -moz-box-shadow: inset 0px 0px 3px rgba(0, 0, 0, 0.25), inset 0px 1px 3px rgba(0, 0, 0, 0.7), 0px 1px 0px rgba(255, 255, 255, 0.03); -webkit-font-smoothing: antialiased; text-align: left; font: 13px "Helvetica Neue", Arial, sans-serif; padding: 4px 10px; border: none; }
#search:focus { width: 180px; opacity: 1.0; }
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
  var links = $('wrapper').all('ul a');

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
    <p><em>X-Powered-By: ` + VERSION + `</em></p>
    </div>
  </body>
</html>
`
var Tmplpolljs = `
<script style="text/javascript">
(function(){
  var oldOnload = window.onload ? window.onload : function(){};
  window.onload = function(){
    oldOnload() || setTimeout(function(){
      var js = document.createElement('script');
      js.src = "/_longpolling.js";
      document.getElementsByTagName("head")[0].appendChild(js);
      if(window.console && console.log) {
        console.log("["+ (new Date()).toLocaleString() +"]orange watcher js is working.");
      }
    }, 800);
  }
}).call(this);
</script>
`
