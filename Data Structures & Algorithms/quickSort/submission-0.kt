// Definition for a pair
// class Pair(var key: Int, var value: String)

class Solution {
    fun quickSort(pairs: MutableList<Pair>): MutableList<Pair> {
        quickSort(pairs, 0, pairs.size - 1)
        return pairs
    }

    private fun quickSort(pairs: MutableList<Pair>, l: Int, r: Int) {
        if (l >= r) {
            return
        }
        val pivot = pairs[r].key
        var spot = l
        for (i in l..<r) {
            if (pairs[i].key < pivot) {
                val tmp = pairs[spot]
                pairs[spot] = pairs[i]
                pairs[i] = tmp
                ++spot
            }
        }
        val tmp = pairs[spot]
        pairs[spot] = pairs[r]
        pairs[r] = tmp
        quickSort(pairs, l, spot - 1)
        quickSort(pairs, spot + 1, r)
    }
}
