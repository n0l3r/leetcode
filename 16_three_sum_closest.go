// Source: https://leetcode.com/problems/3sum-closest/
func threeSumClosest(nums []int, target int) int {
	res := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return sum
			}

			if abs(sum-target) < abs(res-target) {
				res = sum
			}
			if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}