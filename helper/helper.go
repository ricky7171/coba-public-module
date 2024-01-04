package helper

func Factorial(num int) int {
	res := 1
	for i := 1; i <= num; i++ {
		res *= i
	}
	return res
}

func SpecialAddition(num1, num2 int) int {
	if num2%2 == 0 {
		return num1 + num2 + 7
	} else {
		return num1 + num2
	}

}
