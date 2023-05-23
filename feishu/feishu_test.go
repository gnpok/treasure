package feishu

import (
	"fmt"
	"testing"
)

func TestFeiShu_BaseNotify(t *testing.T) {
	feishu := NewFeiShu(DEMO)
	feishu.BaseNotify("测试消息111")
}

func TestFeiShu_ImportanceNotify(t *testing.T) {
	feishu := NewFeiShu(DEMO)
	err := feishu.ImportanceNotify("重要消息", "重要内容")
	fmt.Println(err)
}
