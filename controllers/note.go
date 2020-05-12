package controllers

import (
	"encoding/json"
	"fmt"
	"gob/models"
	"strconv"
	"time"

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
	// var res []models.Note
	page, _ := c.GetInt("page", 0)
	pageSize, _ := c.GetInt("pagesize", 0)
	// queryStr := con.GetString("query")
	o := orm.NewOrm()
	notesList := o.QueryTable(note)
	if pageSize > 0 {
		fmt.Println("#############")
		notesList = notesList.Limit(pageSize, page*pageSize)
		fmt.Println(notesList.Count())
	}
	// notesList.All(&res)
	c.Data["json"] = &notesList
	c.ServeJSON()
}

func (c *NoteController) Post() {
	var note models.Note
	o := orm.NewOrm()
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &note)
	if err == nil {
		note.CreateTime = time.Now()
		note.UpdateTime = time.Now()
		noteID, err := o.Insert(&note)
		if err == nil {
			c.Data["json"] = "{\"Note ID\": \"" + strconv.FormatInt(noteID, 10) + "\"}"
		} else {
			c.Data["json"] = "{\"msg\": \"Insert Error " + err.Error() + "\"}"
		}
	} else {
		c.Data["json"] = "{\"msg\": \"Insert Error\"}"
	}
	c.ServeJSON()
}
