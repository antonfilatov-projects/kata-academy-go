package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		fmt.Println(calculator(text))
	}
}

func calculator(str string) string {

	splitedStr := strings.Split(str, " ")

	if len(splitedStr) != 3 {
		panic("допустимо только два операнда")
	}

	var operand1isArabic bool
	var operand2isArabic bool
	var operand1 int
	var operand2 int
	var result string
	operator := splitedStr[1]

	if oper1, err := strconv.Atoi(splitedStr[0]); err == nil {
		operand1isArabic = true
		operand1 = oper1
	} else {
		operand1isArabic = false
		operand1 = toArabic(splitedStr[0])
	}

	if oper2, err := strconv.Atoi(splitedStr[2]); err == nil {
		operand2isArabic = true
		operand2 = oper2
	} else {
		operand2isArabic = false
		operand2 = toArabic(splitedStr[2])
	}

	if !operand1isArabic && operand2isArabic || operand1isArabic && !operand2isArabic {
		panic("оба операнда должны быть в одной системе счисления")
	}

	if operand1isArabic && operand2isArabic {
		result = calculation(operand1, operand2, operator)
	} else {
		result = calculation(operand1, operand2, operator)

		if testInt, err := strconv.Atoi(result); err == nil {
			if testInt <= 0 {
				panic("в римской системе счисления нет отрицательных чисел и ноля")
			}
			result = toRoman(testInt)
		}
	}
	return result
}

func calculation(operand1 int, operand2 int, operator string) string {
	if operand1 <= 0 || operand1 > 10 {
		panic("числа вне допустимого диапазона")
	}
	if operand2 <= 0 || operand2 > 10 {
		panic("числа вне допустимого диапазона")
	}

	var result int

	switch operator {
	case "+":
		result = operand1 + operand2
		return strconv.Itoa(result)
	case "-":
		result = operand1 - operand2
		return strconv.Itoa(result)
	case "/":
		result = operand1 / operand2
		result = int(result)
		return strconv.Itoa(result)
	case "*":
		result = operand1 * operand2
		return strconv.Itoa(result)
	default:
		panic("недопустимая арифметическая операция")
	}

}

func getArabic(roman string) int {
	switch roman {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}

func toArabic(roman string) int {
	roman = strings.ToUpper(roman)
	var result int
	var arabic int
	var end int
	numArr := strings.Split(roman, "")
	end = len(numArr) - 1

	result = getArabic(numArr[end])

	for i := end - 1; i >= 0; i-- {
		arabic = getArabic(numArr[i])

		if arabic < getArabic(numArr[i+1]) {
			result -= arabic
		} else {
			result += arabic
		}
	}
	return result
	//return strconv.Itoa(result)
}

func toRoman(arabic int) string {

	var result string
	if arabic >= 1000 {
		result = "M" + toRoman(arabic-1000)
		return result
	}

	if arabic >= 500 {
		if arabic >= 900 {
			result = "CM" + toRoman(arabic-900)
		} else {
			result = "D" + toRoman(arabic-500)
		}
		return result
	}

	if arabic >= 100 {
		if arabic >= 400 {
			result = "CD" + toRoman(arabic-400)
		} else {
			result = "C" + toRoman(arabic-100)
		}
		return result
	}

	if arabic >= 50 {
		if arabic >= 90 {
			result = "XC" + toRoman(arabic-90)
		} else {
			result = "L" + toRoman(arabic-50)
		}
		return result
	}

	if arabic >= 10 {
		if arabic >= 40 {
			result = "XL" + toRoman(arabic-40)
		} else {
			result = "X" + toRoman(arabic-10)
		}
		return result
	}

	if arabic >= 5 {
		if arabic == 9 {
			result = "IX"
		} else {
			result = "V" + toRoman(arabic-5)
		}
		return result
	}

	if arabic > 0 {
		if arabic == 4 {
			result = "IV"
		} else {
			result = "I" + toRoman(arabic-1)
		}
		return result
	}

	return ""
}
