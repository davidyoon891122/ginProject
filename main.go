package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"./tools"
)

var logger *log.Logger
var currentDirectory string
var logPath string = "/logs/server.log"

func getLogger() {
	f, err := os.OpenFile(currentDirectory+logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	logger = log.New(f, "INFO : ", log.Ldate|log.Ltime|log.Lshortfile)

}

func main() {
	currentDirectory, _ = os.Getwd()

	getLogger()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*") //load template for rendering.

	r.Static("/css", "./css") //load css

	//login
	r.GET("/", login_get_handler)
	r.POST("/", login_post_handler)

	//join
	r.GET("/join", join_get_handler)
	r.POST("/join", join_post_handler)

	r.Run(":8080")

}

func login_get_handler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login Screen",
	})
}

func login_post_handler(c *gin.Context) {
	id := c.PostForm("userID")
	password := c.PostForm("userPWD")

	log.Printf("logger(test) userID : %s \n", id)
	log.Printf("logger(test) userPWD : %s \n", password)

	result := tools.CheckInfo(id, password)

	if result == true {
		c.HTML(http.StatusOK, "login_post.html", gin.H{
			"id":       id,
			"password": password,
		})
	} else {
		c.String(http.StatusOK, "Access Denied")
	}
}

func join_get_handler(c *gin.Context) {
	c.HTML(http.StatusOK, "join.html", gin.H{
		"title": "Join Screen",
	})
}

func join_post_handler(c *gin.Context) {
	idJoin := c.PostForm("idJoin")
	passwordJoin := c.PostForm("passwordJoin")
	nameJoin := c.PostForm("nameJoin")
	numberJoin := c.PostForm("numberJoin")

	log.Printf("logger(test) idJoin : %s \n", idJoin)
	log.Printf("logger(test) passwordJoin : %s \n", passwordJoin)
	log.Printf("logger(test) nameJoin : %s \n", nameJoin)
	log.Printf("logger(test) numberJoin : %s \n", numberJoin)

}
