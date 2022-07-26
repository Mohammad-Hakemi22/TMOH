package user

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"

	"github.com/Mohammad-Hakemi22/tmoh/db"
)

var users []db.User



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
	users = append(users, db.User{Id: id, Username: username, Password: password, Email: email})
	fmt.Println(users)
}
