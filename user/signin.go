package user

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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

	if len(Users) == 0 {
		http.Error(w, "not user exists", http.StatusInternalServerError)
		return
	}

	for _, user := range Users {
		if user.Username == username {
			if user.Password == password {
				url := fmt.Sprintf("/home/%s", strconv.Itoa(user.Id))
				_ = user.SetLogin()
				http.Redirect(w, r, url, http.StatusSeeOther)
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
