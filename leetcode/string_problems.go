package leetcode

import (
	"sort"
	"strings"
	"unicode"
)

/**
QUESTION - 125. Valid Palindrome
TAGS: Easy | String

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
*/
func IsPalindrome(s string) bool {
    s = strings.ToLower(s)

    i, j := 0, len(s)-1
    for i <= j {
        for i <= j{
            if !isAlphaNumeric(s[i]){
                i++
            }else if !isAlphaNumeric(s[j]){
                j--
            }else {
                if s[i] != s[j] {
                    return false
                }
                i++
                j--
            }
        }
    }
    return true
}

/**
ASCII
a 97 | z 122
A 65 | Z 90
0 48 | 9 57
**/
func isAlphaNumeric(num byte) bool{
    if num >= 48 && num <= 57 {
        return true
    }

    if num >= 65 && num <= 90 {
        return true
    }

    if num >= 97 && num <= 122 {
        return true
    }
    return false
}

func IsPalindrome2(s string) bool {
    s = strings.ToLower(s)

    sChars := []rune(s)
    i, j := 0, len(sChars)-1
    for i <= j{
        if !isAlphaNumericRune(sChars[i]){
            i++
        }else if !isAlphaNumericRune(sChars[j]){
            j--
        }else {
            if sChars[i] != sChars[j] {
                return false
            }
            i++
            j--
        }
    }
    return true
}

func isAlphaNumericRune(value rune) bool{
    return unicode.IsLetter(value) || unicode.IsNumber(value)
}


/**
QUESTION: 14. Longest Common Prefix
TAGS: Easy

Input: strs = ["flower","flow","flight"]
Output: "fl"

First Solution:
Brute Force: We take the first string, since the longest substring will be within it
We iterate over the characters and affirm it in other strings. 
Runtime: O(nm) where n is length of strs, and m is length of substring

Second Solution:
Intuition: We sort strs, this will give us the most dissimilar characters at the end
Then we find the most similar substring amidst the two.
Runtime O(mlogn) where n is length of strs, and m is length of substring
**/
func LongestCommonPrefixFirst(strs []string) string {
    firstString := strs[0]
    if len(strs) == 1 {
        return firstString
    }

    result := ""
    for i:=1; i<=len(firstString); i++{
        substring := firstString[0:i]

        for _, str := range strs {
            if !strings.HasPrefix(str, substring) {
                return result
            }
        }
        result = substring
    }
    return result
}

func LongestCommonPrefixSecond(strs []string) string {
    strLength := len(strs)
    if strLength == 1 {
        return strs[0]
    }

    sort.Strings(strs)
    
    minLength := min(len(strs[0]), len(strs[strLength-1]))

    for i:=0; i<minLength; i++{
        if strs[0][i] != strs[strLength-1][i] {
            return strs[0][:i]
        }
    }
    return strs[0]
}