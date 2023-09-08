package util

import (
	"fmt"
	"strings"
	"testing"
)

func TestReversal(t *testing.T) {
	str := "123456789"
	after := Reversal(str)
	fmt.Println(after)
}

func TestEncodeMD5(t *testing.T) {
	md5 := EncodeMD5("qweasd", "")
	fmt.Println(md5)
}

func TestStrFilterGetChinese(t *testing.T) {
	str := "我爱GO,棒棒哒"
	chinese := StrFilterGetChinese(str)
	fmt.Println(chinese)
}

func TestArr(t *testing.T) {
	arr := []string{"a", "b", "c"}
	join := strings.Join(arr, "")
	fmt.Println(join)
}

func TestArraySearch(t *testing.T) {
	a := make([]int, 6)
	for i := 0; i < 6; i++ {
		a[i] = i + 2
	}
	fmt.Println(a)
	idx := ArraySearch(a, "1")
	fmt.Println(idx)
}

func TestArrayUnique(t *testing.T) {
	//var array []int = []int{1, 1, 3, 3, 4, 4, 5}
	//unique := ArrayUnique(array)
	//fmt.Println(unique == nil)
	//fmt.Printf("%v ,%p", unique, unique)
	//deep := 2
	//nums := make([]int, 0, 2*deep)
	//for i := 0; i < 2; i++ {
	//	for j := 1; j <= deep; j++ {
	//		nums = append(nums, j)
	//	}
	//}
	nums := []int{1, 2, 1, 2}
	fmt.Printf("%#v ,%p\n", nums, nums)
	nums1 := ArrayUnique(nums)
	fmt.Printf("%#v ,%p", nums1, nums1)

}

func TestIsEmail(t *testing.T) {
	email1 := "8888@qq.com"
	fmt.Println(IsEmail(email1))
	email2 := "shir?dong@qq.com"
	fmt.Println(IsEmail(email2))
}

func TestRand1(t *testing.T) {
	nums := GetRandNums(1, 100, 1000)
	fmt.Println(nums)
}

type User struct {
	Name string `-`
	Age  uint8
}

func TestStructToMap(t *testing.T) {
	user := User{
		Name: "demo",
		Age:  18,
	}
	toMap := StructToMap(&user)
	fmt.Println(toMap)
}
