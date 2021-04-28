package model

import "time"

type User struct {
	ID       uint32    `xorm:"'id' bigint pk" json:"id"`
	Email    string    `xorm:"varchar(320) not null unique" json:"email"`
	Password string    `xorm:"varchar(32) not null" json:"-"`
	Created  time.Time `xorm:"created" json:"-"`
	Updated  time.Time `xorm:"updated" json:"-"`
}
