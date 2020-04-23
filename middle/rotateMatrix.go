package middle

import "fmt"

func rotate(matrix [][]int)  {
	n := len(matrix)
	for i :=0 ;i<n/2;i++ {
		matrix[i],matrix[n-i-1] = matrix[n-i-1],matrix[i]
	}

	fmt.Printf("翻转：%v\n",matrix)
	for i := 0 ;i<n; i++ {
		for j := 0;j<i;j++ {
			matrix[i][j],matrix[j][i] = matrix[j][i],matrix[i][j]
		}
	}

	fmt.Printf("对角线：%v\n",matrix)
}
