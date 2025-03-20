package leetcode

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
 }


 /**
QUESTION: 94. Binary Tree Inorder Traversal
TAG: Easy

Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
Output: [4,2,6,5,7,1,3,9,8]
 */
 
func TreeInorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    InorderTraversal(root, &result)
    return result
}

/**
Usecase: Useful for getting tree (binary search trees) values in sorted order
*/
func InorderTraversal(root *TreeNode, result *[]int){
    if root == nil {return}

    if root.Left != nil {
        InorderTraversal(root.Left, result)
    }

    *result = append(*result, root.Val)

    if root.Right != nil {
        InorderTraversal(root.Right, result)
    }
}


/**
Usecase: Useful for copying a tree (binary search tree), cos it maintains the original tree order
*/
func TreePreorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    PreorderTraversal(root, &result)
    return result
}

func PreorderTraversal(root *TreeNode, result *[]int){
	if root == nil {return}

	*result = append(*result, root.Val)

	if root.Left != nil {
		PreorderTraversal(root.Left, result)
	}

	if root.Right != nil {
		PreorderTraversal(root.Right, result)
	}
}


/**
Usecase: Useful for deleting a tree
*/
func TreePostorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    PostorderTraversal(root, &result)
    return result
}

func PostorderTraversal(root *TreeNode, result *[]int){
	if root == nil {return}

	if root.Left != nil {
		PostorderTraversal(root.Left, result)
	}

	if root.Right != nil {
		PostorderTraversal(root.Right, result)
	}

	*result = append(*result, root.Val)
}



/**
QUESTION: 100. Same Tree
TAG: Easy

Input: p = [1,2,3], q = [1,2,3]
Output: true
Input: p = [1,2], q = [1,null,2]
Output: false
*/
func IsSameTree(p *TreeNode, q *TreeNode) bool {
    return preorderTraversal(p, q)
}

func preorderTraversal(p, q *TreeNode) bool {
    if p == nil  && q == nil {
        return true
    }else if (p == nil && q != nil) || (p != nil && q == nil) {
        return false
    }else {
        if p.Val != q.Val {return false}
    }
    return preorderTraversal(p.Left, q.Left) && preorderTraversal(p.Right, q.Right)
}



/**
QUESTION: 101. Symmetric Tree
TAG: Easy

Input: root = [1,2,2,3,4,4,3]
Output: true
Input: root = [1,2,2,null,3,null,3]
Output: false
*/
func IsSymmetric(root *TreeNode) bool {
    if root == nil {return true}
    return traverse(root.Left, root.Right)
}

func traverse(leftNode, rightNode *TreeNode) bool {
    if leftNode == nil && rightNode == nil {return true}

    if leftNode == nil || rightNode == nil {
        return false
    }else{
        if leftNode.Val != rightNode.Val {return false}
    }
    return traverse(leftNode.Left, rightNode.Right) && traverse(leftNode.Right, rightNode.Left)
}