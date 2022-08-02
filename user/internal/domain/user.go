package domain

import "github.com/golang/protobuf/ptypes/timestamp"

type User struct {
	Id               string
	Email            string
	Password         string
	LastLogin        *timestamp.Timestamp
	RegistrationDate *timestamp.Timestamp
}
