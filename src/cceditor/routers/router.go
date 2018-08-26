package routers

import (
	"cceditor/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/cc", &controllers.Cccontroller{}, "*:Index")

}
