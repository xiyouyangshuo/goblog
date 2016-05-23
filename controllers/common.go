package controllers

import (
	"time"

	"github.com/astaxie/beego"
	// "github.com/deepzz0/go-com/log"
	"github.com/deepzz0/goblog/cache"
	"github.com/deepzz0/goblog/models"
)

var sessionname = beego.AppConfig.String("sessionname")

type Common struct {
	beego.Controller
}

func (this *Common) Prepare() {
	this.Layout = "homelayout.html"
	this.Build()
	this.DoRequest()
}
func (this *Common) Leftbar(cat string) {
	this.Data["Picture"] = models.Blogger.HeadIcon
	this.Data["BlogName"] = models.Blogger.BlogName
	this.Data["Introduce"] = models.Blogger.Introduce
	this.Data["Categories"] = models.Blogger.Categories
	this.Data["Socials"] = models.Blogger.Socials
	this.Data["Choose"] = cat
	this.Data["CopyTime"] = time.Now().Year()
}

func (this *Common) Build() {
	this.Data["Build"] = cache.Cache.BuildVersion
}

func (this *Common) Verification() {
	this.Data["Verification"] = models.ManageConf.SiteVerify
}

func (this *Common) DoRequest() {
	requst := models.NewRequest(this.Ctx.Request)
	models.RequestM.Ch <- requst
}

type listOfTopic struct {
	ID        int32
	Title     string
	URL       string
	Time      string
	Preview   string
	PCategory *models.Category
	PTags     []*models.Tag
}
