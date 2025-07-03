package String

// https://www.nowcoder.com/practice/2b317e02f14247a49ffdbdba315459e7

func compare(version1 string, version2 string) int {
	// write code here

	i, j := 0, 0

	for i < len(version1) || j < len(version2) {
		sum1, sum2 := 0, 0

		for ; i < len(version1) && version1[i] != '.'; i++ {
			sum1 = sum1*10 + int(version1[i]-'0')
		}
		i++

		for ; j < len(version2) && version2[j] != '.'; j++ {
			sum2 = sum2*10 + int(version1[j]-'0')
		}
		j++

		if sum1 < sum2 {
			return -1
		} else if sum1 > sum2 {
			return 1
		}
	}
	return 0
}
