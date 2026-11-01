class TreeMap {
    data class Node(
        var left: Node?,
        var right: Node?,
        var key: Int,
        var value: Int
    )

    var root: Node? = null

    fun insert(key: Int, value: Int) {
        if (root == null) {
            root = Node(null, null, key, value) 
            return
        }
        var node: Node = requireNotNull(root)
        while (true) {
            if (node.key == key) {
                node.value = value
                return
            }

            if (node.key > key) {
                if (node.left == null) {
                    node.left = Node(null, null, key, value)
                } else {
                    node = node.left!!
                    continue
                }
            } else {
                if (node.right == null) {
                    node.right = Node(null, null, key, value)
                } else {
                    node = node.right!!
                    continue
                }
            }
        }
    }

    fun get(key: Int): Int {
        var node: Node = root ?: return -1
        while (true) {
            if (node.key == key) return node.value
            if (node.key > key) {
                if (node.left == null) return -1
                node = node.left!!
                continue
            } else {
                if (node.right == null) return -1
                node = node.right!!
                continue
            }
        }
    }

    fun getMin(): Int {
        var node: Node = root ?: return -1
        while (true) {
            if (node.left == null) return node.value
            node = node.left!!
        }
    }

    fun getMax(): Int {
        var node: Node = root ?: return -1
        while (true) {
            if (node.right == null) return node.value
            node = node.right!!
        }
    }

    fun remove(key: Int) {
        root = remove(root, key)
    }

    fun getInorderKeys(): List<Int> {
        if (root == null) return emptyList()
        val result = mutableListOf<Int>()
        getInorderKeys(root!!, result)
        return result
    }

    private fun getInorderKeys(node: Node, resultList: MutableList<Int>) {
        if (node.left != null) {
            getInorderKeys(node.left!!, resultList)
        }
        resultList.add(node.key)
        if (node.right != null) {
            getInorderKeys(node.right!!, resultList)
        }
    }

    private fun remove(node: Node?, key: Int): Node? {
        if (node == null) return null
        if (node.key != key) {
            if (node.key > key) {
                node.left = remove(node.left, key)
                return node
            } else {
                node.right = remove(node.right, key)
                return node
            }
        } else {
            if (node.left == null) {
                return node.right
            }
            if (node.right == null) {
                return node.left
            }
            var minNode = node.right!!
            while (minNode.left != null) minNode = minNode.left!!
            return Node(node.left, remove(node.right!!, minNode.key), minNode!!.key, minNode!!.value)
        }
    }
}