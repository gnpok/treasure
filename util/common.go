package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// Reversal 反转字符串
func Reversal(a string) (re string) {
	//将字符串转为rune数组
	b := []rune(a)
	//遍历
	for i := 0; i < len(b)/2; i++ {
		//交换
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	//转换未字符串类型
	re = string(b)
	return
}

// GenRandStr 生成随机字符串
func GenRandStr(n int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	totalNumber := len(str)

	var token = strings.Builder{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		position := rand.Intn(totalNumber)
		token.WriteByte(str[position])
	}
	return token.String()
}

// GenSmsCode 生成短信验证码
func GenSmsCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

// EncodeMD5 md5加盐加密
func EncodeMD5(val, salt string) string {
	h := md5.New()
	io.WriteString(h, salt)
	io.WriteString(h, val)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// IsMobile 判断是否是手机号码
func IsMobile(phone string) bool {
	if m, _ := regexp.MatchString(`^1[0-9]{10}$`, phone); !m {
		return false
	}
	return true
}

// IsEmail 判断是否是有效的邮箱
// http://www.jsons.cn/reg/ 这个地址里面有更多的正则内容
func IsEmail(email string) bool {
	if m, _ := regexp.MatchString(`^[a-zA-Z0-9_]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+$`, email); !m {
		return false
	}
	return true
}

// StrFilterGetChinese 获取中文字符串
func StrFilterGetChinese(src string) string {
	cnRegexp := regexp.MustCompile("^[\u4e00-\u9fa5]$")
	strNew := ""
	for _, c := range src {
		if cnRegexp.MatchString(string(c)) {
			strNew += string(c)
		}
	}
	return strNew
}

// ExistStr 检查字符串是否在切片中
func ExistStr(target string, array []string) bool {
	for _, item := range array {
		if target == item {
			return true
		}
	}
	return false
}

// ArraySearch 查找一个元素在数组中的位置
func ArraySearch(arr interface{}, d interface{}) int {
	array := reflect.ValueOf(arr)
	num := array.Len()
	for i := 0; i < num; i++ {
		v := array.Index(i)
		if v.Interface() == d {
			return i
		}
	}
	return -1
}

// ArrayUnique 删除元素中重复的
func ArrayUnique(array []int) []int {
	//如果是空切片，则返回nil
	if len(array) == 0 {
		return nil
	}
	//用两个标气来比较相邻位置的值
	//如果一样，则继续
	//如果不一样,则把right指向的赋值给left下一位
	left, right := 0, 1
	for ; right < len(array); right++ {
		if array[left] == array[right] {
			continue
		}
		left++
		array[left] = array[right]
	}
	return array[:left+1]
}

// GetRandNum 获取一个随机数
func GetRandNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	//[min,max)包含最小值,不包含最大值
	return rand.Intn(max-min) + min
}

// GetRandNums 获取多个随机数
func GetRandNums(min, max, num int) []int {
	ret := make([]int, 0, num)
	for i := 0; i < num; i++ {
		time.Sleep(time.Nanosecond * 1)
		tmpNum := GetRandNum(min, max)
		ret = append(ret, tmpNum)
	}
	return ret
}
