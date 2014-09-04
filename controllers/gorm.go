package controllers

import (
	"fmt"
	"os"

	"miniwiki/models"

	"github.com/Unknwon/goconfig"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type GormController struct {
	*beego.Controller
	db *gorm.DB
}

var (
	DB gorm.DB
)

var UploadPath string
var Prefix string
var UploadUrl string

// 自動マイグレーションを行う
func InitDB() {
	var err error

	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver == "" {
		dbDriver = "sqlite3"
	}

	dbSource := os.Getenv("DB_SOURCE")
	if dbSource == "" {
		dbSource = "./wiki.db"
	}

	DB, err = gorm.Open(dbDriver, dbSource)

	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}

	c, err := goconfig.LoadConfigFile("wiki.conf")
	if err != nil {
		panic(fmt.Sprintf("Read config error: %s", err))
	}
	Prefix = c.MustValue("", "prefix")
	UploadPath = c.MustValue("", "upload")
	UploadUrl = c.MustValue("", "uploadurl")
	if UploadPath == "" {
		UploadPath = "./"
	}
	if UploadUrl == "" {
		UploadUrl = "/static/upload/"
	}

	DB.AutoMigrate(models.Page{})
	DB.AutoMigrate(models.Revision{})

	DB.Model(models.Page{}).AddUniqueIndex("unique_title", "title")
}

// リクエスト時にトランザクションを開始する
func (c *GormController) Begin() {
	c.db = DB.Begin()
}

// リクエスト終了時にトランザクションを確定する
func (c *GormController) Commit() {
	if c.db != nil {
		c.db.Commit()
	}
	c.db = nil
}

// 異常時にトランザクションを破棄する
func (c *GormController) Rollback() {
	if c.db != nil {
		c.db.Rollback()
	}
}
