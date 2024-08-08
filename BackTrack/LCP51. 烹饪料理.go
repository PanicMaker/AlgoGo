package BackTrack

// https://leetcode.cn/problems/UEcfPD/

func perfectMenu(materials []int, cookbooks [][]int, attribute [][]int, limit int) int {
	n := len(cookbooks)
	m := len(materials)

	mw := 0     // 美味度
	bf := 0     // 饱腹感
	max_mw := 0 // 最大美味度

	var dfs func(i int)
	dfs = func(i int) {
		// 食材是否满足
		isSupport := true
		for j := 0; j < m; j++ {
			if materials[j] < 0 {
				isSupport = false
			}
		}

		// 食材满足且饱腹感大于等于limit才更新最大美味度
		if isSupport && bf >= limit {
			max_mw = max(max_mw, mw)
		}

		if i == n {
			return
		}

		// 不做当前料理
		dfs(i + 1)

		// 做当前料理
		mw += attribute[i][0]
		bf += attribute[i][1]
		for j := 0; j < m; j++ {
			materials[j] -= cookbooks[i][j]
		}
		dfs(i + 1)
		mw -= attribute[i][0]
		bf -= attribute[i][1]
		for j := 0; j < m; j++ {
			materials[j] += cookbooks[i][j]
		}
	}

	dfs(0)

	if max_mw != 0 {
		return max_mw
	}

	return -1
}
