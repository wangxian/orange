package main

import "net/http"
import "log"
import "fmt"
import "os"
// import "html/template"

// func formatSize(file os.FileInfo) string {
// 	if file.IsDir() {
// 		return "-"
// 	}
// 	size := file.Size()
// 	switch {
// 	case size > 1024*1024:
// 		return fmt.Sprintf("%.1fM", float64(size)/1024/1024)
// 	case size > 1024:
// 		return fmt.Sprintf("%.1fk", float64(size)/1024)
// 	default:
// 		return strconv.Itoa(int(size))
// 	}
// 	return ""
// }

func ServeFile(w http.ResponseWriter, r *http.Request) {
	// println(r.Header.Get("accept"))
	// println(r.Header.Get("host"))
	// println(r.Header.Get("User-Agent"))
	// println(r.Header.Get("Accept-Encoding"))
	// println(r.Header.Get("Accept-Language"))
	// println(r.Header.Get("Pragma"))
	// println(r.Header.Get("Connection"))
	// w.Header().Add("name", "wangxian")
	log.Print(r.URL.Path)

	path := Config.dir + r.URL.Path
	stat, _ := os.Stat(path)
	if stat.IsDir() {
		fmt.Fprintf(w, TmplHeader +"<h1>"+ r.URL.Path +"</h1>" + `<a href="../" id="goback">..</a>`)
		http.ServeFile(w, r, path)
		fmt.Fprintf(w, TmplFooter)
	} else {
		http.ServeFile(w, r, path)
	}

	// t := template.New("DIRLIST")
	// t, _ = t.Parse(TMPL)
	// t.Execute(w, nil)
}

func LongPolling(w http.ResponseWriter, r *http.Request) {

}


func ProxySite(w http.ResponseWriter, r *http.Request) {

}