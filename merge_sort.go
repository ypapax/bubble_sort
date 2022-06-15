package bubble_sort

import "log"

func MergeSort(a []int) []int {
	if len(a) > 1 {
		center := len(a)/2
		log.Println("center", center, "len(a)", len(a))
		return merge(MergeSort(a[:center]), MergeSort(a[center:]))
	}
	log.Println("returning a")
	return a
}

func merge(a, b []int) []int {
	var result []int
	for {
		if len(a) == 0 || len(b) == 0 {
			break
		}
		var c int
		if a[0] < b[0] {
			c = a[0]
			a=a[1:]
		} else {
			c = b[0]
			b=b[1:]
		}
		result = append(result, c)
	}
	result = append(result, a...)
	result = append(result, b...)
	return result
}
