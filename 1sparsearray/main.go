package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

/*

需求：将五子棋盘保存和恢复。

现有10x10的棋盘，有若干黑子（用1表示）和白子（用2表示）。黑白子分布如下图所示：

0 1 0 0 0 0 0 0 0 0
0 1 0 0 0 0 0 2 0 0
0 0 0 0 0 0 0 0 2 0
0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0

如上图分析可得：
黑子有：[0][1], [1][1]
白子有：[1][7], [2][8]

原始实现方式：
定义一个长度为10x10、byte类型的二维数组，默认值为0。
根据图中所示，将黑白子处于二维数组的索引处依次赋值1或2。
将二维数组保存至文件中。

使用稀疏数组实现方式：
再创建一个新的二维数组，存储棋盘上有效的数据。
定义结构体值节点：VNode{ Row, Cell, Value}
标准的稀疏数组，首个元素记录，原始二维数组的大小以及默认值。
后面存入有效的数据。

*/

type vNode struct {
	Row   byte
	Cell  byte
	Value byte
}

func (v *vNode) String() string {
	return fmt.Sprintf("%d %d %d", v.Row, v.Cell, v.Value)
}

func main() {

	// 定义一个长度为10x10、byte类型的二维数组
	var arr [10][10]byte

	// 依次给棋盘赋值

	// 	黑子有：[0][1], [1][1]
	arr[0][1] = 1
	arr[1][1] = 1

	// 白子有：[1][7], [2][8]
	arr[1][7] = 2
	arr[2][8] = 2

	fmt.Println("================================赋值的棋盘===============================")
	printArr(arr)

	var vNodes []vNode

	// 存入原始数组规模以及默认值
	vNodes = append(vNodes, vNode{
		Row:   10,
		Cell:  10,
		Value: byte(0),
	})

	// 遍历原始的二维数组，筛选出有效的数据
	for i, r := range arr {
		for j, c := range r {
			if c != 0 {
				vNode := vNode{
					Row:   byte(i),
					Cell:  byte(j),
					Value: c,
				}
				vNodes = append(vNodes, vNode)
			}
		}
	}
	fmt.Println("================================值节点如下===============================")
	for _, vNode := range vNodes {
		fmt.Println(vNode)
	}

	// 将稀疏数组存入到文件
	f, _ := os.OpenFile("chess.data", os.O_CREATE|os.O_TRUNC, 0600)

	for _, vNode := range vNodes {
		f.WriteString(vNode.String())
		f.WriteString("\n")
	}

	_ = f.Close()

	// 从文件读取并且恢复到棋盘上
	var arr2 [10][10]byte
	data, _ := ioutil.ReadFile("chess.data")
	dataArr := strings.Split(string(data), "\n")
	for i, r := range dataArr {
		// 跳过首个元素
		if i > 0 && r != "" {
			// row cell value
			dataRow := strings.Split(string(r), " ")
			row, _ := strconv.Atoi(dataRow[0])
			cell, _ := strconv.Atoi(dataRow[1])
			value, _ := strconv.Atoi(dataRow[2])
			arr2[row][cell] = byte(value)
		}
	}

	fmt.Println("================================恢复的棋盘===============================")
	printArr(arr)

	_ = os.Remove("chess.data")
}

func printArr(arr [10][10]byte) {
	for _, r := range arr {
		for _, c := range r {
			fmt.Printf("%d\t", c)
		}
		fmt.Println()
	}
}
