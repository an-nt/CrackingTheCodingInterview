package sortingandsearching

// HeapSort dùng cấu trúc Max Heap để sắp xếp 1 dãy tăng dần
func HeapSort(arr []int) {
	right := len(arr) - 1
	if right < 0 {
		return
	}
	for ; right > 0; right-- {
		// tạo cấu trúc Max Heap
		for left := right / 2; left >= 0; left-- {
			heapSort(arr, left, right)
		}
		// phần tử đầu tiên của Max Heap luôn lớn nhất -> chuyển phẩn tử này xuống cuối mảng
		arr[0], arr[right] = arr[right], arr[0]
	}
}

func heapSort(arr []int, left int, right int) {
	if (2*left + 2) <= right {
		// Tồn tại 2 cặp phần tử liên đới
		if arr[left] < arr[2*left+1] || arr[left] < arr[2*left+2] {
			// điều chỉnh vị trí để thỏa mãn Max Heap
			if arr[2*left+1] >= arr[2*left+2] {
				arr[left], arr[2*left+1] = arr[2*left+1], arr[2*left]
				heapSort(arr, 2*left+1, right)
			} else {
				arr[left], arr[2*left+2] = arr[2*left+2], arr[2*left]
				heapSort(arr, 2*left+2, right)
			}
		}
		return
	}
	if (2*left + 1) <= right {
		// Tồn tại 1 cặp phần tử liên đới
		if arr[left] < arr[2*left+1] {
			// điều chỉnh vị trí để thỏa mãn Max Heap
			arr[left], arr[2*left+1] = arr[2*left+1], arr[left]
		}
		return
	}
	// Không tồn tại cặp phần tử liên đới
	return
}
