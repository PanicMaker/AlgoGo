package DataStructure

// https://leetcode.cn/problems/finding-pairs-with-a-certain-sum/description/

type FindSumPairs struct {
	nums1 []int
	nums2 []int
	dict  map[int]int
}

func Constructor1865(nums1 []int, nums2 []int) FindSumPairs {
	dict := make(map[int]int)
	for _, v := range nums2 {
		dict[v]++
	}

	return FindSumPairs{
		nums1: nums1,
		nums2: nums2,
		dict:  dict,
	}
}

func (this *FindSumPairs) Add(index int, val int) {
	oldVal := this.nums2[index]
	this.dict[oldVal]--
	this.nums2[index] += val
	this.dict[this.nums2[index]]++
}

func (this *FindSumPairs) Count(tot int) int {
	count := 0

	for _, v := range this.nums1 {
		if val, ok := this.dict[tot-v]; ok {
			count += val
		}
	}

	return count
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
