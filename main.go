package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"./tools"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*") //load template for rendering.

	r.Static("/css", "./css") //load css

	//login
	r.GET("/", login_get_handler)
	r.POST("/", login_post_handler)

	//join

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

}
