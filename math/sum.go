package baseMath

func Add(num1, num2 int) int {
	return num1 + num2
}

func Minus(num1, num2 int) int {
	return num1 - num2
}

func Max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func Min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func BadLuck(num1, num2 int) int {
	return 13
}
