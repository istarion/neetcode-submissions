class MinHeap {
    private val list: MutableList<Int> = ArrayList<Int>()
    init {
        list.add(Int.MIN_VALUE)
    }

    fun push(value: Int) {
        var idx = list.size
        list.add(value)
        while (list[idx / 2] > list[idx]) {
            swap(idx, idx / 2)
            idx = idx / 2
        }
    }

    private fun swap(a: Int, b: Int) {
        val c = list[a]
        list[a] = list[b]
        list[b] = c
    }

    private fun getSmallestChild(idx: Int): Int {
        if (list.size == idx * 2 + 1) {
            return idx * 2
        }
        if (list[idx * 2] > list[idx * 2 + 1]) {
            return idx * 2 + 1
        } else {
            return idx * 2
        }
    }

    fun pop(): Int {
        if (list.size == 1) return -1
        if (list.size == 2) {
            return list.removeLast()
        }
        swap(1, list.size - 1)
        val elem = list.removeLast()
        if (list.size == 2) {
            return elem
        }
        var idx = 1
        var smallestChild = getSmallestChild(idx) 
        while (list.size > idx * 2 && list[idx] > list[smallestChild]) {
            swap(idx, smallestChild)
            idx = smallestChild
            if (list.size > smallestChild * 2) {
                smallestChild = getSmallestChild(smallestChild)
            } else {
                break
            }
        }
        return elem
    }

    fun top(): Int {
        if (list.size == 1) return -1
        return list[1]
    }

    fun heapify(nums: MutableList<Int>) {
        nums.forEach {
            push(it)
        }
    }
}
