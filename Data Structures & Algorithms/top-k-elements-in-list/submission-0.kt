class Solution {
    fun topKFrequent(nums: IntArray, k: Int): IntArray {
        val freqMap = mutableMapOf<Int, Int>()
        nums.forEach {
            freqMap[it] = freqMap.getOrPut(it) {0} + 1
        }
        val result = mutableListOf<Int>()
        for (i in 0..<k) {
            val maxEntry = freqMap.maxByOrNull { it.value } ?: break
            result.add(maxEntry.key)
            freqMap.remove(maxEntry.key)
        }
        return result.toIntArray()
    }
}
