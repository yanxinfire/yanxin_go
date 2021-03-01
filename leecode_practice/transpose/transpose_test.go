package transpose

import "testing"

func TestA(t *testing.T)  {
	a:=[][]int{{1,2,3},{4,5,6},{7,8,9}}
	t.Log(transpose(a))
}