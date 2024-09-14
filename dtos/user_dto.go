package dtos

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type Client struct {
	User         User         `json:"user"`
	StatusClient StatusClient `json:"status"`
}

type StatusClient struct {
	Status string `json:"status"`
	Points int    `json:"points"`
}
