package algorithm

func Swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func QSort(nums []int) []int {
	l, r := 0, len(nums)
	if r == 0 {
		return nums
	}
	qsortRec(nums, l, r-1)
	return nums
}

func qsortRec(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := partitionV2(nums, l, r)
	qsortRec(nums, l, p-1)
	qsortRec(nums, p+1, r)
}

// from both sides to center
func partitionV1(nums []int, l, r int) int {
	p, r := r, r-1
	for l < r {
		if nums[l] <= nums[p] {
			l++
		} else {
			Swap(nums, l, r)
			r--
		}
	}
	if nums[l] <= nums[p] {
		Swap(nums, l+1, p)
	} else {
		Swap(nums, l, p)
	}
	return l
}

// from side to side
func partitionV2(nums []int, l, r int) int {
	p, i := r, l
	// nums[i] will always be the value that greater than nums[p]
	for j := l; j < p; j++ {
		if nums[j] <= nums[p] {
			Swap(nums, i, j)
			i++
		}
	}
	Swap(nums, i, p)
	return i
}
