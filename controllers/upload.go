package controllers

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

func (this *UploadController) Post() {
	_, header, err := this.GetFile("mypic")
	if err != nil {
		this.Ctx.WriteString("hi")
		log.Println(err)
		return
	}
	fn := strconv.FormatInt(time.Now().Unix(), 10) + header.Filename
	if !strings.HasSuffix(fn, ".jpg") && !strings.HasSuffix(fn, ".jpeg") && !strings.HasSuffix(fn, ".png") && !strings.HasSuffix(fn, ".bmp") && !strings.HasSuffix(fn, ".gif") {
		this.Ctx.WriteString("hi")
		return
	}

	this.SaveToFile("mypic", UploadPath+fn)
	this.Ctx.WriteString(UploadUrl + fn)
}
