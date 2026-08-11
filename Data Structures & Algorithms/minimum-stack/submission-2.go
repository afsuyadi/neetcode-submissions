
// type MinStack struct {
// 	stack    []int
// 	minStack []int
// }

// func Constructor() MinStack {
// 	return MinStack{
// 		stack:    []int{},
// 		minStack: []int{},
// 	}
// }

// func (this *MinStack) Push(val int) { // this method belongs to Minstack
// 	this.stack = append(this.stack, val)
	
// 	if len(this.minStack) == 0 {
// 		this.minStack = append(this.minStack, val)
// 	} else {
// 		topMinStack := this.minStack[len(this.minStack)-1]
	
// 		if topMinStack >= val {
// 			this.minStack = append(this.minStack, val)
// 		} else if topMinStack < val {
// 			this.minStack = append(this.minStack, topMinStack)
// 		}
// 	}
	
	
	
// 	// {
// 	// 	// topMinStack := this.minStack[len(this.minStack)-1]
// 	// if topMinStack >= val { // ex. topMinStack = 5, val = 1
// 	// 	this.minStack[len(this.minStack)-1] = val
// 	// } else {  // ex. topMinStack = 1, val = 5
// 	// 	this.minStack = append(this.minStack, topMinStack)
// 	// }
// 	// }
	
// }

// func (this *MinStack) Pop() {
// 	this.stack = this.stack[:len(this.stack)-1]
// 	this.minStack = this.minStack[:len(this.minStack)-1]
// }

// func (this *MinStack) Top() int {
// 	topValue := this.stack[len(this.stack)-1]
// 	return topValue
// }

// func (this *MinStack) GetMin() int {
// 	min := this.minStack[len(this.minStack)-1]
// 	return min
// }


// // Design a stack class that supports the push, pop, top, and getMin operations.

// // MinStack() initializes the stack object.
// // void push(int val) pushes the element val onto the stack.
// // void pop() removes the element on the top of the stack.
// // int top() gets the top element of the stack.
// // int getMin() retrieves the minimum element in the stack.
// // Each function should run in 
// // O
// // (
// // 1
// // )
// // O(1) time.


type Stack[T any] struct {
	stack []T
}

func (this *Stack[T]) Push(val T) {
	this.stack = append(this.stack, val)
}

func (this *Stack[T]) Seek() T {
	return this.stack[len(this.stack)-1]
}

func (this *Stack[T]) Pop() T {
	popped := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return popped
}

func (this *Stack[T]) Len() int {
	return len(this.stack)
}

type MinStack struct {
	stack    Stack[int] // []int
	minStack Stack[int] //[]int
}

func Constructor() MinStack {
	return MinStack{
		stack:    Stack[int]{},
		minStack: Stack[int]{},
	}
}

func (this *MinStack) Push(val int) { // this method belongs to Minstack
	this.stack.Push(val)

	if this.minStack.Len() == 0 {
		this.minStack.Push(val)
	} else {
		topMinStack := this.minStack.Seek()

		if topMinStack >= val {
			this.minStack.Push(val)
		} else {
			this.minStack.Push(topMinStack)
		}
	}
}

func (this *MinStack) Pop() {
	this.stack.Pop()
	this.minStack.Pop()
}

func (this *MinStack) Top() int {
	return this.stack.Seek()
}

func (this *MinStack) GetMin() int {
	return this.minStack.Seek()
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