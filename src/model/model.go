package model

type User struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Friends  []string `json:"friends_id"`
	Token    string   `json:"token"`
	Posts    []Post   `json:"posts"`
}

type Post struct {
	ID                 string `json:"id"`
	Body               string `json:"body"`
	UserID             string `json:"user_id"`
	PostCountOnSameDay int    `json:"post_count"`
	ShareCount         int    `json:"share_count"`
	TimeStamp          string `json:"time_stamp"`
}

type Token struct {
	Token string `json:"token"`
}

type Links struct {
	FriendID string `json:"friend_id"`
	MyID     string `json:"my_id"`
}
