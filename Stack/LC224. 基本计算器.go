package Stack

import "strconv"

// https://leetcode.cn/problems/basic-calculator/description/

func calculate(s string) int {
	return eval(convert(strToArr(s)))
}

func strToArr(s string) (res []string) {

	for i := 0; i < len(s); {
		v := s[i]
		if v == '+' || v == '-' || v == '*' || v == '/' || v == '(' || v == ')' {
			res = append(res, string(v))
			i++
		} else if v >= '0' && v <= '9' {
			num := ""
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num += string(s[i])
				i++
			}
			res = append(res, num)
		} else {
			i++
		}
	}

	return
}

func convert(strs []string) (res []string) {
	priority := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	stack1, stack2 := make([]string, 0), make([]string, 0)

	for _, v := range strs {
		if v == "(" {
			stack1 = append(stack1, v)
		} else if v == ")" {
			for stack1[len(stack1)-1] != "(" {
				stack2 = append(stack2, stack1[len(stack1)-1])
				stack1 = stack1[:len(stack1)-1]
			}
			stack1 = stack1[:len(stack1)-1]
		} else if v == "+" || v == "-" || v == "*" || v == "/" {
			for len(stack1) > 0 && priority[v] <= priority[stack1[len(stack1)-1]] {
				stack2 = append(stack2, stack1[len(stack1)-1])
				stack1 = stack1[:len(stack1)-1]
			}
			stack1 = append(stack1, v)
		} else {
			stack1 = append(stack1, v)
		}
	}

	for len(stack1) > 0 {
		stack2 = append(stack2, stack1[len(stack1)-1])
		stack1 = stack1[:len(stack1)-1]
	}

	return stack2
}

func eval(tokens []string) int {
	stack := make([]int, 0)

	for _, v := range tokens {
		switch v {
		case "+":
			d1 := stack[len(stack)-1]
			d2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := d1 + d2
			stack = append(stack, res)
		case "-":
			d1 := stack[len(stack)-1]
			d2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := d2 - d1
			stack = append(stack, res)
		case "*":
			d1 := stack[len(stack)-1]
			d2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := d1 * d2
			stack = append(stack, res)
		case "/":
			d1 := stack[len(stack)-1]
			d2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := d2 / d1
			stack = append(stack, res)
		default:
			d, _ := strconv.Atoi(v)
			stack = append(stack, d)
		}
	}

	return stack[len(stack)-1]
}
