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
	CreateTime  time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime  time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Note))
}
