package main

import (
 "fmt"
 "net/http"       
 "html/template" // Added in Lesson 2
)

type SearchPage struct {
 Query string
 Results []Result
}

type Result struct {
 Title string
 Author string
 Year int
 ID string
}

func main() {
 templates := template.Must(template.ParseFiles("templates/index.html"))

 http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
	 results := []Result{
	   Result{"1984", "Orwell", 1950, "123"},
	   Result{"Persuasion", "Austen", 1817, "234"},
	   Result{"Holes", "Sachar", 2000, "345"},
	 }
	 p := SearchPage{Query: r.FormValue("search"), Results: results}

	 if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	 }
	})

 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
   if e := templates.ExecuteTemplate(w, "index.html", nil); e != nil {
     http.Error(w, e.Error(), http.StatusInternalServerError)
   }
 })
 fmt.Println(http.ListenAndServe(":4000", nil))
}