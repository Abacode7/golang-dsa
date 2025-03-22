package leetcode

import (
	"fmt"
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

func TestFindIndexOfFirstOccurrenceInString(t *testing.T){
	result := FindIndexOfFirstOccurrenceInString("mississippi", "issip")
	if result != 4 {
		t.Error("find index of first occurrence should produce correct output 5")
	}
}

func TestValidPath(t *testing.T){
	edges := [][]int{{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4}, {4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4}}
	source := 5
	destination := 9
	result := ValidPathDFS(10, edges, source, destination)
	if result != true {
		t.Error("valid path dfs should return true")
	}
}

func TestTreeTraversal(t *testing.T){
	/**
	Binary Search Tree
				5
			3		7
		2	4	  6		8
	1  nil		
	**/

	tree := TreeNode{5, &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, &TreeNode{4, nil, nil}}, &TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{8, nil, nil}}}
	result := TreeInorderTraversal(&tree)
	if result != nil {
		fmt.Printf("inorder traversal result: %v\n", result)
	}

	result1 := TreePreorderTraversal(&tree)
	if result1 != nil {
		fmt.Printf("preorder traversal result: %v\n", result1)
	}

	result2 := TreePostorderTraversal(&tree)
	if result2 != nil {
		fmt.Printf("postorder traversal result: %v\n", result2)
	}
}

