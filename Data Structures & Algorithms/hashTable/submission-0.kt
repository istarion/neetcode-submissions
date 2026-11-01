class HashTable(capacity: Int) {
    private var keyArr  = IntArray(capacity) { 0 }
    private var valArr  = IntArray(capacity) { -1 }
    private var isZeroExists: Boolean = false
    private var size: Int = 0
    private var capacity = capacity

    fun insert(key: Int, value: Int) {
        if (key == 0) {
            valArr[0] = value
            if (!isZeroExists) {
                isZeroExists = true
                size++
                if (size >= capacity / 2) {
                    resize()
                }
            }
            return
        }
        var curIndex = getBucket(key)
        while (keyArr[curIndex] != 0 && keyArr[curIndex] != key) {
            curIndex++
            if (curIndex == capacity) curIndex = 1
        }
        if (keyArr[curIndex] == 0) size++
        keyArr[curIndex] = key
        valArr[curIndex] = value
        if (size >= capacity / 2) {
            resize()
        }
    }

    fun get(key: Int): Int {
        val index = findIndex(key)
        return valArr[index]
    }

    fun remove(key: Int): Boolean {
        if (key == 0) {
            val removed = isZeroExists
            if (isZeroExists) {
                valArr[0] = -1
                size--
                isZeroExists = false
            }
            return removed
        }
        val index = findIndex(key)
        if (keyArr[index] == 0) return false
        keyArr[index] = 0
        valArr[index] = -1
        size--
        return true
    }

    fun getSize(): Int {
        return size
    }

    fun getCapacity(): Int {
        return capacity
    }

    fun resize() {
        capacity *= 2
        val newKeyArr = IntArray(capacity) { 0 }
        val newValArr = IntArray(capacity) { -1 }
        for (i in 0..<keyArr.size) {
            val key = keyArr[i]
            newKeyArr[getBucket(key)] = key
            newValArr[getBucket(key)] = valArr[i]
        }
        keyArr = newKeyArr
        valArr = newValArr
    }

    private fun getBucket(key: Int): Int {
        if (key == 0) return 0
        val bucket = key % capacity
        if (bucket == 0) return 1
        return bucket
    }

    private fun findIndex(key: Int): Int {
        if (key == 0) return 0
        var curIndex = getBucket(key)
        while (keyArr[curIndex] != key && keyArr[curIndex] != 0) {
            curIndex++
            if (curIndex == capacity) curIndex = 1
        }
        return curIndex
    }
}
