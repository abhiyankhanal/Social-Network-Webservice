package model

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Friends  []User `json:"friends"`
	Token    string `json:"token"`
}

type Post struct {
	ID     string `json:"id"`
	Body   string `json:"body"`
	UserID string `json:"user_id"`
}

type Token struct {
	Token string `json:"token"`
}
