package main

import (
	_ "github.com/XMLPro/todoControl/conf/init"
	_ "github.com/XMLPro/todoControl/controllers"
	_ "github.com/XMLPro/todoControl/models"
	_ "github.com/XMLPro/todoControl/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
