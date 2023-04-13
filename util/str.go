package util

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
