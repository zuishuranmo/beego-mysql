package main

import (
	"github.com/astaxie/beego"
	_ "sgwe/db_mysql"
	_ "sgwe/routers"
)

func main() {
	beego.Run()
}


