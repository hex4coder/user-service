package models

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"-"`
	Role     string `json:"role"`
	Address  string `json:"address"`
}

//type NotFound error

func CreateUser() error
func findUser(userID uint) (*User, error)
func UpdateUser(userID uint, updatedUser *User) error
func DeleteUser(userID uint) error
