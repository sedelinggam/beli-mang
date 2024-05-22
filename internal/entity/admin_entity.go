package entity

import "time"

type Admin struct {
	ID        string    `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
}

func (a Admin) TableName() string {
	return `admins`
}
