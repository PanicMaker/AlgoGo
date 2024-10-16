package Stack

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/basic-calculator/description/

func calculate(s string) int {
	// 由于第一个数可能是负数，为了减少边界判断。一个小技巧是先往 nums 添加一个 0
	nums := make([]int, 1)
	ops := make([]byte, 0)

	cal := func() {
		num1 := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		num2 := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]
		if op == '+' {
			nums = append(nums, num2+num1)
		} else {
			nums = append(nums, num2-num1)
		}
	}

	s = strings.ReplaceAll(s, " ", "")

	for i := 0; i < len(s); {
		ch := s[i]
		switch ch {
		case ' ':
			i++
		case '(':
			ops = append(ops, ch)
			i++
		case ')':
			for ops[len(ops)-1] != '(' {
				cal()
			}
			ops = ops[:len(ops)-1]
			i++
		case '+', '-':
			// 为防止 () 内出现的首个字符为运算符，将所有的空格去掉，并将 (- 替换为 (0-，(+ 替换为 (0+
			if i > 0 && (s[i-1] == '(' || s[i-1] == '+' || s[i-1] == '-') {
				nums = append(nums, 0)
			}
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				cal()
			}
			ops = append(ops, ch)
			i++
		default:
			num := 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			nums = append(nums, num)
		}
	}

	for len(ops) > 0 {
		cal()
	}

	return nums[len(nums)-1]
}

// 转换为逆波兰表达式求值
func calculate1(s string) int {
	strToArr := func(s string) (res []string) {

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

	convert := func(strs []string) (res []string) {
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
				stack2 = append(stack2, v)
			}
		}

		for len(stack1) > 0 {
			stack2 = append(stack2, stack1[len(stack1)-1])
			stack1 = stack1[:len(stack1)-1]
		}

		return stack2
	}

	eval := func(tokens []string) int {
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

	return eval(convert(strToArr(s)))
}
