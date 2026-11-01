
type MinHeap struct {
	heap []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		heap: make([]int, 1),
	}
}

func (mh *MinHeap) Push(val int) {
	mh.heap = append(mh.heap, val)
	mh.bubbleUp(len(mh.heap) - 1)
}

func (mh *MinHeap) bubbleUp(pos int) {
	if pos <= 1 {
		return
	}
	parent := pos / 2
	if mh.heap[pos] < mh.heap[parent] {
		mh.heap[pos], mh.heap[parent] = mh.heap[parent], mh.heap[pos]
		mh.bubbleUp(parent)
	}
}

func (mh *MinHeap) sinkDown(pos int) {
	if pos*2 >= len(mh.heap) {
		return
	}
	leftChild, rightChild := pos*2, pos*2+1
	if rightChild >= len(mh.heap) {
		if mh.heap[pos] > mh.heap[leftChild] {
			mh.heap[pos], mh.heap[leftChild] = mh.heap[leftChild], mh.heap[pos]
		}
		return
	} else if mh.heap[leftChild] < mh.heap[rightChild] {
		if mh.heap[pos] > mh.heap[leftChild] {
			mh.heap[pos], mh.heap[leftChild] = mh.heap[leftChild], mh.heap[pos]
			mh.sinkDown(leftChild)
		}
		return
	} else {
		if mh.heap[pos] > mh.heap[rightChild] {
			mh.heap[pos], mh.heap[rightChild] = mh.heap[rightChild], mh.heap[pos]
			mh.sinkDown(rightChild)
		}
		return
	}
}

func (mh *MinHeap) Pop() int {
	if len(mh.heap) <= 1 {
		return -1
	}
	result := mh.heap[1]
	if len(mh.heap) > 2 {
		mh.heap[1] = mh.heap[len(mh.heap)-1]
		mh.heap = mh.heap[:len(mh.heap)-1]
		mh.sinkDown(1)
	} else {
		mh.heap = mh.heap[:1]
	}
	return result
}

func (mh *MinHeap) Top() int {
	if len(mh.heap) <= 1 {
		return -1
	}
	return mh.heap[1]
}

func (mh *MinHeap) Heapify(nums []int) {
	mh.heap = make([]int, 1)
	mh.heap = append(mh.heap, nums...)
	for i := len(mh.heap) / 2; i >= 1; i-- {
		mh.sinkDown(i)
	}
}
