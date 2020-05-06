package auth

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// A sample user
var SuperUser = User{
	Username: "arjun",
	Password: "qwerty1234",
}
