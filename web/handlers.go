package web

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Mohammad-Hakemi22/tmoh/db"
	"github.com/Mohammad-Hakemi22/tmoh/user"
	"github.com/gorilla/mux"
)

var articles = []db.Article{}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomePage).Methods("GET")
	r.HandleFunc("/form", FormArticle).Methods("GET")
	r.HandleFunc("/create", CreateArticle).Methods("POST")
	r.HandleFunc("/update/{id}", UpdateArticle).Methods("GET")
	r.HandleFunc("/deleteform", DeleteArticleForm).Methods("GET")
	r.HandleFunc("/update", Update).Methods("POST")
	r.HandleFunc("/delete", DeleteArticle).Methods("POST")
	r.HandleFunc("/user/signupform", user.SignUpForm)
	r.HandleFunc("/user/signup", user.SignUp)
	return r
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("D:/Go/TMOH/templates/home.html"))
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
	defer http.Redirect(w, r, "/", http.StatusSeeOther)
	rand.Seed(int64(time.Now().Nanosecond()))
	id := rand.Intn(1000000)
	title := r.FormValue("title")
	text := r.FormValue("text")
	date := time.Now().Format("01-02-2006 Monday")
	rate, _ := strconv.ParseFloat(r.FormValue("rate"), 64)
	name := r.FormValue("AuthorName")
	bio := r.FormValue("AuthorBio")
	age, _ := strconv.Atoi(r.FormValue("AuthorAge"))
	articles = append(articles, db.Article{Id: id, Title: title, Text: text, Date: date, Rate: rate, Athor: &db.Athor{Name: name, Bio: bio, Age: age}})
	fmt.Println(articles)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("D:/Go/TMOH/templates/update.html"))
	params := mux.Vars(r)
	for idx, a := range articles {
		if id, _ := strconv.Atoi(params["id"]); id == a.Id {
			tpl.Execute(w, a)
			articles = append(articles[:idx], articles[idx+1:]...)
		}
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer http.Redirect(w, r, "/", http.StatusSeeOther)
	title := r.FormValue("title")
	text := r.FormValue("text")
	date := time.Now().Format("01-02-2006 Monday")
	rate, _ := strconv.ParseFloat(r.FormValue("rate"), 64)
	name := r.FormValue("AuthorName")
	bio := r.FormValue("AuthorBio")
	age, _ := strconv.Atoi(r.FormValue("AuthorAge"))
	articles = append(articles, db.Article{Title: title, Text: text, Date: date, Rate: rate, Athor: &db.Athor{Name: name, Bio: bio, Age: age}})
}

func DeleteArticleForm(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("D:/Go/TMOH/templates/delete.html"))
	err := tpl.Execute(w, "")
	if err != nil {
		http.Error(w, "Can't execute template", http.StatusInternalServerError)
		return
	}
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	defer http.Redirect(w, r, "/", http.StatusSeeOther)
	id, _ := strconv.Atoi(r.FormValue("id")) 
	for idx, a := range articles {
		if a.Id == id {
			articles = append(articles[:idx], articles[idx+1:]...)
			return
		}
	}
}