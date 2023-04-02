package chinese2digits

import (
	"fmt"
	"testing"
)

func TestTakeChineseNumberFromString(t *testing.T) {
	s := "二三款"
	r := TakeChineseNumberFromString(s)
	fmt.Println(r)

	s = "1点六排量款"
	r = TakeChineseNumberFromString(s)
	fmt.Println(r)

	s = "20w左右的车"
	r = TakeChineseNumberFromString(s)
	fmt.Println(r)

	s = "二十万一左右二零年的车的车"
	r = TakeChineseNumberFromString(s)
	fmt.Println(r)
}
