package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("ծրագիրը գտնում է թվի պարզ արտադրիչները")
	fmt.Println(" մուտքագրե՛ք թիվը")
	n, err := fmt.Scan(&num)
	if n < 1 {
		fmt.Println("մուտքային տվյալների սխալ", err)
		return
	}
	if !(num > 0) {
		fmt.Println("մուտքային տվյալների սխալ,պահանջվում է բնական թիվ")
		return
	}
	r := nextPrime(1)

	for ; r <= num; r = nextPrime(r) {
		for ; num%r == 0; num /= r {
			fmt.Println(r)

		}
	}

}

func isPrime(num int) bool {
	k := 2
	if num == 1 {
		return false
	}
	for ; k <= num/2; k++ {
		if num%k == 0 {
			return false
		}
	}

	return true
}

func nextPrime(num int) int {
	for num = num + 1; isPrime(num) == false; num++ {

	}
	return num

}
