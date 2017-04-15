package controllers //控制器

import (
	"github.com/astaxie/beego"
)

type MainController struct { //声明了一个控制器MainController
	beego.Controller //控制器里面内嵌了beego.Controller
} //也就是 MainController自动拥有了所有beego.Controller 的方法

/*
浏览器的是 GET 请求，那么默认就会执行MainController下的 Get 方法
*/
func (c *MainController) Get() { //重写的方式,重写Get方法
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	//	c.Ctx.WriteString("Hello")   //直接用 this.Ctx.WriteString 输出字符串,没有调用模板

	/*
		这里只是简单的输出数据，我们可以通过各种方式获取数据，然后赋值到 c.Data 中，
		这是一个用来存储输出数据的 map，可以赋值任意类型的值，
	*/

	c.TplName = "index.tpl" //最后一个就是需要去渲染的模板
}
