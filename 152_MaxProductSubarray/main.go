package main

import "fmt"

func main()  {
	fmt.Println(maxProduct([]int{1,2,3,4,-2}))
}

func maxProduct(nums []int)int  {
	if nums == nil {
		return 0
	}

	res,curMax,curMin := nums[0],nums[0],nums[0]

	for i:=1;i<= len(nums)-1 ;i++  {
		num := nums[i]
		curMax,curMin = curMax*num,curMin*num
		curMax,curMin = max(curMax,curMin,num) , min(curMax,curMin,num)

		if curMax > res {
			res = curMax
		}
	}

	return res
}

func max(a,b,c int)int  {
	if a>b&&a>c {
		return a
	}
	if b>a&&b>c {
		return b
	}
	return c
}

func min(a,b,c int)int  {
	if a<=b&&a<=c {
		return a
	}
	if b<=a&&b<=c {
		return b
	}
	return c
}