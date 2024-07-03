package controllers

import "github.com/gin-gonic/gin"

var jwtKey = []byte("this_is_a_test_key")

func Login(c *gin.Context) {
	c.JSON(200, gin.H{"Success" : "You were logged in"})
}

func Signup(c *gin.Context) {
	c.JSON(200, gin.H{"Success" : "You were logged in"})
}

func Home(c *gin.Context) {
	c.JSON(200, gin.H{"Success" : "You were logged in"})
}

func Premium(c *gin.Context) {
	c.JSON(200, gin.H{"Success" : "You were logged in"})
}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{"Success" : "You were logged in"})
}

