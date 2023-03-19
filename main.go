package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"example/web-service-gin/database"
	"example/web-service-gin/structs"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserById)
	router.POST("/users", addUser)
	router.DELETE("/users", delUserById)

	router.Run("localhost:8081")
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Temp_db)
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
	fmt.Printf("ID: %+v", id)
	for _, user := range database.Temp_db {
		if user.Id == int32(id) {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func delUserById(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return 
	}

	delUser := &structs.User{}
	err = json.Unmarshal(data, delUser)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return 
	}

	for idx, user := range database.Temp_db {
		if user.Id == delUser.Id {
			database.Temp_db[idx] = database.Temp_db[len(database.Temp_db)-1]
			database.Temp_db = database.Temp_db[:len(database.Temp_db)-1]
			c.IndentedJSON(http.StatusOK, database.Temp_db)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}
