package model

type User struct {
	Name 		string	`json:"name"`
	Password 	string	`json:"password"`
}

var UserList []User

func Init() {
	UserDemo := User{
		Name: "thinhnx",
		Password: "123",
	}
	UserList = append(UserList, UserDemo)
}