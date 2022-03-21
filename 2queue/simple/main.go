package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用结构体管理队列
type Queue struct {
	maxSize int    // 队列的最大容量
	array   [5]int // 数组模拟队列
	front   int    // 表示指向队首
	rear    int    // 表示指向队尾
}

// 添加到队列
func (q *Queue) AddQueue(val int) (err error) {
	// 先判断队列是否已满
	if q.rear == q.maxSize-1 {
		// rear是队列的尾部（含最后一个元素）
		return errors.New("queue full")
	}

	// 将队尾往后移动
	q.rear++
	// 放入到数组中
	q.array[q.rear] = val

	return
}

// 从队列去获取数据
func (q *Queue) GetQueue() (val int, err error) {

	// 先判断队列是否为空
	// 当队尾等于队首时，队列已空
	if q.rear == q.front {
		return -1, errors.New("queue empty")
	}

	q.front++
	val = q.array[q.front]
	return val, err
}

// 显示队列
// 找到队首 然后遍历到队尾
func (q *Queue) ShowQueue() {
	// q.front是不包含队首的
	// q.rear是包含队尾的
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("arrar[%d]=%d\t", i, q.array[i])
	}
	fmt.Println()
}

func main() {
	// 创建一个队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
	var choice string
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示退出")

		_, _ = fmt.Scanln(&choice)
		switch choice {
		case "add":
			var val int
			fmt.Println("请输入需要添加到队列的数据：")
			_, _ = fmt.Scanln(&val)
			if err := queue.AddQueue(val); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("添加成功")
			}
		case "get":
			if val, err := queue.GetQueue(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("取出的数据是:%d\n", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			fmt.Println("程序退出")
			os.Exit(0)
		}
	}
}
