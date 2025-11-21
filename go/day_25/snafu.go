package day_25

import (
	"fmt"
)

func ConvertToDecimalSum(input *[]*[]byte) (sum int) {
	for _, v := range *input {
		sum += ConvertToDecimal(v)
	}

	return
}

func ConvertToSnafu(sum int) string {
	var result string

	for sum > 0 {
		rem := sum % 5
		sum /= 5
		switch {
		case rem <= 2:
			result = fmt.Sprintf("%d", rem) + result
		case rem == 3:
			result = "=" + result
			sum++
		case rem == 4:
			result = "-" + result
			sum++
		}
	}

	return result
}

func ConvertToDecimal(input *[]byte) int {
	var sum int
	coef := 1
	if (*input)[len(*input)-1] == '\r' || (*input)[len(*input)-1] == '\n' {
		*input = (*input)[:len(*input)-1]
	}

	for i := len(*input) - 1; i >= 0; i-- {
		switch (*input)[i] {
		case '=':
			sum -= 2 * coef
		case '-':
			sum -= coef
		case '0':
			sum += 0 * coef
		case '1':
			sum += coef
		case '2':
			sum += 2 * coef
		}
		coef *= 5
	}
	return sum
}
