package db

import (
	"context"
	"e/model"
	"e/utilities"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DBManager struct {
	URL string
}

type ManagerIface interface {
	Connect(uri string) (client *mongo.Client, err error)
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	CreatePost(ctx *gin.Context)
	AddFriend(ctx *gin.Context)
	Home(ctx *gin.Context)
}

func (db DBManager) Connect(uri string) (client *mongo.Client, err error) {
	client, err = mongo.NewClient(options.Client().ApplyURI(db.URL))
	if err != nil {
		fmt.Println(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
	}
	return
}

// handlers below
func (db DBManager) Login(ctx *gin.Context) {
	var user model.User
	var token model.Token

	ctx.BindJSON(&user)

	client, err := db.Connect(db.URL)
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

func (db DBManager) Register(ctx *gin.Context) {
	var user, temp model.User
	ctx.BindJSON(&user)
	client, err := db.Connect(db.URL)
	if err != nil {
		fmt.Print(err)
	}
	collection := client.Database("test").Collection("users")
	user.Token = utilities.CreateToken(user.Username)
	user.ID = fmt.Sprint(rand.Intn(9999999))
	//validation of unique userID
	err = collection.FindOne(ctx, bson.M{"id": user.ID}).Decode(&temp)
	if err != nil {
		fmt.Println(err)
	}
	if temp.ID == "" {
		_, err = collection.InsertOne(context.Background(), user)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": "User created"})
}

func (db DBManager) CreatePost(ctx *gin.Context) {
	var post model.Post
	ctx.BindJSON(&post)
	client, err := db.Connect(db.URL)
	if err != nil {
		fmt.Print(err)
	}
	collection := client.Database("test").Collection("posts")
	//TODO validate if it is present in db
	post.ID = fmt.Sprint(rand.Int())
	post.UserID = ctx.Param("user")
	post.TimeStamp = time.Now().Format("2017-09-07")
	//TODO check if the PostCountOnSameDay count is less than 10 and increase th count on every insert
	_, err = collection.InsertOne(context.Background(), post)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Error while creating post"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": "Post Created Successfully"})
}

func (db DBManager) AddFriend(ctx *gin.Context) {
	var link model.Links
	ctx.BindJSON(&link)
	client, err := db.Connect(db.URL)
	if err != nil {
		fmt.Print(err)
	}
	collection := client.Database("test").Collection("links")
	link.MyID = ctx.Param("user")
	link.FriendID = ctx.Param("friend_id")
	//todo check if the friend is present and the list is less tha 100
	_, err = collection.InsertOne(context.Background(), link)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Error while creating link"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": "Link created successfully"})
}

func (db DBManager) Home(ctx *gin.Context) {
	var posts []model.Post
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
	client, err := db.Connect(db.URL)
	if err != nil {
		fmt.Print(err)
	}
	collection := client.Database("test").Collection("posts")
	filters := bson.D{}
	cursor, err := collection.Find(context.Background(), filters)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to make DB request"})
		return
	}
	//TODO change the routes to /:user/home and use user to fetch friend list and show only their associated posts
	if err = cursor.All(ctx, &posts); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to read records"})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}
