class Solution {
    fun isAnagram(s: String, t: String): Boolean {
        val countS = HashMap<Char, Int>()
        val countT = HashMap<Char, Int>()
        s.forEach {
            countS[it] = countS.getOrPut(it) {0} + 1
        }
        t.forEach {
            countT[it] = countT.getOrPut(it) {0} + 1
        }
        return countS == countT
    }
}
