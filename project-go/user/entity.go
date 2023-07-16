package user

import "time"

type User struct {
	ID          int
	Name        string
	NoHandphone string
	Email       string
	Password    string
	Role        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
