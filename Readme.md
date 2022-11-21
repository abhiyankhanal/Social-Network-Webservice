# Social Network API
- A simple webservice which can serve as an api to login/register user, add friends, create posts, & share posts
- The following project is implemented using Golang v1.9 and MongoDB.

## Getting Started ##
- Clone the project in your system
- Set environment variable foe URL as a mongo db query string
- Make sure you have docker installed to start the application
- After cloning, open a terminal, go to the root folder of the project and perform the
  following commands.

    1. docker image build -t golang_api .
    2. docker run -p 8080:8080 golang_api

  The first command will build a docker image that consists of go binary file of the project(e). 
  The second command will run the image and it can be accessed locally in port 8080, using the same 8080 port from docker. 
  
## HostUrl
    localhost:8080


# APIs #

## Register ##
- This api will register any new user in the server/db. It will create a `bearer` token using `jwt` which can be used for authorization purpose. Unique 10 million uid is generated which limits the number upto 10 million.

### Resource URL ###

[HOST_URL](#HostUrl)/register

### Method ####

```HTTP POST```

#### Sample Request Body ####

```json
{
    "username":"abhiyan",
    "password":"abhiyan"
}
```
#### Error response ####

```json
{
    "error": "Invalid credentials"
}
```


#### Request parameters ####

| Name  |Description |
| ------------- | ------------- | 
| code  | 201 if success creating, else 401 on unauthorized  |
| username  | type:string |  
| password  |type:string |  

## Login ##
- This api will login existing user in the server/db.

### Resource URL ###

[HOST_URL](#HostUrl)/login

### Method ####
```HTTP POST```

#### Sample Request Body ####

```json
{
    "username":"abhiyan",
    "password":"abhiyan"
}
```

#### Response parameters ####

| Name  |Description |
| ------------- | ------------- | 
| code  | 200 if success  |
| username  | type:string |  
| password  |type:string | 

## CreatePost ##
- This api will register will create new post and store it in post collection of mongo DB. TimeStamp will be used to track number of post created in that particular day.

### Resource URL ###

[HOST_URL](#HostUrl)/:user/post

### Method ####

```HTTP POST```

#### Sample Request Body ####

```json
{
    "body":"abhiyan",
    "user_id": ":user"
}
```

#### Request & URL parameters ####

| Name  |Description |
| ------------- | ------------- | 
| code  | 201 if success creating  |
| body  | type:string |  
| user_id  |type:string, from:URL | 

## Add Friends ##
- This api will rbe used to create a link between two user objects, the link will be stored in db. If the friends array has less that 100 User document/node then new friend will be added to that list.

### Resource URL ###

[HOST_URL](#HostUrl)//:user/add/:friend_id

### Method ####

```HTTP POST```


#### URL parameters ####

| Name  |Description |
| ------------- | ------------- | 
| :user  |type:string  |
| :friend_id | type:string |

## Get Posts ##
- This api will getch the relevantt post from post collection of mongo db. Every use has associated post in their object so that post will be shown for that specific user.

### Resource URL ###

[HOST_URL](#HostUrl)/:user/home

### Method ####

```HTTP POST```

#### Sample Request Body ####

```json
{
    "body":"abhiyan",
    "user_id": ":user"
}
```

#### Response parameters ####

| Name  |Description |
| ------------- | ------------- | 
| code  | 200  |
| message  | type:string, success |  
| data |array of post data | 



## Packages ##

There are 3 high level packages used in this project
- Mongo Driver (" go.mongodb.org/mongo-driver")
- Gin-Gonic ("github.com/gin-gonic/gin")
- jwt-GO ("github.com/dgrijalva/jwt-go")

## Some screenshots ##
<img width="1014" alt="image" src="https://user-images.githubusercontent.com/51784021/203018614-825b7dd7-5372-4892-95a7-d15307681e84.png">
<img width="912" alt="image" src="https://user-images.githubusercontent.com/51784021/203018842-7a1300f6-6383-4559-ac2c-bb450ee76a7a.png">
<img width="1019" alt="image" src="https://user-images.githubusercontent.com/51784021/203018971-7f7181a4-3d30-41bf-8787-bf34b12fb604.png">


