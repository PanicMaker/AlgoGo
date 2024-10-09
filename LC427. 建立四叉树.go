package algogo

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct1(grid [][]int) *Node {

	isSame := func(a, b, x, y int) bool {
		flag := grid[a][b]

		for i := a; i < x+1; i++ {
			for j := b; j < y+1; j++ {
				if flag != grid[i][j] {
					return false
				}
			}
		}

		return true
	}

	var build func(a, b, x, y int) *Node
	build = func(a, b, x, y int) *Node {

		if isSame(a, b, x, y) {
			return &Node{
				Val:    grid[a][b] == 1,
				IsLeaf: true,
			}
		}

		node := &Node{
			Val:         grid[a][b] == 1,
			IsLeaf:      isSame(a, b, x, y),
			TopLeft:     build(a, b, a+(x-a)/2, b+(y-b)/2),
			TopRight:    build(a, b+(y-b)/2+1, a+(x-a)/2, y),
			BottomLeft:  build(a+(x-a)/2+1, b, x, b+(y-b)/2),
			BottomRight: build(a+(x-a)/2+1, b+(y-b)/2+1, x, y),
		}

		return node
	}

	n := len(grid)
	return build(0, 0, n-1, n-1)
}

func construct2(grid [][]int) *Node {

	var build func(a, b, x, y int) *Node
	build = func(a, b, x, y int) *Node {

		flag := grid[a][b]

		for i := a; i < x+1; i++ {
			for j := b; j < y+1; j++ {
				if flag != grid[i][j] {
					return &Node{
						Val:         grid[a][b] == 1,
						IsLeaf:      false,
						TopLeft:     build(a, b, a+(x-a)/2, b+(y-b)/2),
						TopRight:    build(a, b+(y-b)/2+1, a+(x-a)/2, y),
						BottomLeft:  build(a+(x-a)/2+1, b, x, b+(y-b)/2),
						BottomRight: build(a+(x-a)/2+1, b+(y-b)/2+1, x, y),
					}
				}
			}
		}

		return &Node{
			Val:    grid[a][b] == 1,
			IsLeaf: true,
		}
	}

	n := len(grid)
	return build(0, 0, n-1, n-1)
}

func construct3(grid [][]int) *Node {
	n := len(grid)
	preSum := make([][]int, n+1)
	preSum[0] = make([]int, n+1)
	for i := range grid {
		preSum[i+1] = make([]int, n+1)
		for j, v := range grid[i] {
			preSum[i+1][j+1] = preSum[i+1][j] + preSum[i][j+1] - preSum[i][j] + v
		}
	}

	var build func(a, b, x, y int) *Node
	build = func(a, b, x, y int) *Node {
		sum := preSum[x+1][y+1] - preSum[x+1][b] - preSum[a][y+1] + preSum[a][b]

		if sum == 0 || sum == (x-a+1)*(y-b+1) {
			return &Node{
				Val:    grid[a][b] == 1,
				IsLeaf: true,
			}
		}

		return &Node{
			Val:         grid[a][b] == 1,
			IsLeaf:      false,
			TopLeft:     build(a, b, a+(x-a)/2, b+(y-b)/2),
			TopRight:    build(a, b+(y-b)/2+1, a+(x-a)/2, y),
			BottomLeft:  build(a+(x-a)/2+1, b, x, b+(y-b)/2),
			BottomRight: build(a+(x-a)/2+1, b+(y-b)/2+1, x, y),
		}
	}

	return build(0, 0, n-1, n-1)
}
