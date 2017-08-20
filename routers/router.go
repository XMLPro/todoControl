package routers

import (
	"github.com/XMLPro/todoControl/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
