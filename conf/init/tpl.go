package init

import (
	"github.com/astaxie/beego"
)

func init() {
	beego.BConfig.WebConfig.TemplateLeft = "<<<"
	beego.BConfig.WebConfig.TemplateRight = ">>>"
}
