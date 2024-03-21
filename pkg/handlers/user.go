package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hex4coder/user-service/database"
	"github.com/hex4coder/user-service/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type UserPostData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Address  string `json:"address"`
}

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
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

	// validate data
	err := validate.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation error",
			"error":   err.Error(),
		})
		return
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
	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
	// get all data from database
	users := []models.User{}
	database.DB.Find(&users)
	// return it
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	// get user id
	id := c.Param("id")

	// find in database
	user := &models.User{}
	database.DB.Find(user, id)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	// return it
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	// get user id
	id := c.Param("id")

	// find in database
	user := &models.User{}
	database.DB.Find(user, id)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	// bind params to user
	// get user from user body
	var userPost UserPostData
	c.Bind(&userPost)

	// update user
	user.Email = userPost.Email
	user.Name = userPost.Name
	user.Role = userPost.Role
	user.Address = userPost.Address

	// validate data
	err := validate.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation error",
			"error":   err.Error(),
		})
		return
	}

	if userPost.Password != "" {
		// hash the password
		user.Password = userPost.Password
		hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed to hash the password",
				"error":   err.Error(),
			})
			return
		}
		user.Password = string(hashed)
	}

	// update user
	result := database.DB.Save(user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update user",
			"error":   result.Error.Error(),
		})
		return
	}

	// return it
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	// get user id
	id := c.Param("id")

	// find in database
	user := &models.User{}
	database.DB.Find(user, id)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	// delete it
	result := database.DB.Delete(user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to delete user",
			"error":   result.Error.Error(),
		})
		return
	}

	// return it
	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
		"id":      id,
	})
}

/*
**
Login ke sistem dengan email dan password
*/
func LoginUser(c *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// mapping data req
	var req LoginReq
	c.Bind(&req)

	// validate data
	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Credential not valid",
			"message": "Request empty",
		})
		return
	}

	// check email in database
	var user models.User
	database.DB.First(&user, "email = ?", req.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Credential not valid",
			"message": "User not found",
		})
		return
	}

	// check Password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Credential not valid",
			"message": fmt.Sprintf("Invalid password : %s", err.Error()),
		})
		return
	}

	// return it
	c.JSON(http.StatusOK, user)
}
