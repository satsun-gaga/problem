package main

import (
	"fmt"
	"strconv"
)

func test2(a float64) string {
	b := strconv.FormatFloat(a*1000, 'f', 0, 64)

	end := []string{"仟", "佰", "拾", "亿", "仟", "佰", "拾", "万", "仟", "佰", "拾", "元", "角", "分","厘"}
	midd := map[string]string{"0": "零", "1": "壹", "2": "贰", "3": "叁", "4": "肆", "5": "伍", "6": "陆", "7": "柒", "8": "捌", "9": "玖"}

	end1 := end[len(end)-len(b) :]

	str := ""
	for k, v := range b[:] {
		str = str + midd[string(v)] + end1[k]
	}

	return str
}

func main() {
	fmt.Println(test2(123456789.123))
}


