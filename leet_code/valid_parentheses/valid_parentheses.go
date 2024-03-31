package valid_parentheses

func IsValid(s string) bool {
	stack := []rune{}
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			sum := string(last) + string(v)
			if check(sum) {
				continue
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

func check(s string) (isValid bool) {
	if s == "{}" || s == "()" || s == "[]" {
		isValid = true
	}
	return isValid
}
