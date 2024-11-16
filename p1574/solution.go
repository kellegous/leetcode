package p1574

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	l, r := 0, n-1

	// find the non-decreasing prefix
	for l < r && arr[l] <= arr[l+1] {
		l++
	}

	if l == r {
		return 0
	}

	// find the non-decreasing suffix
	for r > l && arr[r] >= arr[r-1] {
		r--
	}

	ix := min(n-l-1, r)
	for i, j := 0, r; i <= l && j < n; {
		if arr[i] <= arr[j] {
			ix = min(ix, j-i-1)
			i++
		} else {
			j++
		}
	}
	return ix
}
