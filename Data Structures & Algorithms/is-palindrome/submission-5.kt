class Solution {
    fun isPalindrome(s: String): Boolean {
        var left = 0
        var right = s.length - 1

        while (left < right) {
            while (left < right && !s[left].isLetterOrDigit()) {
                left++
            }
            while(right > left && !s[right].isLetterOrDigit()) {
                right--
            }
            if (s[left].uppercase() != s[right].uppercase()) {
                return false
            } 
            left++
            right--
        }
        return true
    }
}
