package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userInfo struct {
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age" binding:"required"`
	Color string `json:"color" binding:"required"`
}

var users map[string]userInfo

func getUser(c *gin.Context) {
	name := c.Param("name")
	info, ok := users[name]
	if !ok {
		message := fmt.Sprintf("Could not find user %s", name)
		c.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, info)
}

func addUser(c *gin.Context) {
	var input userInfo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users[input.Name] = input
	message := fmt.Sprintf("Added user %s", input.Name)
	c.JSON(http.StatusOK, gin.H{"success": message})
}

func main() {
	users = make(map[string]userInfo, 1)
	router := gin.Default()

	router.GET("/user/:name", getUser)
	router.POST("/user", addUser)

	router.Run(":8080")
}
