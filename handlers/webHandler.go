package handlers

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	db "github.com/Mohammad-Hakemi22/tmoh/database"
	"github.com/gorilla/mux"
	
)

var articles = []db.Article{}

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
	vip := false
	rand.Seed(int64(time.Now().Nanosecond()))
	id := rand.Intn(1000000)
	title := r.FormValue("title")
	text := r.FormValue("text")
	date := time.Now().Format("01-02-2006 Monday")
	rate, _ := strconv.ParseFloat(r.FormValue("rate"), 64)
	name := r.FormValue("AuthorName")
	bio := r.FormValue("AuthorBio")
	age, _ := strconv.Atoi(r.FormValue("AuthorAge"))
	if r.FormValue("vip") == "0" {
		vip = false
	} else {
		vip = true
	}
	articles = append(articles, db.Article{Id: id, Title: title, Text: text, Date: date, Rate: rate, VIP: vip, Athor: &db.Athor{Name: name, Bio: bio, Age: age}})
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
	vip := false
	title := r.FormValue("title")
	text := r.FormValue("text")
	date := time.Now().Format("01-02-2006 Monday")
	rate, _ := strconv.ParseFloat(r.FormValue("rate"), 64)
	name := r.FormValue("AuthorName")
	bio := r.FormValue("AuthorBio")
	age, _ := strconv.Atoi(r.FormValue("AuthorAge"))
	if r.FormValue("vip") == "0" {
		vip = false
	} else {
		vip = true
	}
	articles = append(articles, db.Article{Title: title, Text: text, Date: date, Rate: rate, VIP: vip, Athor: &db.Athor{Name: name, Bio: bio, Age: age}})
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
