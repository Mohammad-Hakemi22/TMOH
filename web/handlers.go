package web

import (
	"html/template"
	"net/http"
	"time"

	"github.com/Mohammad-Hakemi22/tmoh/db"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomePage).Methods("GET")

	return r
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("D:/Go/TMOH/templates/home.html"))
	articles := []db.Article{}
	articles = append(articles, db.Article{Title: "hi", Text: "hello world", Date: time.Now(), Rate: 2.6, Athor: &db.Athor{Name: "mmd", Bio: "hi there", Age: 25}})
	articles = append(articles, db.Article{Title: "khabar", Text: "salam !", Date: time.Now(), Rate: 3.7, Athor: &db.Athor{Name: "ali", Bio: "salam", Age: 35}})
	tpl.Execute(w, articles)
	// if err != nil {
	// 	http.Error(w, "Can't execute template", http.StatusInternalServerError)
	// 	return
	// }

}
