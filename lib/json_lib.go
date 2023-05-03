package lib

import (
	"encoding/json"
	"io/ioutil"
)

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (j *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile()函数会读取文件的全部内容,并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
