
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) { // this method belongs to Minstack
	this.stack = append(this.stack, val)
	
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
	} else {
		topMinStack := this.minStack[len(this.minStack)-1]
	
		if topMinStack >= val {
			this.minStack = append(this.minStack, val)
		} else if topMinStack < val {
			this.minStack = append(this.minStack, topMinStack)
		}
	}
	
	
	
	// {
	// 	// topMinStack := this.minStack[len(this.minStack)-1]
	// if topMinStack >= val { // ex. topMinStack = 5, val = 1
	// 	this.minStack[len(this.minStack)-1] = val
	// } else {  // ex. topMinStack = 1, val = 5
	// 	this.minStack = append(this.minStack, topMinStack)
	// }
	// }
	
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	topValue := this.stack[len(this.stack)-1]
	return topValue
}

func (this *MinStack) GetMin() int {
	min := this.minStack[len(this.minStack)-1]
	return min
}


// Design a stack class that supports the push, pop, top, and getMin operations.

// MinStack() initializes the stack object.
// void push(int val) pushes the element val onto the stack.
// void pop() removes the element on the top of the stack.
// int top() gets the top element of the stack.
// int getMin() retrieves the minimum element in the stack.
// Each function should run in 
// O
// (
// 1
// )
// O(1) time.