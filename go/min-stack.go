package leetcode

// array implementation of stack data structure
type MinStack struct {
	stack      []int
	minElement int
}

func Constructor() MinStack {
	s := MinStack{stack: []int{}}
	return s
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.minElement = val
		this.stack = append(this.stack, val)
	} else if val < this.minElement {
		this.stack = append(this.stack, 2*val-this.minElement)
		this.minElement = val
	} else {
		this.stack = append(this.stack, val)
	}
}

func (this *MinStack) Pop() {
	if this.stack[len(this.stack)-1] < this.minElement {
		this.minElement = 2*this.minElement - this.stack[len(this.stack)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if this.stack[len(this.stack)-1] < this.minElement {
		return this.minElement
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minElement
}
