class UnionFind(n: Int) {
    val size = n
    val parents = HashMap<Int, Int>()
    val ranks = HashMap<Int, Int>()
    private var numComps = n

    fun find(x: Int): Int {
        var curNode = x
        while(curNode in parents) {
            curNode = parents[curNode]!!
        }
        if (parents[x] != null) {
            parents[x] = curNode
        }
        return curNode
    }

    fun isSameComponent(x: Int, y: Int): Boolean {
        return find(x) == find(y)
    }

    fun union(x: Int, y: Int): Boolean {
        val xPar = find(x)
        val yPar = find(y)
        if (xPar == yPar) return false
        val rankX = ranks[xPar] ?: 0
        val rankY = ranks[yPar] ?: 0
        if (rankX >= rankY) {
            parents[y] = xPar
            if (rankX == rankY) {
                ranks[xPar] = ranks[xPar] ?: 0 + 1
            }
        } else {
            parents[xPar] = yPar
        }
        numComps--
        return true
    }

    fun getNumComponents(): Int {
        return numComps
    }
}
