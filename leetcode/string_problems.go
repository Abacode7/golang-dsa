package leetcode

import (
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