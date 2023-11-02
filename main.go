package main

// Bobobek Turd
import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	InitDB()

	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		users, err := GetAllUsers()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, users)
	})

	r.GET("/users/:id", func(c *gin.Context) {

		idStr := c.Param("id")
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}
		user, err := GetUserByID(uint(idInt))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, user)
	})

	r.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := CreateUser(&user); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, user)
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := UpdateUser(&user); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, user)
	})

	r.DELETE("/users/:id", func(c *gin.Context) {

		idStr := c.Param("id")
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}

		if err := DeleteUser(uint(idInt)); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "User deleted successfully"})
	})

	r.Run(":8080")
}
