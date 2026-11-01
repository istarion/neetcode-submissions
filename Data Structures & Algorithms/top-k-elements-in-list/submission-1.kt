class Solution {
    fun topKFrequent(nums: IntArray, k: Int): IntArray {
        val freqMap = mutableMapOf<Int, Int>()
        nums.forEach {
            freqMap[it] = freqMap.getOrPut(it) {0} + 1
        }

        val pq = PriorityQueue<Pair<Int, Int>>(compareBy {it.first})

        freqMap.forEach {
            pq.add(it.value to it.key)
            if (pq.size > k) {
                pq.poll()
            }
        }

        val result = IntArray(k)
        for (i in k-1 downTo 0) {
            result[i] = pq.poll().second
        }
        return result
    }
}
