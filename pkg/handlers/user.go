package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hex4coder/user-service/database"
	"github.com/hex4coder/user-service/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type UserPostData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
	Address  string `json:"address"`
}

func CreateUser(c *gin.Context) {
	// get user from user body
	var userPost UserPostData
	c.Bind(&userPost)

	// create user
	user := models.User{
		Password: userPost.Password,
		Email:    userPost.Email,
		Name:     userPost.Name,
		Role:     userPost.Role,
		Address:  userPost.Address,
	}

	// hash the password
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to hash the password",
			"error":   err.Error(),
		})
		return
	}
	user.Password = string(hashed)

	// post to database
	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to insert user to database",
			"error":   result.Error.Error(),
		})
		return
	}

	// return it
	c.JSON(200, user)

}
