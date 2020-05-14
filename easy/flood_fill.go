package easy

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	bfs(image,sr,sc,image[sr][sc],newColor)
	return image
}

func bfs(image [][]int,sr int,sc int,value int,newColor int) {
	if sr < 0 || sc < 0 || sr >= len(image) || sc >= len(image[0]) {
		return
	}

	if image[sr][sc] == value {
		image[sr][sc] = newColor
	} else {
		return
	}

	bfs(image, sr-1, sc, value, newColor)
	bfs(image, sr, sc-1, value, newColor)
	bfs(image, sr+1, sc, value, newColor)
	bfs(image, sr, sc+1, value, newColor)
}
