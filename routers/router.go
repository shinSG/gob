package routers

import (
	"gob/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/note", &controllers.NoteController{})
}
