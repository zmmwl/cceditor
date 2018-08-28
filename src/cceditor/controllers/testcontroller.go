package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type TestHandler struct {
	beego.Controller
}



func (c *TestHandler) GetFileHandler() {
	fmt.Println("test handler")
	d := c.Data;
	fmt.Println(d)
	fmt.Println(c.Data["p1"])
	fmt.Println(c.GetString("p1"))
	fmt.Println("CopyRequestBody is: ",beego.BConfig.CopyRequestBody)
	r := c.Ctx.Input;
	fmt.Println(r.RequestBody)

	fmt.Println(string(c.Ctx.Input.RequestBody))
}