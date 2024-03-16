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

func Merge_sorted_array_without_using_another_array(nums1 []int, m int, nums2 []int, n int) []int {
	total := m + n - 1
	J_done := n
	for i, j := m-1, n-1; i >= 0 && j >= 0 && total > 0; {
		if nums1[i] < nums2[j] {
			nums1[total] = nums2[j]
			j--
			total--
		} else if nums1[i] == nums2[j] {
			nums1[total] = nums2[j]
			total--
			nums1[total] = nums1[i]
			j--
			i--
			total--
		} else if nums1[i] > nums2[j] {
			nums1[total] = nums1[i]
			i--
			total--
		}
		J_done = j
	}
	for x := total; x >= 0 && J_done >= 0 && n > 0; {
		nums1[x] = nums2[total]
		total--
		x--
	}
	return nums1

}
