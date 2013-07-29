package main

import (
	"net/http"
	"strings"
	"log"
	"fmt"
	"os"
	"io"
	// "time"
	// "html/template"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL.Path)

	// println(r.Header.Get("accept"))
	// println(r.Header.Get("host"))
	// println(r.Header.Get("User-Agent"))
	// println(r.Header.Get("Accept-Encoding"))
	// println(r.Header.Get("Accept-Language"))
	// println(r.Header.Get("Pragma"))
	// println(r.Header.Get("Connection"))
	w.Header().Add("Server", "orange/1.1")

	path := Config.dir + r.URL.Path
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Fprintf(w, "Error 404:\r\n"+ path +" is not exist.")
	} else if stat.IsDir() {
		fmt.Fprintf(w, TmplHeader +"<h1>"+ r.URL.Path +"</h1>" + `<a href="../" id="goback">..</a>`)
		http.ServeFile(w, r, path)
		fmt.Fprintf(w, TmplFooter)
	} else {

		if strings.Contains(r.URL.Path, ".html") {
			f, _ := os.Open(path)
			defer f.Close()

			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(200)
			io.Copy(w, f)
			w.Write([]byte(Tmplpolljs))

			return
		}

		http.ServeFile(w, r, path)
	}

	// t := template.New("DIRLIST")
	// t, _ = t.Parse(TMPL)
	// t.Execute(w, nil)
}


// Handler long polling request
func LongPolling(w http.ResponseWriter, r *http.Request) {
	// w.Header().Add("Content-Type", "text/javascript")
	// w.Header().Add("Cache-Control", "no-cache")
	// body := "console.log(123);"
	// w.Write([]byte(body))

	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
		return
	}
	conn, bufrw, err := hj.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Clients = append(Clients, Client{bufrw, conn})
	log.Println("Now Clients count:", len(Clients))

	// changed := <- Config.pipchan
	// log.Println("Hijack: ", changed, ", reload browser page.")

	// Don't forget to close the connection:
	// defer conn.Close()
	// w.Header().Set("Content-Type", "text/html")
	// bufrw.WriteString("console.log(new Date())")
	// bufrw.Flush()
}

func RefreshBrowser() {
	for _, c := range Clients {
		defer c.conn.Close()
		body := "HTTP/1.1 200 OK\r\n"
		body += "Cache-Control: no-cache\r\nContent-Type: text/javascript\r\n\r\n"
		body += "window.location.reload();"
		c.bufrw.Write([]byte(body))
		c.bufrw.Flush()
	}
	Clients = make([]Client, 0)
}


func ProxySite(w http.ResponseWriter, r *http.Request) {
}