package leetcode

import "math"

/**
QUESTION 2138 - Divide a String into groups of size k - Easy

Input: s = "abcdefghi", k = 3, fill = "x"
Output: ["abc","def","ghi"]

Input: s = "abcdefghij", k = 3, fill = "x"
Output: ["abc","def","ghi","jxx"]
**/

func divideString(s string, k int, fill byte) []string {
	stringLength := len(s)
    startIndex := 0
    result := make([]string, 0)

    for startIndex < stringLength {
        endIndex := int(math.Min(float64(startIndex + k), float64(stringLength)))
        substring := s[startIndex:endIndex]
        if len(substring) < k {
            numOfFills := k - len(substring)
            for i:=0; i<numOfFills; i++{
                substring += string(fill)
            }
        }
        result = append(result, substring)
        startIndex += k
    }
    return result
}

/**
QUESTION 1 - Two Sums - Easy

nums = [7,11,2,15], target = 9
Map[7:0, 11:1]

- Check if value is the pair of the any previous element
    - If yes, return it's index, with the current value index
- Store item and itemIndex pair in set for a future pair
*/

func twoSum(nums []int, target int) []int {
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
