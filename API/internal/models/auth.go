package models


type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type About struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Author  string `json:"author"`
	}


type LoginResponse struct {
	Token string `json:"token"`
}

type User struct {
	ID       int64
	Username string
	Password string
}