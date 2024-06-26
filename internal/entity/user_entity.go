package entity

import (
	"strconv"
	"time"

	"github.com/lib/pq"
)

type User struct {
	ID                    string      `db:"id"`
	NIP                   string      `db:"nip"`
	Name                  string      `db:"name"`
	Password              string      `db:"password"`
	UserRole              string      `db:"role"`
	IdentityCardScanImage *string     `db:"identity_card_scan_img"`
	CreatedAt             time.Time   `db:"created_at"`
	DeletedAt             pq.NullTime `db:"deleted_at"`
}

func (s User) TableName() string {
	return `users`
}

func (s User) CheckNIP(login bool) bool {
	//Change string to INT
	//Check if NIP length is 13
	if len(s.NIP) < 13 || len(s.NIP) > 15 {
		return false
	} else {
		year, _ := strconv.Atoi(s.NIP[4:8])
		month, _ := strconv.Atoi(s.NIP[8:10])
		randDigit, _ := strconv.Atoi(s.NIP[10:len(s.NIP)])
		if s.UserRole == "t" && s.NIP[0:3] != "303" && !login {
			return false
		} else if s.UserRole == "v" && s.NIP[0:3] != "615" && !login {
			return false
		} else if s.NIP[3:4] != "1" && s.NIP[3:4] != "2" {
			return false
		} else if year < 2000 || year > time.Now().Year() {
			return false
		} else if month < 1 || month > 12 {
			return false
		} else if randDigit < 1 || randDigit > 99999 {
			return false
		}
	}
	return true
}
