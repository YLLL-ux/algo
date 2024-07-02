package main

func main() {

}

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	// 将元素加入inStack
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
	popNum := this.Peek()
	this.outStack = this.outStack[:len(this.outStack)-1]
	return popNum
}

func (this *MyQueue) Peek() int {
	// 判断队列是否为空
	if this.Empty() {
		return -1
	}

	if len(this.outStack) == 0 {
		// 元素从inStack移出，然后移入outStack
		length := len(this.inStack)
		for i := length - 1; i >= 0; i-- {
			this.outStack = append(this.outStack, this.inStack[i])
		}
		// 移出完毕后置空
		this.inStack = []int{}
	}

	return this.outStack[len(this.outStack)-1]
}

func (this *MyQueue) Empty() bool {
	if len(this.inStack) == 0 && len(this.outStack) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
