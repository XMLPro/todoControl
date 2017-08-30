package controllers

import (
	"github.com/astaxie/beego"
)

type TaskController struct {
	beego.Controller
}

func (c *TaskController) Get() {
	c.TplName = "index.html"
	//c.Ctx.WriteString("index pages")
}
