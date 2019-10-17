package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/**
字符串转json
*/
func JsonConverter(str string) map[string]interface{} {
	json1 := make(map[string]interface{})
	var err = json.Unmarshal([]byte(str), &json1)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil
	}
	return json1
}

/**
创建vue文件
*/
func CreateVueFile(name string) *os.File {
	dir, err := os.Getwd()
	path := dir + "/main/out/" + name
	f, err := os.Create(path)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil
	}
	return f
}

/**
模板读取
*/
func ReadTemp(name string) string {
	dir, err := os.Getwd()
	data, err := ioutil.ReadFile(dir + "/main/template/" + name + ".temp")
	if err != nil {
		fmt.Println("File reading error", err)
	}
	return string(data)
}
