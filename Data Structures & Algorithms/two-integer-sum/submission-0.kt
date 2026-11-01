class Solution {
    fun twoSum(nums: IntArray, target: Int): IntArray {
        val backIndex = HashMap<Int, Int>()
        nums.forEachIndexed { i, v ->
            val left = target - v
            if (left !in backIndex) {
                backIndex[v] = i
            } else {
                return intArrayOf(backIndex[left]!!, i)
            }
        }
        error("No pair")
    }
}
