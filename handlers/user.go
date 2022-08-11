package handlers

import (
	"fmt"
	"go/token"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/Mohammad-Hakemi22/tmoh/config"
	"github.com/Mohammad-Hakemi22/tmoh/database"
	"github.com/Mohammad-Hakemi22/tmoh/model"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func SignInForm(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./templates/user/signin.html"))
	err := tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Can't execute sign in template", http.StatusInternalServerError)
		return
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	// username := r.FormValue("username")
	// password := r.FormValue("password")

	// if len(Users) == 0 {
	// 	http.Error(w, "not user exists", http.StatusInternalServerError)
	// 	return
	// }

	// for _, user := range Users {
	// 	if user.Username == username {
	// 		if user.Password == password {
	// 			fmt.Fprintln(w, "login successfully")
	// 		} else {
	// 			fmt.Fprintln(w, "wrong password ")
	// 			return
	// 		}
	// 	} else {
	// 		fmt.Fprintln(w, "wrong username")
	// 		return
	// 	}
	// }
}

func SignUpForm(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./templates/user/signup.html"))
	err := tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Can't execute sign up template", http.StatusInternalServerError)
		return
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var dbuser model.User
	var user model.User
	connection, err := database.GetDatabase()
	if err != nil {
		log.Fatalln("something wrong in database connection", err)
	}

	defer func() {
		database.Closedatabase(connection.Conn)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}()

	rand.Seed(int64(time.Now().Nanosecond()))
	id := rand.Intn(1000000)
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")
	role := r.FormValue("role")
	user.Id = id
	user.Username = username
	user.Password = password
	user.Email = email
	user.Role = role

	connection.Conn.Where("username = ?", user.Username).First(&dbuser)
	if dbuser.Username != "" {
		fmt.Fprintln(w, "username already in use!")
		return
	}

	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("error in password hash")
	}
	connection.Conn.Create(&user)
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func GenerateJWT(username, role string) (string, error) {
	var signInKey = []byte(config.AppConfig.SECRET_KEY)
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["username"] = username
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	tokenString, err := token.SignedString(signInKey)
	if err != nil {
		log.Fatalln("Something Went Wrong in generate jwt:", err)
		return "", err
	}
	return tokenString, nil
}