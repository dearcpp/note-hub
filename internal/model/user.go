package model

import (
	"time"
)

type User struct {
	ID       int64     `xorm:"'id' autoincr pk" json:"-"`
	Email    string    `xorm:"varchar(320) not null unique" json:"email"`
	Password string    `xorm:"varchar(32) not null" json:"-"`
	Created  time.Time `xorm:"created" json:"-"`
	Updated  time.Time `xorm:"updated" json:"-"`
}
