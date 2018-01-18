package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl_iter.html")
	daysOfWeek := []string{"月", "火", "水", "木", "金", "土", "日"}
	t.Execute(w, daysOfWeek)
}

func noProcess(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl_iter.html")
	nothing := []string{}
	t.Execute(w, nothing)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/noprocess", noProcess)
	server.ListenAndServe()
}
