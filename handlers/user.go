package handlers

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/Mohammad-Hakemi22/tmoh/config"
	"github.com/Mohammad-Hakemi22/tmoh/database"
	"github.com/Mohammad-Hakemi22/tmoh/model"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// var tokenCache []string // temp

func SignInForm(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./templates/user/signin.html"))
	err := tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Can't execute sign in template", http.StatusInternalServerError)
		return
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var authDetail model.Authentication
	var authUser model.User
	var token model.Token

	connection, err := database.GetDatabase()
	if err != nil {
		log.Fatalln("something wrong in database connection", err)
	}
	defer database.Closedatabase(connection.Conn)

	username := r.FormValue("username")
	password := r.FormValue("password")

	authDetail.Username = username
	authDetail.Password = password

	connection.Conn.Where("username = ?", authDetail.Username).First(&authUser)
	if authUser.Username == "" {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "Username or Password is incorrect")
		return
	}

	check := CheckPasswordHash(authDetail.Password, authUser.Password)
	if !check {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "Username or Password is incorrect")
		return
	}

	validToken, err := GenerateJWT(authUser.Username, authUser.Role)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "Failed to generate token")
		return
	}

	token.Username = authUser.Username
	token.Role = authUser.Role
	token.TokenString = validToken
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token.TokenString,
		Path:    "/vip",
		Expires: time.Now().Add(time.Minute * 5),
	})
	http.Redirect(w, r, "/vip", http.StatusSeeOther)
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
		w.Header().Set("Content-Type", "application/json")
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
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["username"] = SpaceFieldsJoin(username)
	claims["role"] = SpaceFieldsJoin(role)
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	tokenString, err := token.SignedString(signInKey)
	if err != nil {
		log.Fatalln("Something Went Wrong in generate jwt:", err)
		return "", err
	}
	return tokenString, nil
}

func CheckPasswordHash(password, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}

func SpaceFieldsJoin(str string) string {
	return strings.Join(strings.Fields(str), "")
}
