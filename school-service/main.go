package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int
	Username string
	School   string
}

var users []User

func main() {
	router := gin.Default()

	router.GET("/randomuser", randomUser)
	router.GET("/getuser/:id", getUser)
	router.POST("/adduser", addUser)

	router.Run(":8080")
}

func randomUser(c *gin.Context) {
	if len(users) > 0 {
		user := users[rand.Intn(len(users))]
		c.JSON(http.StatusOK, gin.H{
			"user": gin.H{
				"Id":       user.Id,
				"Username": user.Username,
				"School":   user.School,
			},
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"user": gin.H{
			"Id": -1,
		},
		"message": "No users exist",
	})
}

func addUser(c *gin.Context) {
	username := c.PostForm("username")
	school := c.PostForm("school")

	if len(username) > 0 && len(school) > 0 {
		newuser := User{
			Id:       len(users),
			Username: username,
			School:   school,
		}
		users = append(users, newuser)

		msg := "Created user " + username + " (" + strconv.Itoa(newuser.Id) + ") with school " + school
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": msg,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failure",
			"message": "POST request must include username and school fields",
		})
	}
}

func getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if len(users) > id && err == nil {
		user := users[id]
		c.JSON(http.StatusOK, gin.H{
			"user": gin.H{
				"Id":       user.Id,
				"Username": user.Username,
				"School":   user.School,
			},
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"user": gin.H{
			"Id": -1,
		},
		"message": "No users exist",
	})
}
