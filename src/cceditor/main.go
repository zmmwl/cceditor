package main

import (
	_ "cceditor/routers"
	"github.com/astaxie/beego"
	"github.com/b3log/wide"
)

func main() {
	wide.Start()
	beego.Run()

}

