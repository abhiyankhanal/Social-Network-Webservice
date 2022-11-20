package main

import (
	"context"
	"e/model"
	"e/utilities"
	"fmt"
	"math/rand"

	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	r := gin.Default()
	db := dbManager{
		URL: "mongodb+srv://root:1234@cluster0.ik76ncs.mongodb.net/?retryWrites=true&w=majority",
	}
	r.POST("/login", db.login)
	r.POST("/register", db.register)
	r.POST("/:user/post", db.createPost)
	r.GET("/home", home)

	r.Run()
}

func (db dbManager) login(ctx *gin.Context) {
	var user model.User
	var token model.Token

	ctx.BindJSON(&user)

	client, err := db.connect(db.URL)
	if err != nil {
		fmt.Print(err)
	}
	collection := client.Database("test").Collection("users")
	filter := bson.M{"username": user.Username, "password": user.Password}
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token.Token = utilities.CreateToken(user.Username)
	ctx.JSON(http.StatusOK, token)
}

func (db dbManager) register(ctx *gin.Context) {
	var user model.User
	ctx.BindJSON(&user)
	client, err := db.connect(db.URL)
	if err != nil {
		fmt.Print(err)
	}
	collection := client.Database("test").Collection("users")
	user.Token = utilities.CreateToken(user.Username)
	//TODO validate if it is present in db
	user.ID = fmt.Sprint(rand.Intn(9999999))
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": "User created"})
}

func (db dbManager) createPost(ctx *gin.Context) {
	var post model.Post
	ctx.BindJSON(&post)
	client, err := db.connect(db.URL)
	if err != nil {
		fmt.Print(err)
	}
	collection := client.Database("test").Collection("posts")
	//TODO validate if it is present in db
	post.ID = fmt.Sprint(rand.Int())
	post.UserID = ctx.Param("user")
	_, err = collection.InsertOne(context.Background(), post)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Error while creating post"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": "Post Created Successfully"})
}

func home(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	token = token[7:len(token)]
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "Home"})
}
