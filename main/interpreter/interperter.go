package interperter

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"vue/main/data"
	"vue/main/utils"
)

type Vue struct {
	Name   string
	File   *os.File
	Json   map[string]interface{}
	Html   strings.Builder
	Data   map[string]interface{}
	Method strings.Builder
	Css    strings.Builder
}

func (vue *Vue) Run() {
	vue.Data["formStruct"] = make(map[string]interface{})
	vue.Data["formData"] = make(map[string]interface{})
	vue.coreInterperter()
}

/**
核心解释器
*/
func (vue *Vue) coreInterperter() {
	//解释生成子结构
	for _, v := range data.Fields {
		println(v)
		vue.typeInterperter(v, vue.Json[v])
	}

	//解释生成主结构
	vue.mainInterperter()
	fmt.Println("代码生成结束！")
}

/**
主文件最后处理
*/
func (vue *Vue) mainInterperter() {
	main := utils.ReadTemp("main")
	//替换vue组件name
	main = strings.ReplaceAll(main, "${name}", vue.Name)
	//替换html content
	main = strings.ReplaceAll(main, "${content}", vue.Html.String())
	//替换data
	mjson, _ := json.Marshal(vue.Data)
	main = strings.ReplaceAll(main, "${data}", string(mjson))

	vue.File.WriteString(main)
}

/**
类型转换
*/
func (vue *Vue) typeInterperter(k string, v interface{}) {
	// 其他项处理
	var mapV = v.(map[string]interface{})
	switch mapV["type"] {
	case "input":
		vue.inputInterperter(k, mapV)
		break
	case "select":
		vue.selectInterperter(k, mapV)
		break
	case "card":
		vue.cardInterperter(mapV)
		break
	case "checkbox":
		vue.checkboxInterperter(k, mapV)
		break
	}

}

/**
card
*/
func (vue *Vue) cardInterperter(v map[string]interface{}) {
	card := utils.ReadTemp("card")
	var elementLabel = v["label"]
	card = strings.ReplaceAll(card, "${label}", elementLabel.(string))
	vue.Html.WriteString(card)
}

/**
普通输入框解释器
*/
func (vue *Vue) inputInterperter(k string, v map[string]interface{}) {
	//var elementType = v["type"]
	var elementLabel = v["label"]
	//var elementData = v["data"]
	//var elementNotHave = v["notHave"]
	str := utils.ReadTemp("input")
	str = strings.ReplaceAll(str, "${v-model}", "formData."+k)
	str = strings.ReplaceAll(str, "${label}", elementLabel.(string))
	vue.Html.WriteString(str)
	var formData = vue.Data["formData"].(map[string]interface{})
	formData[k] = nil
}

/**
下拉选择
*/
func (vue *Vue) selectInterperter(k string, v map[string]interface{}) {
	//var elementType = v["type"]
	var elementLabel = v["label"]
	//var elementData = v["data"]
	//var elementNotHave = v["notHave"]
	str := utils.ReadTemp("select")
	//
	str = strings.ReplaceAll(str, "${v-model}", "formData."+k)
	str = strings.ReplaceAll(str, "${label}", elementLabel.(string))
	str = strings.ReplaceAll(str, "${items}", "formStruct."+k)
	vue.Html.WriteString(str)
	var formData = vue.Data["formData"].(map[string]interface{})
	formData[k] = nil
}

func (vue *Vue) checkboxInterperter(k string, v map[string]interface{}) {
	//var elementType = v["type"]
	var elementLabel = v["label"]
	//var elementData = v["data"]
	//var elementNotHave = v["notHave"]
	str := utils.ReadTemp("checkbox")
	//
	str = strings.ReplaceAll(str, "${v-model}", "formData."+k)
	str = strings.ReplaceAll(str, "${label}", elementLabel.(string))
	str = strings.ReplaceAll(str, "${k}", k)
	vue.Html.WriteString(str)
	var formData = vue.Data["formData"].(map[string]interface{})
	formData[k] = []int{}
}
