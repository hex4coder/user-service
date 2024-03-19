package main

import (
	"fmt"
	"net/http"

	"github.com/hex4coder/user-service/pkg/router"
)

func main() {

	port := ":9000"

	fmt.Println("user-service running on port", port)
	http.ListenAndServe(port, router.SetupUserAPI())
}
