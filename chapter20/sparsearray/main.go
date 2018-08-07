package main

import (
	"bufio"
	"fmt"
	"os"
)

type ValNode struct {
	row, col, val int
}

func main() {
	//1,先创建原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //白子

	//2,输出看看原始的数组
	for _, v := range chessMap {
		//一维数组
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	//3,转成稀疏数组
	//思路
	//(1) 遍历二维数组,如果元素不为0,创建一个node结构体
	//(2) 将其翻入到对应的切片

	var sparseArr []ValNode

	//标准的一个稀疏数组,还有一个 记录元素的二维数组的规模(行和列,默认值)
	//创建一个ValNode值节点
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}

	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	//输出稀疏数组
	fmt.Println("当前的稀疏数组是:::::")
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}

	//将这个稀疏数组，存盘 d:/chessmap.data
	file, err := os.Create("chessmap.data")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, valNode := range sparseArr {
		line := fmt.Sprintf("%d\t%d\t%d", valNode.row, valNode.col, valNode.val)
		_, err := fmt.Fprintln(writer, line)
		if err != nil {
			fmt.Println("err=", err)
			return
		}
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	//如何恢复原始的数组

	//1. 打开这个d:/chessmap.data => 恢复原始数组.

	//2. 这里使用稀疏数组恢复

	// 先创建一个原始数组

	openFile, err := os.Open("chessmap.data")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer openFile.Close()

	var newSparseArr []ValNode

	scanner := bufio.NewScanner(openFile)
	scanner.Split(bufio.ScanLines)
	/*
		ScanLines (默认)
		ScanWords
		ScanRunes (遍历UTF-8字符非常有用)
		ScanBytes
	*/

	//是否有下一行
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}
		var row, col, val int
		_, err = fmt.Sscanf(line, "%d\t%d\t%d", &row, &col, &val)
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		newSparseArr = append(newSparseArr, ValNode{
			row: row,
			col: col,
			val: val,
		})
	}
	fmt.Println("newSparseArr=", newSparseArr)

	var chessMap2 [11][11]int

	// 遍历 sparseArr [遍历文件每一行]
	//for i, valNode := range sparseArr {
	//	if i != 0 { //跳过第一行记录值
	//		chessMap2[valNode.row][valNode.col] = valNode.val
	//	}
	//}
	for i, valNode := range newSparseArr {
		if i != 0 { //跳过第一行记录值
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}

	// 看看chessMap2 是不是恢复.
	fmt.Println("恢复后的原始数据......")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
