package utils

import "sync"

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/23 11:41
  @describe :
*/

// Queue 队列
type Queue struct {
	mute sync.Mutex
	list []interface{}
}

// Push 推入信息
func (q *Queue) Push(data interface{}) {
	q.mute.Lock()
	defer q.mute.Unlock()
	q.list = append(q.list, data)
}

// Pop 推出信息
func (q *Queue) Pop() interface{} {
	if len(q.list) == 0 {
		return nil
	}
	q.mute.Lock()
	defer q.mute.Unlock()
	item := q.list[0]
	q.list = q.list[1:]
	return item
}

// BPop 以阻塞的形式推出信息
func (q *Queue) BPop() interface{} {
	q.mute.Lock()
	defer q.mute.Unlock()

	// 阻塞,直到列表长度不为0
	for len(q.list) == 0 {
	}

	item := q.list[0]
	q.list = q.list[1:]
	return item
}

// InitFuncQueue 初始方法队列
var InitFuncQueue = &funcQueue{queue: Queue{}}

// funcQueue 方法队列
type funcQueue struct {
	queue Queue
}

// Push 推入一个方法
func (q *funcQueue) Push(fn func()) {
	q.queue.Push(fn)
}

// Run 执行已存在的方法
func (q *funcQueue) Run() {
	for fn, ok := q.queue.Pop().(func()); ok; fn, ok = q.queue.Pop().(func()) {
		fn()
	}
}
