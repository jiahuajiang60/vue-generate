package main

import (
	"strings"
	"vue/main/data"
	interperter "vue/main/interpreter"
	"vue/main/utils"
)

// 页面名称
const name = "FamilyProfileInfoAdd"

/**
启动程序
*/
func start(data string, _name string) {
	var m = utils.JsonConverter(data)
	f := utils.CreateVueFile(_name)
	vue := &interperter.Vue{
		name,
		f,
		m,
		strings.Builder{},
		make(map[string]interface{}),
		strings.Builder{},
		strings.Builder{}}
	vue.Run()
}

/**
入口函数
*/
func main() {
	var str = data.JSON
	start(str, name+".vue")
}
