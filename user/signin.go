package user

import (
	"fmt"
	"html/template"
	"net/http"
)

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

	if len(users) == 0 {
		http.Error(w, "not user exists", http.StatusInternalServerError)
		return
	}

	for _, user := range users {
		if user.Username == username {
			if user.Password == password {
				fmt.Fprintln(w, "welcome ")
			} else {
				fmt.Fprintln(w, "wrong password ")
			}
		} else {
			fmt.Fprintln(w, "wrong username")
		}
	}
}
