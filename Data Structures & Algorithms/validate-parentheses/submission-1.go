
func isValid(s string) bool {
	// all closed parantheses will always summed up to even numbers
	// half-left of it will always be the open paranthesis
	// let's try to check until half of it, similar to pallindrome
	truth := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']', 
	}
    // for i := 0; i < len(s)/2; i++ { // ex. len(s) = 6
	// 	o := s[i]
	// 	c := s[len(s) - 1 - i]
	// 	if o != '(' && o != '[' && o != '{' {
	// 		return false
	// 	}
	// 	if c != ')' && c != ']' && c != '}' {
	// 		return false
	// 	}
	// 	if truth[o] == c {
	// 		continue
	// 	} else {
	// 		return false
	// 	}

	// }
	// return true
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 { // if the stack is empty, pushed the current value into that
			stack = append(stack, s[i])			
			continue
		} // ex. stack = ["(", "}"]
		top := stack[len(stack)-1] // pop the furthest value
		couple_top := truth[top] // couple_top = "{"
		if s[i] != couple_top { // ex. s[i] = "[", if the values are NOT coupled parantheses
			stack = append(stack, s[i])
		} else { // if the iteration and stack value matches as coupled parantheses
			stack = stack[:len(stack)-1] // pop the value out
		}
		if len(stack) == 0 && i == len(s)-1 {
			return true
		}
	}
	return false
}
