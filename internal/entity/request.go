package entity

type Request struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
	Mail     string `json:"Mail"`
}
