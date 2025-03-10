package leetcode

import "math"

/**
Quetions 2138 - Divide a String into groups of size k - Easy

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

