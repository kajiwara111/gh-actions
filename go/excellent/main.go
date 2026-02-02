package main

// 偶数か奇数かを判定する関数
func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}
