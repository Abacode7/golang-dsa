package leetcode

import (
	"testing"
)


func TestDivideString(t *testing.T){
	/**
	Input: s = "abcdefghij", k = 3, fill = "x"
	Output: ["abc","def","ghi","jxx"]
	**/
	
	s := "abcdefghij"
	k := 3
	fill := byte('x')

	result := divideString(s, k, fill)
	if result[0] != "abc"{
		t.Error("divideString of s = abcdefghij, k = 3, fill = x should produce abc first")
	}
	if result[1] != "def"{
		t.Error("divideString of s = abcdefghij, k = 3, fill = x should produce def first")
	}
	if result[2] != "ghi"{
		t.Error("divideString of s = abcdefghij, k = 3, fill = x should produce ghi first")
	}
	if result[3] != "jxx"{
		t.Error("divideString of s = abcdefghij, k = 3, fill = x should produce jxx first")
	}
}

func TestRemoveElements(t *testing.T){
	/**
	Input: nums = [0,1,2,2,3,0,4,2], val = 2
	Output: 5, nums = [0,1,4,0,3,_,_,_]
	**/
	
	nums := []int{0,1,2,2,3,0,4,2}
	result := RemoveElementFirst(nums, 2)
	if result != 5 {
		t.Error("remove elements should produce correct output 5")
	}
}