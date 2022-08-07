package domain

type User struct {
	Id             string
	Email          string
	Password       string
	RepeatPassword string `json:"repeat_password"`
}
