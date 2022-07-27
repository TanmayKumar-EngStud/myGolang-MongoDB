# Instagram Backend API
Tanmay Kumar
<img src="https://www.vhv.rs/dpng/d/211-2119308_transparent-mongodb-png-mongodb-update-one-collection-golang.png" align="right"
     alt="GoLang img didn't loaded :(" width="120" >

Aim is to Design and Develop an **HTTP**, **JSON**, **API** capable of the following operations,<br>

* Create A User
  * [X] Should Be a POST request
  * [X] Should use JSON request body
  * [X] URL should be '/users'
* Get a user using id
  * [X] Should be a GET reuqest
  * [X] Id Should be in URL parameter
  * [X] URL should be '/users/*user id*
* Create a Post
  * [X] Should be a POST request
  * [X] Use JSON request body
  * [X] URL should be ‘/posts'

* Get a post using id
  * [X] Should be a GET request
  * [X] Id should be in the url parameter
  * [X] URL should be ‘/posts/*post id*

* List all posts of a user
  * [ ] Should be a GET request                  
  * [ ] URL should be ‘/posts/users/*user id* 😅
- [X] password Protection  (I used SHA 512 Hashing algorithm to protect User Password)
- [ ] pagination
- [ ] unit Test   (I will try these 2 things before 11:59pm)
# Constraints
- [X] The API should be developed using [Go](https://golang.org/).
- [X] [MongoDB](https://www.mongodb.com/) should be used for storage.
- [X] Only packages/libraries listed [here](https://pkg.go.dev/std) and [here](https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.4.0) Can be used only.



## Setup
## Since I am a linux user, I will give the direction in that perspective only.
- install [golang](https://golang.org/doc/install) and [MongoDB](https://www.mongodb.com/try/download/community)
- create workspace only in the ''go/src''
- install driver from [here](https://github.com/mongodb/mongo-go-driver)
- in terminal run 
```
sudo apt install dep
dep init
```


## How It Works

1. Open Workspace
2. Open internal terminal using **crtl + `**
3. run command in the terminal
```
go run .
```
This thing must be visible in the terminal window :- 

<img src= "https://github.com/TanmayKumar-EngStud/myGolang-MongoDB/blob/master/images/terminal%20Start.png">
<br>
# After running the go files you can Test for API via Postman in this way :
<img src= "https://github.com/TanmayKumar-EngStud/myGolang-MongoDB/blob/master/images/finally.png">
<br>
<img src= "https://github.com/TanmayKumar-EngStud/myGolang-MongoDB/blob/master/images/post%20POST.png">
<br>
# In AppointyDatabase.users
<br>
<img src= "https://github.com/TanmayKumar-EngStud/myGolang-MongoDB/blob/master/images/New%20Data%20is%20added%20in%20the%20This%20database.png">

