package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	_ "hello/routers"
)

func main() {
	//在配置文件中的配置项都是采用 键值对 的方式,即 key = value
	/*
			通过beego.AppConfig.String()函数获取配置文件中的信息
		    形如 beego.AppConfig.String("key"),key代表配置文件中的配置项目
	*/
	beego.BConfig.AppName = "AAAAAAAA"
	a := beego.AppConfig.String("appname")
	fmt.Println(a)

	b, err := beego.AppConfig.Int("httpport") //多返回值
	if err == nil {
		fmt.Println(b)
	} else {
		fmt.Println(err)
	}

	c := beego.AppConfig.String("runmode")
	fmt.Println(c)

	//json
	/*调用 config.NewConfig()构建不同格式 config 解析器*/
	aa, err := config.NewConfig("json", "conf/test.json") //初始化文件,解析json,默认未init
	if err == nil {
		val := aa.String("dev::appname")
		fmt.Println("val:", val)

		port, _ := aa.Int("dev::httpport")
		fmt.Println("port:", port)

		str := aa.String("dev::runmode")
		fmt.Println(str)

	} else {
		fmt.Println(err)
	}

	fmt.Println()
	fmt.Println()

	//api   初始化默认ini配置文件,config
	configTest, err := config.NewConfig("ini", "conf/app.conf")
	val_bool, err := configTest.Bool("test_bool")
	fmt.Println(val_bool)

	val_int, _ := configTest.Int("test_int")
	fmt.Println(val_int)

	val_int64, _ := configTest.Int("test_int64")
	fmt.Println(val_int64)

	val_string := configTest.String("test_string")
	fmt.Println(val_string)

	testStr1 := configTest.String("appname")
	fmt.Println(testStr1)

	//配置文件的设置
	configTest.Set("appname", "hello")
	testStr2 := configTest.String("appname")
	fmt.Println(testStr2)

	//default
	def := configTest.DefaultString("hello", "hello jikexueyuan") //设置默认值
	fmt.Println(def)
	//save 文件保存
	configTest.SaveConfigFile("conf/configTest.conf")

	/*
	   BConfig是程序在整个运行过程中都需要用的,而 AppConfig是用来解析配置项的接口,
	   所以整个配置项的解析过程来说就是 使用 AppConfig去完成 BConfig的初始化
	*/
	beego.Run()
}
