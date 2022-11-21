package main

import (
	"e/db"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	var URL string
	r := gin.Default()
	dbM := db.DBManager{
		URL: os.Getenv(URL),
	}
	r.POST("/login", dbM.Login)
	r.POST("/register", dbM.Register)
	r.POST("/:user/post", dbM.CreatePost)
	r.POST("/:user/add/:friend_id", dbM.AddFriend)
	r.GET(":user/home", dbM.Home)

	r.Run()
}
