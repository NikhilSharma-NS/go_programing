package merge_sorted_array

func Merge_sorted_array(nums1 []int, m int, nums2 []int, n int) []int {
	output := make([]int, m+n)
	count := 0
	var final_m, final_n int
	for i, j := 0, 0; i < m && j < n; {
		if nums1[i] > nums2[j] {
			output[count] = nums2[j]
			j++
			count++
		} else if nums1[i] == nums2[j] {
			output[count] = nums1[i]
			count++
			output[count] = nums2[j]
			i++
			j++
			count++
		} else {
			output[count] = nums1[i]
			i++
			count++
		}
		final_m = i
		final_n = j
	}
	if m != final_m {
		for i := final_m; i < m; i++ {
			output[count] = nums1[i]
			count++
		}
	} else if n != final_n {
		for i := final_n; i < n; i++ {
			output[count] = nums2[i]
			count++
		}
	}
	return output

}
