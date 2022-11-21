package main

import (
	"e/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	dbM := db.DBManager{
		URL: "mongodb+srv://root:1234@cluster0.ik76ncs.mongodb.net/?retryWrites=true&w=majority",
	}
	r.POST("/login", dbM.Login)
	r.POST("/register", dbM.Register)
	r.POST("/:user/post", dbM.CreatePost)
	r.POST("/:user/add/:friend_id", dbM.AddFriend)
	r.GET("/home", dbM.Home)

	r.Run()
}
