package router

import (
	"net/http"

	"github.com/hex4coder/user-service/pkg/handlers"
)

/*
**
Membuat fungsi untuk mapping handler ke models method
*/
func SetupUserAPI() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/get-user", handlers.GetUserHandler)
	router.HandleFunc("/create-user", handlers.CreateUserHandler)
	router.HandleFunc("/update-user", handlers.UpdateUserHandler)
	router.HandleFunc("/delete-user", handlers.DeleteUserHandler)

	return router
}
