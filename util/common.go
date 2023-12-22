package util

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
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

// InArray 是否在数组中
func InArray[T comparable](item T, slice []T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// ArrayUnique 切片数据去重
func ArrayUnique[T comparable](slice []T) []T {
	mData := make(map[T]struct{})
	arr := make([]T, 0, len(slice))
	for _, val := range slice {
		if _, ok := mData[val]; !ok {
			mData[val] = struct{}{}
			arr = append(arr, val)
		}
	}
	return arr
}

// ArrayToMap 将切片转为map
func ArrayToMap[T comparable](data []T) map[T]struct{} {
	m := make(map[T]struct{})
	for _, _val := range data {
		if _, ok := m[_val]; !ok {
			m[_val] = struct{}{}
		}
	}
	return m
}

// ArrayFilter 切片根据条件过滤
func ArrayFilter[T comparable](data []T, fn func(item T) bool) []T {
	res := make([]T, 0)
	for _, _val := range data {
		match := fn(_val)
		if match == true {
			res = append(res, _val)
		}
	}
	return res
}

// ArrayFind 切片查找指定条件的数据
func ArrayFind[T comparable](data []T, fn func(item T) bool) *T {
	for _, _val := range data {
		match := fn(_val)
		if match == true {
			return &_val
		}
	}
	return nil
}

// ArrayFindIdx 切片查找指定条件数据的索引，若未找到则为-1
func ArrayFindIdx[T comparable](data []T, fn func(item T) bool) int {
	for _k, _val := range data {
		match := fn(_val)
		if match == true {
			return _k
		}
	}
	return -1
}

// ArrayMap 生成对应格式的数据切片
func ArrayMap[T comparable, E comparable](data []T, fn func(item T) E) []E {
	ret := make([]E, 0, len(data))
	for _, _val := range data {
		tmpVal := fn(_val)
		ret = append(ret, tmpVal)
	}
	return ret
}

// GetRandNum 获取一个随机数
func GetRandNum(min, max int) int {
	//rand.Seed(time.Now().UnixNano()) //rand.Seed在go高版本中弃用了
	rand.NewSource(time.Now().UnixNano())
	//[min,max)包含最小值,不包含最大值
	return rand.Intn(max-min) + min
}

// GetRandNums 获取多个随机数
func GetRandNums(min, max, num int) []int {
	ret := make([]int, 0, num)
	for i := 0; i < num; i++ {
		//time.Sleep(time.Nanosecond * 1)
		tmpNum := GetRandNum(min, max)
		ret = append(ret, tmpNum)
	}
	return ret
}

// IsOpenId 判断是否是openid
func IsOpenId(openid string) bool {
	if m, _ := regexp.MatchString(`^[_A-Za-z0-9\-]{1,40}$`, openid); !m {
		return false
	}
	return true
}

// IsAppId 判断是否是appid
func IsAppId(appid string) bool {
	if m, _ := regexp.MatchString(`^wx[_A-Za-z0-9\-]{10,20}$`, appid); !m {
		return false
	}
	return true
}

// IsNumeric 判断字符串是否是数字字符串
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// SubString 截取字符串子字符串
func SubString(s string, start int, length int) string {
	runes := []rune(s)
	end := start + length
	if start < 0 {
		start = 0
	}
	if end > len(runes) {
		end = len(runes)
	}
	return string(runes[start:end])
}

// StructToMap Convert struct to map.
// This function first tries to use the bson tag, and if the bson tag does not exist, it will use the json tag.
// if both bson and json tags do not exist, then it will use the field name as the key. Additionally,
// if the tag value is "-", this field will be ignored and not added to the map.
// https://github.com/mix-go/mix/blob/master/src/xutil/xconv/conv.go 采用的
// go语言里面struct里面变量如果大写则是public,如果是小写则是private的，private的时候通过反射不能获取其值
// struct里面需要都是大写不然panic: reflect.Value.Interface: cannot return value obtained from unexported field or method
func StructToMap(i interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	inputType := reflect.TypeOf(i)
	inputVal := reflect.ValueOf(i)

	if inputType.Kind() == reflect.Ptr {
		inputType = inputType.Elem()
		inputVal = inputVal.Elem()
	}

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputVal.Field(i).Interface()

		key := field.Tag.Get("bson")
		if key == "" {
			key = field.Tag.Get("json")
		}

		if key == "-" {
			continue
		}

		key = strings.Split(key, ",")[0]
		if key == "" {
			key = field.Name
		}

		result[key] = value
	}

	return result
}

// Pack 二进制数据打包
func Pack(binData []byte) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	length := int32(len(binData))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	//写入实体消息
	err = binary.Write(pkg, binary.LittleEndian, []byte(binData))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Unpack 消息解包
func Unpack(reader *bufio.Reader) (string, error) {
	// 读取消息的长度
	lengthByte, _ := reader.Peek(4)
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// 读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}

func UnpackByte(binData []byte) ([]byte, error) {
	// 读取消息的长度
	dataBuff := bytes.NewReader(binData)
	var (
		msgId  int16 //2个字节消息ID
		length int32 //4个字节的消息长度
	)
	//读MsgID
	if err := binary.Read(dataBuff, binary.BigEndian, &msgId); err != nil {
		return nil, err
	}
	//读消息长度
	if err := binary.Read(dataBuff, binary.BigEndian, &length); err != nil {
		return nil, err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	pack := make([]byte, int(length))
	if err := binary.Read(dataBuff, binary.BigEndian, &pack); err != nil {
		return nil, err
	}
	// 读取真正的消息数据
	return pack, nil
}

func TT(byteContent []byte) {
	//byte与*bufio.Reader的转换
	rd := bytes.NewReader(byteContent)
	r := bufio.NewReader(rd)
}
