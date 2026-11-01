// Definition for a pair
// class Pair(var key: Int, var value: String)

class Solution {
    fun mergeSort(pairs: MutableList<Pair>): MutableList<Pair> {
        if (pairs.size <= 1) return pairs
        val mid = pairs.size / 2
        return merge(
            mergeSort(pairs.subList(0, mid)), 
            mergeSort(pairs.subList(mid, pairs.size))
        )
    }

    fun merge(a: List<Pair>, b: List<Pair>): MutableList<Pair> {
        var ia = 0
        var ib = 0
        return ArrayList<Pair>(a.size + b.size).apply {
            while (ia < a.size && ib < b.size) {
                if (a[ia].key <= b[ib].key) {
                    add(a[ia++])
                } else {
                    add(b[ib++])
                }
            }
            while (ia < a.size) {
                add(a[ia++])
            }
            while (ib < b.size) {
                add(b[ib++])
            }
        }
    }
}
