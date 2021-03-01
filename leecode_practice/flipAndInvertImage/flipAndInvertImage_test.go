package faii

import "testing"

func Test(t *testing.T)  {
	A:=[][]int{{1,1,0},{1,0,1},{0,0,0}}
	t.Log(flipAndInvertImage(A))
}
