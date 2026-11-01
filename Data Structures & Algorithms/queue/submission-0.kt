class Deque {
    data class Node (
        var prev: Node?,
        var next: Node?,
        val value: Int
    )

    var head: Node? = null
    var tail: Node? = null

    fun isEmpty(): Boolean {
        return head == null
    }

    fun append(value: Int) {
        if (head == null) {
            val node = Node(null, null, value)
            head = node
            tail = node
            return
        }
        val node = Node(tail, null, value)
        tail!!.next = node
        tail = node
    }

    fun appendleft(value: Int) {
        if (head == null) {
            val node = Node(null, null, value)
            head = node
            tail = node
            return
        }
        val node = Node(null, head, value)
        head!!.prev = node
        head = node
    }

    fun pop(): Int {
        if (tail == null) return -1
        val result = tail!!.value
        tail = tail!!.prev
        tail?.next = null
        if (tail == null) head = null
        return result
    }

    fun popleft(): Int {
        if (head == null) return -1
        val result = head!!.value
        head = head!!.next
        head?.prev = null
        if (head == null) tail = null
        return result
    }
}
