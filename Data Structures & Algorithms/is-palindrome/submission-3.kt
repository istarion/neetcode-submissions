class Solution {
    fun isPalindrome(s: String): Boolean {
        var left = 0
        var right = s.length - 1

        while (right - left > s.length % 2) {
            while (!s[left].isLetterOrDigit()) {
                left++
                if (left == right) return true
            }
            while(!s[right].isLetterOrDigit()) {
                right--
                if (left == right) return true
            }
            if (right - left <= s.length % 2) return true
            if (s[left].uppercase() != s[right].uppercase()) {
                return false
            } else {
                left++
                right--
            }
        }
        return true
    }
}
