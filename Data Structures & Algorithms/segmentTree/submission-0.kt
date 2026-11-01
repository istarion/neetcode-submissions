class SegmentTree(nums: IntArray) {
    val root: TreeNode = TreeNode(nums, 0, nums.size - 1)

    class TreeNode(val nums: IntArray, val l: Int, val r: Int) {
        private var leftBranch: TreeNode? = null
        private var rightBranch: TreeNode? = null
        private var sum: Int = 0

        init {
            if (l == r) {
                sum = nums[l]
            } else {
                val mid = (r - l) / 2 + l
                leftBranch = TreeNode(nums, l, mid)
                rightBranch = TreeNode(nums, mid + 1, r)
                sum = (leftBranch?.sum ?: 0) + (rightBranch?.sum ?: 0)
            }
        }

        fun update(index: Int, value: Int) {
            if (r == l) {
                sum = value
                return
            }
            val mid = (r - l) / 2 + l
            if (index <= mid) {
                leftBranch?.update(index, value)
                sum = (leftBranch?.sum ?: 0) + (rightBranch?.sum ?: 0)
            } else {
                rightBranch?.update(index, value)
                sum = (leftBranch?.sum ?: 0) + (rightBranch?.sum ?: 0)
            }
        }

        fun query(L: Int, R: Int): Int {
            if (L <= l && r <= R) {
                return sum
            }
            val mid = (r - l) / 2 + l
            if (R <= mid) {
                return leftBranch?.query(L, R) ?: 0
            }
            if (L > mid) {
                return rightBranch?.query(L, R) ?: 0
            }
            return (leftBranch?.query(L, mid) ?: 0) + (rightBranch?.query(mid + 1, R) ?: 0)
        }
    }

    fun update(index: Int, value: Int) {
        root.update(index, value)
    }

    fun query(L: Int, R: Int): Int {
        return root.query(L, R)
    }
}
