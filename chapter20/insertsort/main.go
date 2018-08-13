package main

import "fmt"

func InsertSort(a *[]int) {
	//完成第一次,给第二个元素找到合适的位置并插入

	arr := *a

	for j := 1; j < len(arr); j++ {
		insertVal := arr[j]
		insertIndex := j - 1

		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] //数据后移
			insertIndex--
		}

		//插入
		if insertIndex+1 != j {
			arr[insertIndex+1] = insertVal
		}

		fmt.Printf("第%d次插入后arr=%v \n", j, arr)
	}

}

func main() {
	arr := []int{23, 0, 12, 56, 34}
	fmt.Printf("原始数组arr=%v \n", arr)
	InsertSort(&arr)

	fmt.Printf("排序后数组arr=%v \n", arr)
}
