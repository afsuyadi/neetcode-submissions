
type Stack struct {
	stack []byte
}

func Constructor() Stack {
	return Stack{
		stack: []byte{},
	}
}

func (this *Stack) Push(val byte) {
	this.stack = append(this.stack, val)
}

func (this *Stack) Seek() byte {
	return this.stack[len(this.stack)-1]
}

func (this *Stack) Pop() byte {
	popped := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return popped
}

func (this *Stack) Len() int {
	return len(this.stack)
}

func isValid(s string) bool {
	// all closed parantheses will always summed up to even numbers
	// half-left of it will always be the open paranthesis
	// let's try to check until half of it, similar to pallindrome	

	// v1.0 - Using Hashmap
	truth := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']', 
	}
	// stack := make([]byte, 0)
	// for i := 0; i < len(s); i++ {
	// 	if len(stack) == 0 { // if the stack is empty, pushed the current value into that
	// 		stack = append(stack, s[i])			
	// 		continue
	// 	} // ex. stack = ["(", "}"]
	// 	top := stack[len(stack)-1] // pop the furthest value
	// 	couple_top := truth[top] // couple_top = "{"
	// 	if s[i] != couple_top { // ex. s[i] = "[", if the values are NOT coupled parantheses
	// 		stack = append(stack, s[i])
	// 	} else { // if the iteration and stack value matches as coupled parantheses
	// 		stack = stack[:len(stack)-1] // pop the value out
	// 	}
	// 	if len(stack) == 0 && i == len(s)-1 {
	// 		return true
	// 	}
	// }
	// return false

	// v2.0 - Using Stack
	// ex. input = ()[]{}
	store := Constructor()
	for i := 0; i < len(s); i++ {
		if store.Len() == 0 {
			store.Push(s[i]) // ex. "("
		} else {
			seek := store.Seek()
			if truth[seek] == s[i] { // which means that they are a couple
				store.Pop()
			} else { // which means that they are not a couple, so just store it into the stack
				store.Push(s[i])
			}
		}
	}
	if store.Len() != 0 {
		return false
	} else {
		return true
	}

}
