class Solution {
    data class CountedString(val s: String, val m: HashMap<Char, Int>) {
        companion object {
            fun fromString(s: String): CountedString {
                val map = HashMap<Char, Int>()
                s.forEach{
                    map[it] = map.getOrPut(it){0} + 1
                }
                return CountedString(s, map)
            }
        }
    }

    fun groupAnagrams(strs: Array<String>): List<List<String>> {
        return strs.map(CountedString::fromString)
            .groupBy({it.m}, {it.s})
            .values.toList()
    }
}
