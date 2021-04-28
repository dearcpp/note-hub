package model

import "time"

type Note struct {
	ID      uint32    `xorm:"'id' bigint pk" json:"id"`
	Name    string    `xorm:"varchar(64) not null" json:"name"`
	Content string    `xorm:"varchar(256) not null" json:"content"`
	Author  User      `xorm:"not null" json:"author"`
	Created time.Time `xorm:"created" json:"-"`
	Updated time.Time `xorm:"updated" json:"-"`
}
