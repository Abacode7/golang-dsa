package leetcode

/**
QUESTION 1 - Two Sums 
TAGS: Easy | Array

nums = [7,11,2,15], target = 9
Map[7:0, 11:1]

- Check if value is the pair of the any previous element
    - If yes, return it's index, with the current value index
- Store item and itemIndex pair in set for a future pair
*/

func TwoSum(nums []int, target int) []int {
    pairMap := make(map[int]int, 0)
    numsLength := len(nums)

    for i := 0; i < numsLength; i++{
        currentValue := nums[i]
        prevValue := target - currentValue
        if prevValueIndex, ok := pairMap[prevValue]; ok {
            return []int{prevValueIndex, i}
        }
        pairMap[currentValue] = i
    }
    return nil
}



/**
QUESION AE - IsValidSubquence 
TAGS: Easy | Array

array
[5 1 22 25 6 -1 8 10]
sequence
[1 6 -1 10]
Answer: true
**/

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
    i, j := 0, 0
    for i < len(array) && j < len(sequence){
        if array[i] == sequence[j] {
            j++
        }
        i++
    }
	return j == len(sequence)
}



/**
QUESTION 977 - Squares of a Sorted Array 
TAGS: Easy | Array

Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]

To avoid an nlogn solution
- We use two pointer from behind
**/

func SortedSquares1(nums []int) []int {
    i, j := 0, len(nums) - 1
    result := make([]int, 0)

    // Find squares in reverse order
    for i <= j {
        iValueSquare := square(nums[i])
        jValueSquare := square(nums[j])

        if iValueSquare <= jValueSquare {
            result = append(result, jValueSquare)
            j--
        }else{
            result = append(result, iValueSquare)
            i++
        }
    }

    result = reverse(result)
    return result
}

func SortedSquares2(nums []int) []int {
    numsLength := len(nums)
    i, j := 0, numsLength - 1
    result := make([]int, numsLength)
    resultIndex := numsLength - 1

    // Find squares in reverse order
    for i <= j {
        iValueSquare := square(nums[i])
        jValueSquare := square(nums[j])

        if iValueSquare <= jValueSquare {
            result[resultIndex] = jValueSquare
            j--
        }else{
            result[resultIndex] = iValueSquare
            i++
        }
        resultIndex--
    }
    return result
}

func square(num int) int {
    return num * num
}

func reverse(nums []int) []int {
    i, j := 0, len(nums) -1
    for i <= j {
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
    return nums 
}



/**
QUESTION 2529 - Maximum Count of Positive Integer and Negative Integer 
TAGS: Easy | Array

Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]

To avoid an nlogn solution
- We use two pointer from behind
**/
func MaximumCount(nums []int) int {
    pos, neg := 0, 0
    for _, value := range nums {
        if value > 0 {
            pos++
        }else if value < 0{
            neg++
        }
    }

    if pos >= neg {
        return pos
    }else{
        return neg
    }
}
