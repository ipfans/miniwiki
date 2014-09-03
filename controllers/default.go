package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Redirect(this.UrlFor("PageController.Get", ":page", ""), 302)
}
