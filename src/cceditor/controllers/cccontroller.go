package controllers

import "github.com/astaxie/beego"

type Cccontroller struct {
	beego.Controller
}

func (c *Cccontroller) Index() {
	c.TplName = "ccindex.html"

}
