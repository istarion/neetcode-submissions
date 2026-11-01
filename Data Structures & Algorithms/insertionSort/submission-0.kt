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
        var place = pos
        for (i in 0..<pos) {
            if (key < pairs[i].key) {
                place = i
                break
            }
        }
        return place
    }
}
