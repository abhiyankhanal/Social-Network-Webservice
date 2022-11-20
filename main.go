package main

import (
	"context"
	"e/model"
	"e/utilities"
	"fmt"

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
	r.GET("/home", home)

	r.Run()
}

func (db dbManager) login(c *gin.Context) {
	var user model.User
	var token model.Token

	c.BindJSON(&user)

	client, err := db.connect(db.URL)
	if err != nil {
		fmt.Print(err)
	}
	collection := client.Database("test").Collection("users")

	filter := bson.M{"username": user.Username, "password": user.Password}

	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token.Token = utilities.CreateToken(user.Username)
	c.JSON(http.StatusOK, token)
}

func (db dbManager) register(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	client, err := db.connect(db.URL)
	if err != nil {
		fmt.Print(err)
	}
	collection := client.Database("test").Collection("users")

	user.Token = utilities.CreateToken(user.Username)
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": "User created"})
}

func home(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	token = token[7:len(token)]
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Home"})
}
