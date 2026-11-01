class Solution {
    fun hasDuplicate(nums: IntArray): Boolean {
        val s = HashSet<Int>()
        for (n in nums) {
            if (n in s) {
                return true
            } else {
                s.add(n)
            }
        }
        return false
    }
}
