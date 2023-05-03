package lib

import (
	"fmt"
	"testing"
)

type Conf struct {
	Name string
	Age  int
}

func TestNewJsonStruct(t *testing.T) {
	v := Conf{}
	j := NewJsonStruct()
	j.Load("./conf.json", &v)
	fmt.Println(v)
}
