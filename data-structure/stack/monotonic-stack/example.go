package main

import "fmt"

/*
input: 5,3,1,2,4
return: -1 3 1 1 -1
*/

func main(){
	input:=[]int{5,3,1,2,4}
	fmt.Println(input)
	fmt.Println(nextExceed(input))
}

func nextExceed(input []int) []int{
	n:=len(input)
	result := make([]int, n)
	for i:=0;i<n;i++{
		result[i]=-1
	}
	stack := make([]int,0)
	for i:=0;i<n;i++{
		for len(stack)>0 && input[stack[len(stack)-1]]<input[i]{
			result[stack[len(stack)-1]]=i-stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return result
}