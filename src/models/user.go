package models

import "time"

type User struct {
	ID       int64     `xorm:"'id' autoincr pk" json:"id"`
	Mail     string    `xorm:"varchar(320) not null unique" json:"mail"`
	Password string    `xorm:"varchar(32) not null" json:"-"`
	Created  time.Time `xorm:"created" json:"-"`
	Updated  time.Time `xorm:"updated" json:"-"`
}
