package middle

type point struct {
	x, y int
}

func numOfIslands(grid [][]byte) int {
	ans := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 1 {
				ans++
				dfs(grid, x, y)
			}
		}
	}
	return ans
}

// 深度优先遍历 deepth first search
func dfs(grid [][]byte, x, y int) {

	if x < 0 || y < 0 || y > len(grid)-1 || x > len(grid[0])-1 {
		return
	}

	if grid[y][x] == 1 {
		grid[y][x] = 0
	} else {
		return
	}

	dfs(grid, x, y-1)
	dfs(grid, x, y+1)
	dfs(grid, x-1, y)
	dfs(grid, x+1, y)
}

func numOfIslands2(grid [][]byte) int {
	var ans int
	queue := make([]point, 0, len(grid)*len(grid[0]))
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 1 {
				ans++
				queue = append(queue, point{
					x: x,
					y: y,
				})
				bfs(grid, &queue)
			}
		}
	}
	return ans
}

// 广度优先遍历 breadth first search
func bfs(grid [][]byte, queue *[]point) {
	tmp := (*queue)[0]
	*queue = (*queue)[1:]

	if tmp.y-1 >= 0 && grid[tmp.y-1][tmp.x] == 1 {
		grid[tmp.y+1][tmp.x] = 2
		*queue = append(*queue, point{
			x: tmp.x,
			y: tmp.y - 1,
		})
	}

	if tmp.y+1 < len(grid) && grid[tmp.y+1][tmp.x] == 1 {
		grid[tmp.y+1][tmp.x] = 2
		*queue = append(*queue, point{
			x: tmp.x,
			y: tmp.y + 1,
		})
	}

	if tmp.x+1 < len(grid[0]) && grid[tmp.y][tmp.x+1] == 1 {
		grid[tmp.y][tmp.x+1] = 2
		*queue = append(*queue, point{
			x: tmp.x + 1,
			y: tmp.y,
		})
	}

	if tmp.x-1 >= 0 && grid[tmp.y][tmp.x-1] == 1 {
		grid[tmp.y][tmp.x-1] = 2
		*queue = append(*queue, point{
			x: tmp.x - 1,
			y: tmp.y,
		})
	}

	if len(*queue) != 0 {
		bfs(grid, queue)
	}
}
