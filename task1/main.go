package main

import (
	"fmt"
	"strconv"
)

func main() {
	var price int
	fmt.Print("Input price : ")
	fmt.Scanln(&price)

	for k, v := range MoneyCount(price) {
		fmt.Println("Rp", k, ":", v)
	}
}

func MoneyCount(price int) map[string]int {
	var moneys = [...]int{
		100000,
		50000,
		20000,
		10000,
		5000,
		2000,
		1000,
		500,
		200,
		100,
	}

	result := make(map[string]int)
	var total int
	var tmpCount = price

	for _, money := range moneys {
		total = tmpCount / money
		if total != 0 {
			result[strconv.Itoa(money)] = total
		}
		tmpCount = tmpCount % money

		if tmpCount > 0 && tmpCount < 100 {
			result[strconv.Itoa(100)]++
			break
		}
	}
	return result

}

// func Word(input1, input2 string) bool {

// }
