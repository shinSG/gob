package controllers

import (
	"encoding/json"
	"gob/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// NoteController is note controller
type NoteController struct {
	beego.Controller
}

// Get function is GET METHOD
func (c *NoteController) Get() {
	note := new(models.Note)
	var res []models.Note
	page, _ := c.GetInt("page", 0)
	pageSize, _ := c.GetInt("pagesize", 0)
	// queryStr := con.GetString("query")
	o := orm.NewOrm()
	qs := o.QueryTable(note)
	if pageSize > 0 {
		qs = qs.Limit(pageSize).Offset((page - 1) * pageSize)
	}
	qs.All(&res)
	c.Data["json"] = &res
	c.ServeJSON()
}

// Post Create Note
func (c *NoteController) Post() {
	var note models.Note
	o := orm.NewOrm()
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &note)
	msg := make(map[string]string)
	if err == nil {
		noteID, err := o.Insert(&note)
		if err == nil {
			msg["ID"] = strconv.FormatInt(noteID, 10)
		} else {
			msg["msg"] = "Insert Error " + err.Error()
		}
	} else {
		msg["msg"] = "Insert Error " + err.Error()
	}
	c.Data["json"] = msg
	c.ServeJSON()
}
