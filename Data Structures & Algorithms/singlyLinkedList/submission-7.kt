class LinkedList {
    data class Node(var next: Node?, val value: Int)

    private var head: Node? = null

    fun get(index: Int): Int {
        var curNode: Node? = head ?: return -1
        for (i in 0..<index) {
            curNode = curNode?.next
            if (curNode == null) return -1
        }
        return curNode?.value ?: return -1
    }

    fun insertHead(value: Int) {
        head = Node(head, value)
    }

    fun insertTail(value: Int) {
        var curNode: Node? = head
        if (head == null) {
            head = Node(null, value)
            return
        }
        while (curNode?.next != null) curNode = curNode?.next
        curNode!!.next = Node(null, value)
    }

    fun remove(index: Int): Boolean {
        if (head == null) return false
        if (index == 0) {
            head = head?.next
            return true
        }
        var curNode: Node? = head
        for (i in 0..<index-1) {
            if (curNode?.next == null) return false
            curNode = curNode?.next
        }
        if (curNode?.next == null) return false
        curNode!!.next = curNode!!.next?.next
        return true
    }

    fun getValues(): List<Int> {
        var curNode: Node? = head
        val result = mutableListOf<Int>()
        while (curNode != null) {
            result.add(curNode!!.value)
            curNode = curNode?.next
        }
        return result
    }
}
