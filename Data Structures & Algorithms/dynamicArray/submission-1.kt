class DynamicArray(capacity: Int) {
    private var currentCapacity: Int = capacity
    private var arr: Array<Int> = Array<Int>(capacity) {0}
    private var allocated: Int = 0

    fun get(i: Int): Int {
        return arr[i]
    }

    fun set(i: Int, n: Int) {
        arr[i] = n
    }

    fun pushback(n: Int) {
        if (allocated >= currentCapacity) {
            resize()
        }
        arr[allocated] = n
        allocated++
    }

    fun popback(): Int {
        return arr[--allocated]
    }

    private fun resize() {
        val oldCapacity = currentCapacity
        currentCapacity = currentCapacity * 2
        var newArr = Array<Int>(currentCapacity) {0}
        for (i in 0..<oldCapacity) {
            newArr[i] = arr[i]
        }
        arr = newArr
    }

    fun getSize(): Int {
        return allocated
    }

    fun getCapacity(): Int {
        return currentCapacity
    }
}
