package threesum

import "fmt"

func ThreeSum(nums []int) (output [][]int) {
	mapv := map[int]int{}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			for k := 0; k < len(nums); k++ {
				if i != j && j != k && i != k && nums[i]+nums[j]+nums[k] == 0 {
					for k := range mapv {
						delete(mapv, k)
					}
					mapv[nums[i]] = nums[i]
					mapv[nums[j]] = nums[j]
					mapv[nums[k]] = nums[k]
					if !check(output, nums[i], nums[j], nums[k], mapv) {
						output = append(output, []int{nums[i], nums[j], nums[k]})
					}
				}
			}
		}
	}
	return output

}

func check(input [][]int, ikey int, jkey int, kkey int, mapv map[int]int) bool {
	fmt.Println("check")
	n := len(input)
	fmt.Println("Input", input)
	for i := 0; i < n; i++ {
		if ikey == input[i][0] && jkey == input[i][1] && kkey == input[i][2] {
			return true
		} else if ikey == input[i][0] && kkey == input[i][1] && jkey == input[i][2] {
			return true
		} else if jkey == input[i][0] && ikey == input[i][1] && kkey == input[i][2] {
			return true
		} else if jkey == input[i][0] && kkey == input[i][1] && ikey == input[i][2] {
			return true
		} else if kkey == input[i][0] && ikey == input[i][1] && jkey == input[i][2] {
			return true
		} else if kkey == input[i][0] && jkey == input[i][1] && ikey == input[i][2] {
			return true
		}

	}
	return false
}
