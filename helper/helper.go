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

func SpecialSubstraction(num1, num2 int) int {
	res := 0
	if num2%2 == 0 {
		res = num1 - num2 - 7
	} else {
		res = num1 - num2
	}
	if res < 0 {
		return 0
	} else {
		return res + 1
	}

}
