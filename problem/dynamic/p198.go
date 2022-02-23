package main

func main() {

}

func rob(nums []int) int {
	var robMap map[int]int
	robMap = make(map[int]int)
	robMap[0] = nums[0]
	if len(nums) > 1 {
		if nums[1] > nums[0] {
			robMap[1] = nums[1]
		} else {
			robMap[1] = nums[0]
		}
	}

	if len(nums) > 2 {
		for i := 2; i < len(nums); i++ {
			if robMap[i-2]+nums[i] > robMap[i-1] {
				robMap[i] = robMap[i-2] + nums[i]
			} else {
				robMap[i] = robMap[i-1]
			}
		}
	}

	return robMap[len(nums)-1]

}
