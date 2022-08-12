package routes

import (
	"github.com/Mohammad-Hakemi22/tmoh/handlers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomePage).Methods("GET")
	r.HandleFunc("/{user}", handlers.HomePageVip).Methods("GET")
	r.HandleFunc("/form", handlers.FormArticle).Methods("GET")
	r.HandleFunc("/create", handlers.CreateArticle).Methods("POST")
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
