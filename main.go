package bubble_sort

func BubbleSort(ii []int) []int {
	for i := 0; i < len(ii); i++ {
		for j := 0; j < len(ii)-i-1; j++ {
			if ii[j] > ii[j+1] {
				ii[j],ii[j+1] = ii[j+1],ii[j]
			}
		}
	}
	return ii
}


func SelectionSort(ii []int) []int {
	for i := 0; i < len(ii); i++ {
		minI := i
		for j := minI; j < len(ii)-1; j++ {
			if ii[j] < ii[minI] {
				minI = j
			}
		}
		if minI != i {
			ii[minI],ii[i] = ii[i],ii[minI]
		}
	}
	return ii
}
