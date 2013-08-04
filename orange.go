// Copyright 2013 The Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
//
// Test Run: go run defines.go orange.go --port 4000 --dir ../ --portproxy 80 --ignores .go -precmd "ls -lhr"
// INSTALL: go install
//
// @Author: wangxian
// @Created: 2013-07-27
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
	flag.StringVar(&Config.rootdir, "rootdir", "./", "Server root dir, default current dir")
	flag.StringVar(&Config.watchdir, "watchdir", "./", "Watch dir which change will refresh the browser, default current dir")
	flag.StringVar(&Config.ignores, "ignores", "", "Not watch files, split width `,` Not regexp eg: `.go,.git/`, default no ignores")
	flag.StringVar(&Config.precmd, "precmd", "", "Before refresh browser, execute precmd command. eg: `ls {0}`, {0} is the changed file")
	flag.Parse()

	// Current dir
	curdir, _ := os.Getwd()

	os.Chdir(Config.rootdir)
	Config.rootdir, _ = os.Getwd()

	os.Chdir(curdir)
	os.Chdir(Config.watchdir)
	Config.watchdir, _ = os.Getwd()

	println("-------------------------------------------------------------------------")
	println("port:		",	Config.port)
	println("portproxy:	", Config.portproxy)
	println("dir:		", Config.rootdir)
	println("watchdir:	", Config.watchdir)
	println("ignores:	", Config.ignores)
	println("precmd:	", Config.precmd)

	port := ":" + strconv.Itoa(Config.port)
	println("HostAt: 	 http://0.0.0.0"+ port +"/")
	println("-------------------------------------------------------------------------")
	// println(Config.ignores != "" && Config.ignores != ".")
	// os.Exit(1)
	http.HandleFunc("/", dispatch)

	// Start Watcher
	Watcher(Config.watchdir)

	// open browser
	webbrowser.Open("http://localhost"+ port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}