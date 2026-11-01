// Definition for a pair
// class Pair(var key: Int, var value: String)

class Solution {
    fun insertionSort(pairs: MutableList<Pair>): List<List<Pair>> {
        val result = mutableListOf<List<Pair>>()
        var currentStep = pairs
        for (currentElement in 0..<pairs.size) {
            val newPlace = searchPlaceToInsert(pairs[currentElement].key, currentStep, currentElement)
            val element = currentStep.removeAt(currentElement)
            currentStep.add(newPlace, element)
            result.add(currentStep.toList())
        }
        return result
    }

    fun searchPlaceToInsert(key: Int, pairs: MutableList<Pair>, pos: Int): Int {
        return binarySearch(key, pairs, 0, pos)
    }

    fun binarySearch(key: Int, pairs: MutableList<Pair>, from: Int, to: Int): Int {
        if (from == to) {
            return from
        }
        val middle: Int = ((to - from) / 2) + from
        if (key >= pairs[middle].key) {
            return binarySearch(key, pairs, middle + 1, to)
        } else {
            return binarySearch(key, pairs, from, middle)
        }
    }
}
