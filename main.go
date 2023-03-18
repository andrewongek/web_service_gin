package main

import (
	"net/http"
	"strconv"

	"example/web-service-gin/database"
	"example/web-service-gin/structs"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", addUser)
	router.GET("/users/:id", getUserById)

	router.Run("localhost:8080")
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Temp_db)
	return
}

func addUser(c *gin.Context) {
	var newUser structs.User

	if err := c.BindJSON(&newUser); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return 
	}

	database.Temp_db = append(database.Temp_db, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func getUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	
	for _, user := range database.Temp_db {
		if user.Id == int32(id) {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}