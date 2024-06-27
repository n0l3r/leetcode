func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	x, y := len(nums1), len(nums2)
	low, high := 0, x
	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (x+y+1)/2 - partitionX
		maxLeftX, minRightX := math.MinInt32, math.MaxInt32
		maxLeftY, minRightY := math.MinInt32, math.MaxInt32
		if partitionX > 0 {
			maxLeftX = nums1[partitionX-1]
		}
		if partitionX < x {
			minRightX = nums1[partitionX]
		}
		if partitionY > 0 {
			maxLeftY = nums2[partitionY-1]
		}
		if partitionY < y {
			minRightY = nums2[partitionY]
		}
		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (x+y)%2 == 0 {
				return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2
			}
			return float64(max(maxLeftX, maxLeftY))
		} else if maxLeftX > minRightY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}
	return 0
}