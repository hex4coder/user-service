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
Membuat fungsi untuk mapping handler ke models method
*/
func SetupUserAPI() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// create/post user
	r.POST("/create-user", handlers.CreateUser)

	// r.HandleFunc("/get-user", handlers.GetUserHandler)
	// r.HandleFunc("/update-user", handlers.UpdateUserHandler)
	// r.HandleFunc("/delete-user", handlers.DeleteUserHandler)

	return r
}

func Run(app *config.AppConfig) {
	fmt.Printf("[User Service] Running server on endpoint %s\n", app.BackendServer.GetEndpoint())
	log.Fatal(http.ListenAndServe(app.BackendServer.GetEndpoint(), SetupUserAPI()))
}
