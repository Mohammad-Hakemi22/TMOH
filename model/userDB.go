package model

type User struct {
	Id       int    `gorm:"foreignkey"`
	Username string `json:"usernagorm"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Username    string `json:"username"`
	Id          string `json:"id"`
	TokenString string `json:"token"`
}
