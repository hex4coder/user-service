package handlers

import "net/http"

func CreateUserHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	rw.Write([]byte("create user"))
}

func UpdateUserHandler(rw http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	rw.Write([]byte("update user"))
}
func DeleteUserHandler(rw http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	rw.Write([]byte("delete user"))
}
func GetUserHandler(rw http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	rw.Write([]byte("get user"))
}
