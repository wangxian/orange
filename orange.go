// Copyright 2013 The Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
//
// Run: go run defines.go orange.go --port 4000 --dir ../ --portproxy 80 --ignores .go -precmd "ls -lhr"
// INSTALL: go install
//
// Project URL: https://github.com/wangxian/orange
package main

import (
	"os"
	"log"
	"flag"
	"strconv"
	"net/http"
	"github.com/toqueteos/webbrowser"
)

// Hander http request
func dispatch(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	} else if r.URL.Path == "/_longpolling.js" {
		LongPolling(w, r)
	} else if Config.portproxy == 0 {
		ServeFile(w, r)
	} else {
		ProxySite(w, r)
	}
}

// func init() {
//     runtime.GOMAXPROCS(runtime.NumCPU())
// }

func main() {
	flag.IntVar(&Config.port, "port", 4000, "Static server port, The port must>1024, default 4000")
	flag.IntVar(&Config.portproxy, "portproxy", 0, "Proxy http://localhost:{{port}}/ when file saved refresh browser, set 0 not proxy")
	flag.StringVar(&Config.dir, "dir", "./", "Watch dir which change will refresh the browser, default current dir")
	flag.StringVar(&Config.ignores, "ignores", "", "Not watch files, split width `,` Not regexp eg: `.go,.git/`, default no ignores")
	flag.StringVar(&Config.precmd, "precmd", "", "Before refresh browser, execute precmd command. eg: `ls {0}`, {0} is the changed file")
	flag.Parse()

	os.Chdir(Config.dir)
	Config.dir, _ = os.Getwd()

	println("-------------------------------------------------------------------------")
	println("port:		",	Config.port)
	println("portproxy:	", Config.portproxy)
	println("dir:		", Config.dir)
	println("ignores:	", Config.ignores)
	println("precmd:	", Config.precmd)

	port := ":" + strconv.Itoa(Config.port)
	println("HostAt: 	 http://0.0.0.0"+ port +"/")
	println("-------------------------------------------------------------------------")

	http.HandleFunc("/", dispatch)

	// Stare Watcher
	Watcher(Config.dir)

	// open browser
	webbrowser.Open("http://localhost"+ port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}