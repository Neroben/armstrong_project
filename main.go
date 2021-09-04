package main
import (
	"fmt"
	"math"
	"os"
)

func main() {
	number := input()
	fmt.Printf("isArmstrong = %t\n", isArmstrongNumber(number))
}

func input() int {
	fmt.Printf("Введите число для проверки: ")
	var number int
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		os.Exit(404)
	}
	return number
}

func isArmstrongNumber(number int) bool {
	result := 0
	checkNumber := number
	for result < number && checkNumber != 0 {
		result += int(math.Pow(float64(checkNumber % 10), 3))
		checkNumber /= 10
	}
	return result == number
}
