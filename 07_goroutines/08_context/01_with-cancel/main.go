package main

func main() {
	/*
		context 用于在多个goroutine之间传递取消信号、超时控制和请求相关的上下文信息
		是Go语言并发编程中的一个关键组件，能够有效地管理不同任务之间的协作和资源释放
		context，即上下文，在Go语言中是一个接口，定义了四个方法：CancelFunc, Deadline, Done, 和 Err
		主要用于在不同的goroutine之间传递取消信号和上下文信息
		核心功能
			1. 取消（cancel）
				上游调用 cancel()，所有持有这个 context 的 goroutine 都会收到信号，自己检查后退出
			2. 超时（timeout）
				自动在指定时间后发出取消信号
			3. 截止时间（deadline）
				和超时类似，只是指定一个绝对时间点
			4, 传值（value）
				可以在一个请求范围内传递一些公共数据（比如 request id），但不建议用来传业务参数
		context可以通过context.Background()和context.WithCancel等方法创建
	*/
	/*
		1. 根 context：所有context都应从context.Background()开始
			这是整个上下文树的根节点，常用在 main、初始化、顶层，永远不会被取消
			context.TODO() 当还不确定用哪种 context 时先占位

	*/
	//ctx := context.Background()
	//ctx := context.TODO()

	/*
		2. 可取消的 Context
			使用 context.WithCancel(parent)
			返回一个新的子 Context 和一个 cancel 函数
			调用 cancel() 后，所有持有该 context 的 goroutine 都会收到 <-ctx.Done() 信号
	*/
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel

	/*
		3. 带超时 / 截止时间的 Context
			context.WithTimeout(parent, duration)
			自动在一段时间后触发取消（等价于 WithDeadline + time.Now().Add(duration))
	*/
	// 使用Deadline
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	//defer cancel()

	// 使用Timeout
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	/*
		4. 带值的 Context
			context.WithValue(parent, key, value)
			用来在调用链上传递少量、和请求范围相关的数据（例如 traceID、用户ID）
			注意：不推荐放业务数据或大对象，只放元信息
	*/
	//ctx = context.WithValue(context.Background(), "requestID", "12345")
}
