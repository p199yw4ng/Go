package alog

// 冒泡排序
func bubbleSort(arr []int) []int {
	n := len(arr)
	// 从第一个元素开始
	// 比较相邻的两个元素
	// 如果前一个元素大于后一个元素，则交换它们的位置
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}


// 选择排序
func selectionSort(arr []int) []int {
	n := len(arr)
	// 从第一个元素开始
	// 在剩余的元素中找到最小的元素
	// 将最小的元素与当前元素交换

	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

// 插入排序 
func insertionSort(arr []int) []int {
	n := len(arr)
	// 从第二个元素开始
	// 将当前元素插入到前面已经有序的部分中
	// 直到所有元素都插入到前面已经有序的部分中
	
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}



// 归并排序
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// 分割数组为左右两部分
	// 递归地对左右两部分进行排序
	// 合并左右两部分
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}	

// 快速排序
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// 选择第一个元素作为基准
	pivot := arr[0]
	// 分割数组为左右两部分
	// 小于基准的元素放在左边，大于基准的元素放在右边
	// 递归地对左右两部分进行排序
	left, right := make([]int, 0), make([]int, 0)
	for _, num := range arr[1:] {
		if num <= pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}	
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, pivot), right...)
}

// 堆排序
func heapSort(arr []int) []int {
	n := len(arr)
	// 构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	// 依次将堆顶元素与最后一个元素交换
	// 然后重新构建最大堆
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
	return arr
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest!= i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

