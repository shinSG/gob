package controllers

import (
	"github.com/astaxie/beego"
)

// MainController main page
type MainController struct {
	beego.Controller
}

// Get method GET
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
