package entity

type CreateUserRequest struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
	Mail     string `json:"Mail"`
}
