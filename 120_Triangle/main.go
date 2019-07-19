package main

import "fmt"

func main()  {
	arr := [][]int{
		{1},{2,3},{4,5,6},{7,8,9,10},
	}
	fmt.Println(minimumTotal(arr))
}

func minimumTotal(triangle [][]int)int{
	mini := triangle[len(triangle)-1]
	fmt.Println("mini:",mini)
	i:=len(triangle) - 2
	for ; i>=0;i--  {
		for j:=0;j<len(triangle[i]) ;j++  {
			mini[j] = triangle[i][j] + min(mini[j],mini[j+1])
		}
		fmt.Println("mini:",mini)
	}
	return mini[0]
}

func min(a int, b int)int  {
	if b<a {
		return b
	}
	return a
}