class Graph {

    val nodes = HashMap<Int, MutableSet<Int>>()

    fun addEdge(src: Int, dst: Int) {
        nodes.getOrPut(src) { mutableSetOf<Int>() }.add(dst)
    }

    fun removeEdge(src: Int, dst: Int): Boolean {
        return nodes[src]?.remove(dst) ?: false
    }

    fun hasPath(src: Int, dst: Int): Boolean {
        return bfs(src, dst) >= 0
    }

    fun bfs(src: Int, dst: Int): Int {
        if (src == dst) return 0
        var len = 0
        val deque = ArrayDeque<Int>()
        deque.add(src)
        val visited = mutableSetOf<Int>()
        while (deque.isNotEmpty()) {
            val cur = deque.pop()
            visited.add(cur)
            for (nbr in nodes[cur] ?: emptySet()) {
                if (nbr !in visited) {
                    if (nbr == dst) return len
                    visited.add(nbr)
                    deque.add(nbr)
                }
            }
            len++
        }
        return -1
    }
}
