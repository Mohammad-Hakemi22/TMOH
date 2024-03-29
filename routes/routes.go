package routes

import (
	"github.com/Mohammad-Hakemi22/tmoh/handlers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.IndexPage).Methods("GET")
	r.HandleFunc("/home", handlers.HomePage).Methods("GET")
	r.HandleFunc("/form/{id}", handlers.FormArticle).Methods("GET")
	r.HandleFunc("/create/{id}", handlers.CreateArticle).Methods("POST")
	r.HandleFunc("/update/{id}", handlers.UpdateArticle).Methods("GET")
	r.HandleFunc("/deleteform", handlers.DeleteArticleForm).Methods("GET")
	r.HandleFunc("/update", handlers.Update).Methods("POST")
	r.HandleFunc("/delete", handlers.DeleteArticle).Methods("POST")

	// User
	r.HandleFunc("/user/signupform", handlers.SignUpForm)
	r.HandleFunc("/user/signup", handlers.SignUp)
	r.HandleFunc("/user/signinform", handlers.SignInForm)
	r.HandleFunc("/user/signin", handlers.SignIn)
	return r
}
