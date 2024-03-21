package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hex4coder/user-service/config"
	"github.com/hex4coder/user-service/pkg/handlers"
)

/*
**
Membuat custom middleware untuk API
*/
func APIMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// set content type for json
		c.Header("Content-Type", "application/json")
		// next
		c.Next()
	}
}

/*
**
Membuat fungsi untuk mapping handler ke models method
*/
func SetupUserAPI() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(APIMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// create/post user
	r.POST("/create-user", handlers.CreateUser)

	// get all users
	r.GET("/get-users", handlers.GetUsers)

	// get specific user by id
	r.GET("/get-user/:id", handlers.GetUser)

	// update user
	r.PUT("/update-user/:id", handlers.UpdateUser)

	// delete a user by id
	r.DELETE("/delete-user/:id", handlers.DeleteUser)

	// return router
	return r
}

func Run(app *config.AppConfig) {
	fmt.Printf("[User Service] Running server on endpoint %s\n", app.BackendServer.GetEndpoint())
	log.Fatal(http.ListenAndServe(app.BackendServer.GetEndpoint(), SetupUserAPI()))
}
