package routers

import (
	"miniwiki/controllers"

	"github.com/astaxie/beego"
)

func init() {
	controllers.InitDB()
	beego.Router("/page/", &controllers.PageController{})
	beego.Router("/page/:page", &controllers.PageController{})
	beego.Router("/page/:page/modify", &controllers.ModifyController{})
	beego.Router("/", &controllers.MainController{})
}
