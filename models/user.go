package models

import(
	"time"
	_"regexp"
)

type User struct {
	Id int64
	Name string `sql:"size:255"`
	Password string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAT time.Time
	DeletedAt time.Time
}
