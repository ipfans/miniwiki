package controllers

import (
	"log"
	"miniwiki/models"
	"miniwiki/wikihelper"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/oal/beego-pongo2"
)

type PageController struct {
	GormController
	beego.Controller
}

func (this *PageController) Get() {
	this.Begin()
	pageName := this.Ctx.Input.Param(":page")
	if pageName == "" {
		pageName = "Home"
	}

	page := models.Page{}
	this.db.Where("id = ?", pageName).Or("title = ?", pageName).First(&page)

	id, _ := strconv.Atoi(pageName)
	if id != 0 && page.Id == id {
		this.Commit()
		this.Redirect(this.UrlFor("PageController.Get", "page", page.Title), 302)
		return
	}

	body := page.Body
	html := wikihelper.Render(body)

	// リビジョン番号を取得
	revision := 0
	this.db.Model(models.Revision{}).Where("page_id = ?", page.Id).Count(&revision)
	log.Println(revision)

	// 最近登録されたページ一覧を取得
	recentCreatedPages := []models.Page{}
	this.db.Order("created_at desc").Limit(10).Find(&recentCreatedPages)

	// 最近更新されたページ一覧を取得
	recentUpdatedPages := []models.Page{}
	this.db.Where("created_at != updated_at").Order("updated_at desc").Limit(10).Find(&recentUpdatedPages)

	this.Commit()
	pongo2.Render(this.Ctx, "page.html", pongo2.Context{
		"title":    pageName,
		"html":     html,
		"revision": revision,
		"prefix":   Prefix,
	})
}
