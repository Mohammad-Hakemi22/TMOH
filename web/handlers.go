package web

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/Mohammad-Hakemi22/tmoh/db"
	"github.com/gorilla/mux"
)

var articles = []db.Article{}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomePage).Methods("GET")
	r.HandleFunc("/form", FormArticle).Methods("GET")
	r.HandleFunc("/create", CreateArticle).Methods("POST")

	return r
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	// wd, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	tpl := template.Must(template.ParseFiles("D:/Go/TMOH/templates/home.html"))
	// articles = append(articles, db.Article{Title: "hi", Text: "hello world", Date: time.Now(), Rate: 2.6, Athor: &db.Athor{Name: "mmd", Bio: "hi there", Age: 25}})
	// articles = append(articles, db.Article{Title: "khabar", Text: "salam !", Date: time.Now(), Rate: 3.7, Athor: &db.Athor{Name: "ali", Bio: "salam", Age: 35}})
	err := tpl.Execute(w, articles)
	if err != nil {
		http.Error(w, "Can't execute template", http.StatusInternalServerError)
		return
	}

}

func FormArticle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("D:/Go/TMOH/templates/create.html"))
	err := tpl.Execute(w, "")
	if err != nil {
		http.Error(w, "Can't execute template", http.StatusInternalServerError)
		return
	}
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	text := r.FormValue("text")
	date := time.Now()
	rate, _ := strconv.ParseFloat(r.FormValue("rate"), 32)
	name := r.FormValue("AuthorName")
	bio := r.FormValue("AuthorBio")
	age, _ := strconv.Atoi(r.FormValue("AuthorAge"))
	articles = append(articles, db.Article{Title: title, Text: text, Date: date, Rate: rate, Athor: &db.Athor{Name: name, Bio: bio, Age: age}})
	http.Redirect(w, r, "/", http.StatusOK)
}
