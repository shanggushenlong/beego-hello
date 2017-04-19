package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func init() {
	//固定路由:就是全匹配的路由
	//常用的路由方式:一个固定的路由,一个控制器,然后根据用户请求方法的不同请求控制器中对应的方法,典型的RESTful方式
	beego.Router("/", &controllers.MainController{})

	beego.Router("/api/list", &controllers.RestController{}, "post:ListFood")

}
