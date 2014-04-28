// Copyright 2013 The Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
//
// INSTALL: go install
//
// @Author: wangxian
// @Created: 2013-07-27
//
// Project URL: https://github.com/wangxian/orange
package main

import (
	"flag"
	"github.com/toqueteos/webbrowser"
	"log"
	"net/http"
	"os"
	"strings"
)

// Hander http request
func dispatch(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	} else if r.URL.Path == "/_longpolling.js" {
		LongPolling(w, r)
	} else if Config.proxy != "" {
		ProxySite(w, r)
	} else {
		ServeFile(w, r)
	}
}

// func init() {
//     runtime.GOMAXPROCS(runtime.NumCPU())
// }

// watch Dir string from cli
var watchdirStr string

func main() {
	flag.StringVar(&Config.http, "http", ":4000", "Static server port, The port must>1024, default :4000")
	flag.StringVar(&Config.proxy, "proxy", "", "Proxy webserver when file saved refresh browser, like :80")
	flag.StringVar(&Config.rootdir, "rootdir", "./", "Server root dir, default current dir")
	flag.StringVar(&watchdirStr, "watchdir", "", "Watch dir which change will refresh the browser, default watch Nothing")
	flag.StringVar(&Config.openURL, "openurl", "/", "Open URL in browser. eg: /dir/apps/, default '/'")
	flag.StringVar(&Config.ignores, "ignores", "", "Not watch files, split width `,` Not regexp eg: `.go,.git/`, default no ignores")
	flag.StringVar(&Config.precmd, "precmd", "", "Before refresh browser, execute precmd command. eg: `ls {0}`, {0} is the changed file")

	flag.StringVar(&Config.rootdir, "r", "./", "Alias -rootdir")
	flag.StringVar(&watchdirStr, "w", "", "Alias -watchdir")
	flag.StringVar(&Config.openURL, "o", "/", "Alias -openurl")
	flag.StringVar(&Config.ignores, "i", "", "Alias -ignores")

	flag.Parse()

	// Current dir
	curdir, _ := os.Getwd()

	os.Chdir(Config.rootdir)
	Config.rootdir, _ = os.Getwd()

	if watchdirStr != "" {
		tList := strings.Split(watchdirStr, ",")
		for _, v := range tList {
			os.Chdir(curdir)
			os.Chdir(v)

			tPath, _ := os.Getwd()
			Config.watchdir = append(Config.watchdir, tPath)
		}
	}

	if !strings.Contains(Config.http, ":") {
		log.Fatal("Config.http must Contains `:`")
	}

	httpList := strings.Split(Config.http, ":")
	openURL := ""
	if httpList[0] == "" {
		openURL = "http://localhost:" + httpList[1]
	} else {
		openURL = "http://" + Config.http
	}

	if Config.proxy != "" {
		if !strings.Contains(Config.proxy, ":") {
			log.Fatal("Config.proxy must Contains `:`")
		}

		proxyList := strings.Split(Config.proxy, ":")
		if proxyList[0] == "" {
			Config.proxy = "http://127.0.0.1:" + proxyList[1]
		} else {
			Config.proxy = "http://" + Config.proxy
		}
	}

	println("-------------------------------------------------------------------------")
	println("http	:", Config.http)
	println("proxy	:", Config.proxy)
	println("rootdir	:", Config.rootdir)
	println("watchdir:", watchdirStr)
	println("ignores	:", Config.ignores)
	println("precmd	:", Config.precmd)

	println("HostAt	:", openURL)
	println("-------------------------------------------------------------------------")

	http.HandleFunc("/", dispatch)

	// Start Watcher
	if len(Config.watchdir) > 0 {
		for _, v := range Config.watchdir {
			Watcher(v)
		}
	}

	// Open browser
	webbrowser.Open(openURL + Config.openURL)

	err := http.ListenAndServe(Config.http, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
