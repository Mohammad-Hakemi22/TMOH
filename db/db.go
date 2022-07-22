package db

type Article struct {
	Id    int
	Title string
	Text  string
	Date  string
	Rate  float64
	Athor *Athor
}

type Athor struct {
	Name string
	Bio  string
	Age  int
}

type User struct {
	Id       int
	Username string
	Password string
	Email    string
}
