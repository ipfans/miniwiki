package routers

import (
	"miniwiki/controllers"

	"github.com/astaxie/beego"
)

func init() {
	controllers.InitDB()

	beego.Router(controllers.Prefix+"/upload/", &controllers.UploadController{})
	beego.Router(controllers.Prefix+"/page/", &controllers.PageController{})
	beego.Router(controllers.Prefix+"/page/:page", &controllers.PageController{})
	beego.Router(controllers.Prefix+"/page/:page/modify", &controllers.ModifyController{})
	beego.Router(controllers.Prefix+"/", &controllers.MainController{})
}
