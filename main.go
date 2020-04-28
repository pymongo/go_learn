package main

import (
	"fmt"
)

const (
	gender_male = iota
	gender_femail = iota
)

const (
	trade_type_buy = iota
	trade_type_sell
)

func main() {
	gender := gender_femail
	trade_type := trade_type_sell
	fmt.Println(gender) // 1
	fmt.Println(trade_type) // 1
}

func add(x int32, y int32) int32 {
	return x + y
}

