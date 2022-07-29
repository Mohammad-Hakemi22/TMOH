package db

type Article struct {
	Id    int
	Title string
	Text  string
	Date  string
	Rate  float64
	Athor *Athor
	VIP   bool
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
	isLogin  bool
}

func (u *User) SetLogin() bool {
	u.isLogin = true
	return u.isLogin
}

func (u *User) SetLogout() bool {
	u.isLogin = false
	return u.isLogin
}
