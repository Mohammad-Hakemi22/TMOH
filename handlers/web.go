package handlers

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/Mohammad-Hakemi22/tmoh/config"
	"github.com/Mohammad-Hakemi22/tmoh/database"
	"github.com/Mohammad-Hakemi22/tmoh/model"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type Data map[string]interface{}

var articles = []model.Article{}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./templates/index.html"))
	err := tpl.Execute(w, articles)
	if err != nil {
		http.Error(w, "Can't execute template", http.StatusInternalServerError)
		return
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	var authorId string

	connection, err := database.GetDatabase()
	if err != nil {
		log.Fatalln("something wrong in database connection", err)
	}
	defer database.Closedatabase(connection.Conn)

	tok, _ := r.Cookie("token")
	if tok == nil {
		http.Redirect(w, r, "/user/signinform", http.StatusSeeOther)
		return
	}
	tokenString := tok.Value
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
		var articles []model.Article
		if claims["role"] == "vipUser" || claims["role"] == "author" {
			authorId = fmt.Sprintf("%v", claims["id"])
			results := connection.Conn.Find(&articles)
			fmt.Println(results.RowsAffected)
			tpl := template.Must(template.ParseFiles("./templates/home.html"))
			data := Data{
				"Url":      authorId,
				"Articles": articles,
			}
			err := tpl.Execute(w, data)
			if err != nil {
				log.Println(err)
				http.Error(w, "Can't execute template", http.StatusInternalServerError)
				return
			}
			return

		} else if claims["role"] == "user" {
			// showing just non vip articles
			r.Header.Set("Role", "user")
			return
		}
	}

}

func FormArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	data := Data{
		"Url": params["id"],
	}
	tpl := template.Must(template.ParseFiles("./templates/create.html"))
	err := tpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Can't execute template", http.StatusInternalServerError)
		return
	}
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var article model.Article
	connection, err := database.GetDatabase()
	if err != nil {
		log.Fatalln("something wrong in database connection", err)
	}
	defer func() {
		database.Closedatabase(connection.Conn)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}()

	vip := false
	rand.Seed(int64(time.Now().Nanosecond()))
	id := rand.Intn(1000000)
	title := r.FormValue("title")
	text := r.FormValue("text")
	date := time.Now().Format("01-02-2006 Monday")
	rate, _ := strconv.ParseFloat(r.FormValue("rate"), 64)
	if r.FormValue("vip") == "0" {
		vip = false
	} else {
		vip = true
	}
	article.Id = id
	article.Title = title
	article.Text = text
	article.Date = date
	article.Rate = rate
	article.Vip = vip
	article.AuthorID, _ = strconv.Atoi(params["id"])
	connection.Conn.Create(&article)

}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var article model.Article
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	connection, err := database.GetDatabase()
	if err != nil {
		log.Fatalln("something wrong in database connection", err)
	}
	defer database.Closedatabase(connection.Conn)
	tpl := template.Must(template.ParseFiles("./templates/update.html"))
	connection.Conn.Find(&article, id)
	tpl.Execute(w, article)
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer http.Redirect(w, r, "/", http.StatusSeeOther)
	vip := false
	title := r.FormValue("title")
	text := r.FormValue("text")
	date := time.Now().Format("01-02-2006 Monday")
	rate, _ := strconv.ParseFloat(r.FormValue("rate"), 64)
	if r.FormValue("vip") == "0" {
		vip = false
	} else {
		vip = true
	}
	articles = append(articles, model.Article{Title: title, Text: text, Date: date, Rate: rate, Vip: vip})
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
	connection, err := database.GetDatabase()
	if err != nil {
		log.Fatalln("something wrong in database connection", err)
	}
	defer func() {
		database.Closedatabase(connection.Conn)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}()
	id, _ := strconv.Atoi(r.FormValue("id"))
	connection.Conn.Delete(&model.Article{}, id)
}
