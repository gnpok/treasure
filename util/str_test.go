package util

import (
	"fmt"
	"testing"
)

func TestReversal(t *testing.T) {
	str := "123456789"
	after := Reversal(str)
	fmt.Println(after)
}
