package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Note Model
type Note struct {
	ID          int
	Title       string
	Content     string
	HTMLContent string
	CreateTime  time.Time
	UpdateTime  time.Time
}

func init() {
	orm.RegisterModel(new(Note))
}
