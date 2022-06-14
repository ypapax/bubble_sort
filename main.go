package bubble_sort

func Sort(ii []int) []int {
	for i := 0; i < len(ii); i++ {
		for j := 0; j < len(ii)-i-1; j++ {
			if ii[j] > ii[j+1] {
				ii[j],ii[j+1] = ii[j+1],ii[j]
			}
		}
	}
	return ii
}
