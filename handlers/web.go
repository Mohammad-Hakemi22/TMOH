package handlers

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Mohammad-Hakemi22/tmoh/config"
	db "github.com/Mohammad-Hakemi22/tmoh/model"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

var articles = []db.Article{}

func HomePage(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./templates/home.html"))
	err := tpl.Execute(w, articles)
	if err != nil {
		http.Error(w, "Can't execute template", http.StatusInternalServerError)
		return
	}
}

func HomePageVip(w http.ResponseWriter, r *http.Request) {
	tok, _ := r.Cookie("token")
	tokenString := tok.Value
	if tokenString == "" {
		fmt.Fprintln(w, "No token found!")
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return []byte(config.AppConfig.SECRET_KEY), nil
	})
	if err != nil {
		fmt.Println(w, "Your Token has been expired", err.Error())
		http.Redirect(w, r, "/user/signinform", http.StatusSeeOther)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["role"] == "admin" {
			tpl := template.Must(template.ParseFiles("./templates/home.html"))
			err := tpl.Execute(w, articles)
			if err != nil {
				http.Error(w, "Can't execute template", http.StatusInternalServerError)
				return
			}
			return

		} else if claims["role"] == "user" {

			r.Header.Set("Role", "user")
			return
		}
	}

}

func FormArticle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./templates/create.html"))
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
	tpl := template.Must(template.ParseFiles("./templates/update.html"))
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
	tpl := template.Must(template.ParseFiles("./templates/delete.html"))
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
