package models

import (
	"time"
)

type Note struct {
	ID      int64     `xorm:"'id' autoincr pk" json:"id"`
	Name    string    `xorm:"varchar(256) not null" json:"name"`
	Content string    `xorm:"varchar(2048) not null" json:"content"`
	Author  User      `xorm:"not null" json:"author"`
	Created time.Time `xorm:"created" json:"-"`
	Updated time.Time `xorm:"updated" json:"-"`
}
