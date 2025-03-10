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