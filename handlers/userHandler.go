package handlers

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"

	db "github.com/Mohammad-Hakemi22/tmoh/database"
)

var Users []db.User

func SignInForm(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("D:/Go/TMOH/templates/user/signin.html"))
	err := tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Can't execute sign in template", http.StatusInternalServerError)
		return
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if len(Users) == 0 {
		http.Error(w, "not user exists", http.StatusInternalServerError)
		return
	}

	for _, user := range Users {
		if user.Username == username {
			if user.Password == password {
				fmt.Fprintln(w, "login successfully")
			} else {
				fmt.Fprintln(w, "wrong password ")
				return
			}
		} else {
			fmt.Fprintln(w, "wrong username")
			return
		}
	}
}

func SignUpForm(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("D:/Go/TMOH/templates/user/signup.html"))
	err := tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Can't execute sign up template", http.StatusInternalServerError)
		return
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	defer http.Redirect(w, r, "/", http.StatusSeeOther)
	rand.Seed(int64(time.Now().Nanosecond()))
	id := rand.Intn(1000000)
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")
	Users = append(Users, db.User{Id: id, Username: username, Password: password, Email: email})
	fmt.Println(Users)
}
