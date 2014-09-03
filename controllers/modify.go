package controllers

import (
	"crypto/sha1"
	"fmt"
	"log"
	"miniwiki/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/oal/beego-pongo2"
	"github.com/pmezard/go-difflib/difflib"
)

type ModifyController struct {
	beego.Controller
	GormController
}

func (this *ModifyController) Get() {
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
	hash := fmt.Sprintf("%x", sha1.Sum([]byte(body)))
	this.Commit()
	pongo2.Render(this.Ctx, "modify.html", pongo2.Context{
		"title": pageName,
		"body":  body,
		"hash":  hash,
	})
}

func (this *ModifyController) Post() {
	this.Begin()
	pageName := this.Ctx.Input.Param(":page")
	if pageName == "" {
		pageName = "Home"
	}
	page := models.Page{}
	this.db.Where("title = ?", pageName).First(&page)

	// POSTで送信された本文を取得
	body := this.GetString("page.Body")
	//body := this.Ctx.Input.Params["page.Body"]
	log.Println(body)

	// ページは存在するが変更が一切ない場合には更新しない
	if page.Id > 0 && page.Body == body {
		this.Commit()
		this.Redirect(this.UrlFor("PageController.Get", ":page", page.Title), 302)
		return
	}

	// ページを保存する
	page.Title = pageName
	page.Body = body
	this.db.Save(&page)

	// 最新のリビジョンを取得
	previous := models.Revision{}
	this.db.Where("page_id = ?", page.Id).Order("id desc").First(&previous)

	// 追加行、削除行を数えるため差分を取得
	unifiedDiff := difflib.UnifiedDiff{
		A:       difflib.SplitLines(previous.Body),
		B:       difflib.SplitLines(page.Body),
		Context: 65535,
	}
	diffString, _ := difflib.GetUnifiedDiffString(unifiedDiff)
	diffLines := difflib.SplitLines(diffString)

	// 追加行、削除行を数える
	revision := models.Revision{}
	for i, line := range diffLines {
		if i > 2 {
			if strings.HasPrefix(line, "+") {
				revision.AddedLines++
			}
			if strings.HasPrefix(line, "-") {
				revision.DeletedLines++
			}
		}
	}

	// リビジョンを保存
	revision.Title = page.Title
	revision.Body = page.Body
	revision.PageId = page.Id
	this.db.Save(&revision)
	this.Commit()

	this.Redirect(this.UrlFor("PageController.Get", ":page", page.Title), 302)
	return
}
